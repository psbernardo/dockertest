# dockertest

TDD - help us to check the correctness of the business logic with mocking some test data but testing with actual database or third party service will add another layer of testing since we testing not only the business logic, we aslo test the production code.


![1111111](https://user-images.githubusercontent.com/70035042/228185363-466a4fe3-f773-4c2f-8ed6-2ad61b7868bf.png)

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AllRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RequestURI)

}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RequestURI)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
}

type Response struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

func create(w http.ResponseWriter, r *http.Request) {
	data := Response{}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func handleRequests() {

	http.HandleFunc("/", AllRoute)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}

