package main

import (
	"Tutorial/blockchain"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my API")
}

//Person Represents Struct of the Entity
type Person struct {
	Name           string  `json:"name"`
	Department     float64 `json:"department"`
	LifeStage      float64 `json:"lifestage"`
	Gender         float64 `json:"gender"`
	Comorbidity    float64 `json:"comorbidity"`
	Symptomatology float64 `json:"symptomatology"`
	IsPositive     string  `json:"positive"`
}

//EuclideanDistance calculate the distance beetween two points
func EuclideanDistance(instance1 Person, instance2 Person) float64 {
	var sum float64
	sum += math.Pow((instance1.Department - instance2.Department), 2)
	sum += math.Pow((instance1.LifeStage - instance2.LifeStage), 2)
	sum += math.Pow((instance1.Gender - instance2.Gender), 2)
	sum += math.Pow((instance1.Comorbidity - instance2.Comorbidity), 2)
	sum += math.Pow((instance1.Symptomatology - instance2.Symptomatology), 2)

	return math.Sqrt(sum)
}

type classVote struct {
	key   string
	value int
}

type sortVotes []classVote

func (sortedvotes sortVotes) Len() int { return len(sortedvotes) }
func (sortedvotes sortVotes) Less(i int, j int) bool {
	return sortedvotes[i].value < sortedvotes[j].value
}
func (sortedvotes sortVotes) Swap(i int, j int) {
	sortedvotes[i], sortedvotes[j] = sortedvotes[j], sortedvotes[i]
}
func contains(votesMap map[string]int, name string) bool {
	for s := range votesMap {
		if s == name {
			return true
		}
	}

	return false
}

func getResponse(neighbors []Person) sortVotes {
	classVotes := make(map[string]int)

	for i := range neighbors {
		response := neighbors[i].IsPositive
		if contains(classVotes, response) {
			classVotes[response]++
		} else {
			classVotes[response] = 1
		}
	}

	scv := make(sortVotes, len(classVotes))
	i := 0
	for k, v := range classVotes {
		scv[i] = classVote{k, v}
		i++
	}

	sort.Sort(sort.Reverse(scv))
	return scv
}

//getAccuracy obtains the percentage of correct predictions
func getAccuracy(testSet []Person, predictions []string) float64 {
	correct := 0

	for x := range testSet {
		if testSet[x].IsPositive == predictions[x] {
			correct++
		}
	}

	return (float64(correct) / float64(len(testSet))) * 100.00
}

type distancePair struct {
	record   Person
	distance float64
}

type distanceCompare []distancePair

func (slice distanceCompare) Len() int               { return len(slice) }
func (slice distanceCompare) Less(i int, j int) bool { return slice[i].distance < slice[j].distance }
func (slice distanceCompare) Swap(i int, j int)      { slice[i], slice[j] = slice[j], slice[i] }

//getNeighbors calculate the neighbors at a k distance
func getNeighbors(trainingSet []Person, testInstance Person, k int) []Person {
	var distances distanceCompare
	for i := range trainingSet {
		dist := EuclideanDistance(testInstance, trainingSet[i])
		distances = append(distances, distancePair{trainingSet[i], dist})
	}
	sort.Sort(distances)
	var neighbors []Person
	for x := 0; x < k; x++ {
		neighbors = append(neighbors, distances[x].record)
	}
	return neighbors
}

func receivePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Request-Method", "POST")
	var testInstance Person
	var testInstances []Person
	var predictions []string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Insert Valid Information")
	}
	json.Unmarshal(reqBody, &testInstance)
	k := 4
	json.NewEncoder(w).Encode(testInstance)
	testInstances = append(testInstances, testInstance)
	predictions = calculateKNN(persons, testInstances, k)
	testInstance.IsPositive = predictions[0]
	Perso := blockchain.AddPerson(testInstance.Name, testInstance.Department, testInstance.LifeStage, testInstance.Gender, testInstance.Comorbidity, testInstance.Symptomatology, testInstance.IsPositive)
	chain.AddBlock(Perso)
	addedPersons = append(addedPersons, testInstance)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(predictions)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addedPersons)
}

func calculateKNN(persons []Person, testInstances []Person, k int) []string {
	var predictions []string
	for i := range testInstances {
		neighbors := getNeighbors(persons, testInstances[i], k)
		result := getResponse(neighbors)
		predictions = append(predictions, result[0].key)
		//fmt.Printf("Predicted: %s, Actual: %s\n", result[0].key, flowerTestSet[i].Specie)
	}
	return predictions
}

//GlobalVariables
var persons []Person
var addedPersons []Person
var chain = blockchain.InitBlockChain()

func main() {
	csvFile, err := os.Open("persons.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Succesfully Opened CSV file")
	defer csvFile.Close()

	csvLines := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := csvLines.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		name := line[0]
		department, _ := strconv.ParseFloat(line[1], 8)
		lifestage, _ := strconv.ParseFloat(line[2], 8)
		gender, _ := strconv.ParseFloat(line[3], 8)
		comorbidity, _ := strconv.ParseFloat(line[4], 8)
		symptomatology, _ := strconv.ParseFloat(line[5], 8)
		ispositive := line[6]
		persons = append(persons, Person{
			Name:           name,
			Department:     department,
			LifeStage:      lifestage,
			Gender:         gender,
			Comorbidity:    comorbidity,
			Symptomatology: symptomatology,
			IsPositive:     ispositive,
		})
	}
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/knn", receivePerson).Methods("POST")
	router.HandleFunc("/persons", getPeople).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(router)))
	/*personsJSON, _ := json.Marshal(persons)
	fmt.Println(string(personsJSON))*/

}
