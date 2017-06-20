package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! To send images, go to /image")
}

func ImageShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageName := vars["imageName"]
	fmt.Fprintln(w, "Image show:", imageName)
}

func NewImage(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(handler.Filename, data, 0777)
	if err != nil {
		fmt.Println(err)
	}

}
