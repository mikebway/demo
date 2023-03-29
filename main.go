package main

import (
	"log"
	"net/http"
	"os"
)

var (
	name string
)

func main() {

	name = os.Args[1]

	http.HandleFunc("/ping", PingFunc)

	http.HandleFunc("/pong", PongFunc)

	http.HandleFunc("/", HomeFunc)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Unexpected error starting server: %v", err)
	}
}

func PingFunc(writer http.ResponseWriter, request *http.Request) {
	count, err := writer.Write([]byte("Hello " + name))
	if err != nil {
		log.Printf("Unexpected error writing response: %v", err)
	} else {
		log.Printf("Byte count %d\n", count)
	}
}

func PongFunc(writer http.ResponseWriter, request *http.Request) {
	count, err := writer.Write([]byte("Goodbye " + name))
	if err != nil {
		log.Printf("Unexpected error writing response: %v", err)
	} else {
		log.Printf("Byte count %d\n", count)
	}
}

func HomeFunc(writer http.ResponseWriter, request *http.Request) {
	count, err := writer.Write([]byte("Try /ping or /pong %s" + name))
	if err != nil {
		log.Printf("Unexpected error writing response: %v", err)
	} else {
		log.Printf("Byte count %d\n", count)
	}
}
