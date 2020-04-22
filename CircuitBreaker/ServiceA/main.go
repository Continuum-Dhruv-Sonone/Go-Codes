package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {

	hystrix.ConfigureCommand("Dhruv", hystrix.CommandConfig{
		// How long to wait for command to complete, in milliseconds
		Timeout: 50000,

		// MaxConcurrent is how many commands of the same type
		// can run at the same time
		MaxConcurrentRequests: 300,

		// VolumeThreshold is the minimum number of requests
		// needed before a circuit can be tripped due to health
		RequestVolumeThreshold: 10,

		// SleepWindow is how long, in milliseconds,
		// to wait after a circuit opens before testing for recovery
		SleepWindow: 2000,

		// ErrorPercentThreshold causes circuits to open once
		// the rolling measure of errors exceeds this percent of requests
		ErrorPercentThreshold: 50,
	})

	http.HandleFunc("/", logger(handler))

	fmt.Println("ServiceA started")
	log.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}

// handler send request to Service B
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	resultCh := make(chan []byte)
	errCh := hystrix.Go("Dhruv", func() error {
		resp, err := http.Get("http://localhost:9090")
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		resultCh <- b
		return nil
	}, nil)

	select {
	case res := <-resultCh:
		log.Println("Received response from service B:", string(res))
		w.WriteHeader(http.StatusOK)
	case err := <-errCh:
		log.Println("failed to get response from ServiceB:", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

// logger is a middleware for logging
func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, r.Method)
		fn(w, r)
	}
}

func handlerAsync(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var b []byte
	err := hystrix.Do("Dhruv", func() error {
		resp, err := http.Get("http://localhost:9090")
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		b, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return nil
	}, nil)

	if err != nil {
		log.Println("failed to get response from ServiceB:", err)
		w.WriteHeader(http.StatusServiceUnavailable)

		return
	}

	log.Println("Received response from service B:", string(b))
	w.WriteHeader(http.StatusOK)

}
