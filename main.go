package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	readConfig("config.json")
	fmt.Printf("%+v\n", config)

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":"+config.ServerPort, router))
}
