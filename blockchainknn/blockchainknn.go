package blockchainknn

import (
	"bytes"
	"crypto/sha256"
)

//BlockChain contains all the blocks
type BlockChain struct {
	blocks []*Block
}

//Person is the struct to abstract of a Person
type Person struct {
	Name           string
	Department     float64
	LifeStage      float64
	Gender         float64
	Comorbidity    float64
	Symptomatology float64
	IsPositive     string
}

//Block represents a Block of BlockChains
type Block struct {
	Hash     []byte
	Data     Person
	PrevHash []byte
}

//AddPerson function to add a Person to a Block
func AddPerson(name string, department float64, lifestage float64, gender float64, comorbidity float64, symptomatology float64, ispositive string) Person {
	return Person{name, department, lifestage, gender, comorbidity, symptomatology, ispositive}
}

//AddBlock to add the block to de blocks from Blockchain
func (chain *BlockChain) AddBlock(data Person) {
	prevBloc := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBloc.Hash)
	chain.blocks = append(chain.blocks, new)
}

//ReturnBlockValues return the value of Person Stored
func (b *Block) ReturnBlockValues() Person {
	return b.Data
}

//ReturnBlockName return The name of the block
func (b *Block) ReturnBlockName() string {
	return b.Data.Name
}

//DeriveHash gives a Hash based on PrevHash and using crypto/sha256
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

//CreateBlock creates a block an gives a hash
func CreateBlock(data Person, prevHash []byte) *Block {
	block := &Block{[]byte{}, data, prevHash}
	block.DeriveHash()
	return block
}

//ReturnBlocks returns all the blocks on the blockchain
func (chain *BlockChain) ReturnBlocks() []*Block {
	return chain.blocks
}

//Genesis it's the creation of the first block
func Genesis() *Block {
	return CreateBlock(Person{"Genesis", 0.00, 0.00, 0.00, 0.00, 0.00, "false"}, []byte{})
}

//InitBlockChain initializes the Blockchain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
