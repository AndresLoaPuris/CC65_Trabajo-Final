package api

import (
	"Tutorial/blockchain"
	"Tutorial/kmeans"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIPerson struct {
	Name      string    `json:Name`
	Values    []float64 `json:Values`
	Category  int       `json:Category`
	Algorithm string    `json:Algorithm`
}

type allPerson []APIPerson

var returnpeople = allPerson{}

var algorithmUse = ""
var team = 0

var chain = blockchain.InitBlockChain()

func IndexRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to API")
}

func GetPeople(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(returnpeople)

	/*for _, block := range chain.ReturnBlocks() {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Name in Block: %s\n", block.ReturnBlockName())
		fmt.Println("Data in Block:", block.ReturnBlockValues())
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}*/
}

func CreatePeople(w http.ResponseWriter, r *http.Request) {

	reqBody, error := ioutil.ReadAll(r.Body)

	if error != nil {
		fmt.Fprint(w, "Insert a Valid Person")
	}

	var people = allPerson{}

	json.Unmarshal(reqBody, &people)

	/*people.Category = 0
	algorithmUse = people.Algorithm
	team = people.Category

	if success, centroids := kmeans.Train(observations, team, 10); success {
		for i, observation := range observations {

			index := kmeans.Nearest(observation, centroids)
			people[i].Category = index + 1
			Person := blockchain.AddPerson(people[i].Name, observation, index+1, algorithmUse)
			chain.AddBlock(Person)
		}
	}

	people = append(people, people)
	observations = append(observations, kmeans.Node{people.Values[0], people.Values[1], people.Values[2], people.Values[3], people.Values[4]})
	*/
	team = people[0].Category
	var observations = []kmeans.Node{}

	for i := 0; i < len(people); i++ {
		observations = append(observations, kmeans.Node{people[i].Values[0], people[i].Values[1], people[i].Values[2], people[i].Values[3], people[i].Values[4]})
	}

	if success, centroids := kmeans.Train(observations, team, 100); success {
		for i, observation := range observations {
			index := kmeans.Nearest(observation, centroids)
			people[i].Category = index + 1
			Person := blockchain.AddPerson(people[i].Name, observation, index+1, algorithmUse)
			chain.AddBlock(Person)
		}
	}

	returnpeople = people

	for _, block := range chain.ReturnBlocks() {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Name in Block: %s\n", block.ReturnBlockName())
		fmt.Println("Data in Block:", block.ReturnBlockValues())
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(returnpeople)

}
