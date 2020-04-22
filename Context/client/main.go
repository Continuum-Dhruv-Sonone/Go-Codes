package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/", nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error from get: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}
	io.Copy(os.Stdout, res.Body)
}
