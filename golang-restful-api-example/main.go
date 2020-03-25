package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Person struct
type person struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
	Age  string `json:"Age"`
}

type allPeople []person

var people = allPeople{
	{
		Id:   "1",
		Name: "john",
		Age:  "22",
	},
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint - index")
	fmt.Fprintf(w, "index")
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var newPerson person
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newPerson)
	people = append(people, newPerson)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPerson)
}

//Get a Specific person based on the ID
func getOnePerson(w http.ResponseWriter, r *http.Request) {
	personID := mux.Vars(r)["id"]

	for _, person := range people {
		if person.Id == personID {
			json.NewEncoder(w).Encode(person)
		}
	}
}

func everyone(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	PersonID := mux.Vars(r)["id"]
	var updatedPerson person

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "invalid data.")
	}
	json.Unmarshal(reqBody, &updatedPerson)
	for i, person := range people {
		if person.Id == PersonID {
			person.Age = updatedPerson.Age
			person.Name = updatedPerson.Name
			people = append(people[:i], person)
			json.NewEncoder(w).Encode(person)
		}
	}
}

func removePerson(w http.ResponseWriter, r *http.Request) {
	personId := mux.Vars(r)["id"]

	for i, person := range people {
		if personId == person.Id {
			people = append(people[:i], people[i+1:]...)
			fmt.Fprintf(w, "%v was removed", personId)
		}
	}
}

func main() {
	//gorilla mux
	router := mux.NewRouter().StrictSlash(true)

	//root
	router.HandleFunc("/", index)
	//Add person
	router.HandleFunc("/add", createPerson).Methods("POST")
	//Get all People
	router.HandleFunc("/everyone", everyone).Methods("GET")
	//Get specific person
	router.HandleFunc("/person/{id}", getOnePerson).Methods("GET")
	//Update person
	router.HandleFunc("/person/{id}", updatePerson).Methods("PATCH")
	//Remove person
	router.HandleFunc("/person/{id}", removePerson).Methods("DELETE")

	//address
	log.Fatal(http.ListenAndServe(":9999", router))

}
