package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	blocks []*Block
}

type Person struct {
	name   string
	values []float64
}

type Block struct {
	Hash     []byte
	Data     Person
	PrevHash []byte
}

func AddPerson(name string, value []float64) Person {
	return Person{name, value}
}

func (b *Block) ReturnBlockValues() []float64 {
	return b.Data.values
}

func (b *Block) ReturnBlockName() string {
	return b.Data.name
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data Person, prevHash []byte) *Block {
	block := &Block{[]byte{}, data, prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data Person) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func (chain *BlockChain) ReturnBlocks() []*Block  {
	return chain.blocks
}

func Genesis() *Block {
	return CreateBlock(Person{"Genesis", make([]float64, 0)}, []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}