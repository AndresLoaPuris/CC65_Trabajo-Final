package main

import (
	"Tutorial/api"
	"Tutorial/blockchain"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", api.IndexRouter)
	router.HandleFunc("/people", api.GetPeople).Methods("GET")
	router.HandleFunc("/people", api.CreatePeople).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(router)))

	chain := blockchain.InitBlockChain()

	for _, v := range api.Example() {
		Person := blockchain.AddPerson(v.Name, v.Values, v.Category, v.Algorithm)
		chain.AddBlock(Person)
	}

	for _, block := range chain.ReturnBlocks() {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Name in Block: %s\n", block.ReturnBlockName())
		fmt.Println("Data in Block:", block.ReturnBlockValues())
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
