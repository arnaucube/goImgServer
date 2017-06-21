package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! To send images, go to /image")
}

func ImageShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageName := vars["imageName"]

	file, err := ioutil.ReadFile(config.Folder + "/" + imageName)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	img, err := dataToImage(file, imageName)

	if err != nil {
		fmt.Fprintln(w, "la imatge no existeix al server")
	} else {
		jpeg.Encode(w, img, nil) // Write to the ResponseWriter
	}
}

func NewImage(w http.ResponseWriter, r *http.Request) {
	err := ipFilter(r)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	imageName := strings.Split(handler.Filename, ".")[0]
	fileName := imageName + ".png"
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	img, err := dataToImage(data, handler.Filename)
	if err != nil {
		fmt.Fprintln(w, "error al processar la imatge")
	}
	img = Resize(img)
	data, err = imageToPNG(img)
	if err != nil {
		fmt.Fprintln(w, "error al processar la imatge")
	}
	err = ioutil.WriteFile(config.Folder+"/"+fileName, data, 0777)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w, "http://"+config.ServerIP+":"+config.ServerPort+"/images/"+fileName)
}
