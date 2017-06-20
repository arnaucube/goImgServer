package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! To send images, go to /image")
}

func ImageShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageName := vars["imageName"]
	//fmt.Fprintln(w, "Image show:", imageName)
	var imageExtension = strings.Split(imageName, ".")[1]

	file, err := os.Open(imageName)
	if err != nil {
		//log.Fatal(err)
		//fmt.Fprintln(w, "la imatge no existeix al server")
		fmt.Fprintln(w, err)
	}

	var img image.Image
	switch imageExtension {
	case "png":
		img, err = png.Decode(file)
	case "jpg":
		img, err = jpeg.Decode(file)
	case "jpeg":
		img, err = jpeg.Decode(file)
	default:
		img = nil
	}

	if err != nil {
		//log.Fatal(err)
		fmt.Fprintln(w, "la imatge no existeix al server")
	} else {
		file.Close()

		jpeg.Encode(w, img, nil) // Write to the ResponseWriter
	}
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

	fmt.Fprintln(w, "url:", handler.Filename)
}
