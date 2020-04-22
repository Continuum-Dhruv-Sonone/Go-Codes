package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Printf("Handler started")
	defer log.Printf("Handler Ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "Hello world")
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
