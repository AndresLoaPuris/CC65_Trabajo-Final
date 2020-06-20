package api

import (
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

var people = allPerson{

	{
		Name:      "Jose Carlos Paredes",
		Values:    []float64{0.0034, 0.0012, 0.0123, 0.0071, 0.0016},
		Category:  1,
		Algorithm: "knn",
	},
}

func Example() []APIPerson {
	return people
}

func IndexRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to API")
}

func GetPeople(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func CreatePeople(w http.ResponseWriter, r *http.Request){
	var newPeople APIPerson
	reqBody , error := ioutil.ReadAll(r.Body)

	if error != nil {
		fmt.Fprint(w,"Insert a Valid Person")
	}
	json.Unmarshal(reqBody, &newPeople)
	people = append(people, newPeople)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPeople)
}