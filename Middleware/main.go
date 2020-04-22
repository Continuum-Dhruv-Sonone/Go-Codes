package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var circuitState = make(map[string]bool)

func handleRequests() {
	circuitState["Kafka"] = false
	circuitState["Cassandra"] = false
	circuitState["Cherwell"] = false
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", middleware(incidentClosure, "Kafka", "Cassandra", "Cherwell"))
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {

	handleRequests()
}

func middleware(next http.HandlerFunc, cbName ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		for _, name := range cbName {

			if circuitState[name] == true {

				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}
		next(w, r)
		//or use next.ServeHTTP(w,r)
	}
}

func incidentClosure(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pickle Rick!")
}
