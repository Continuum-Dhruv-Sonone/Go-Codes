package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	setUpRoutes()
	http.ListenAndServe(":8080", nil)
}

func setUpRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func reader(conn *websocket.Conn) {

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println(err)
		}
	}

}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {

	//Dirty way of checking the origin of request
	//Allows request from any source
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	wsConnection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading the websocket connection: %v", err)
	}

	fmt.Println("Successfull client websocket connection.")

	reader(wsConnection)
}
