package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	readConfig("config.json")
	fmt.Println(config)

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":3050", router))
}
