package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"strings"

	"github.com/nfnt/resize"
)

func dataToImage(data []byte, imageName string) (image.Image, error) {
	reader := bytes.NewReader(data)
	var imageExtension = strings.Split(imageName, ".")[1]
	var img image.Image
	var err error
	switch imageExtension {
	case "png":
		img, err = png.Decode(reader)
	case "jpg":
		img, err = jpeg.Decode(reader)
	case "jpeg":
		img, err = jpeg.Decode(reader)
	default:
		img = nil
	}
	if err != nil {
		return img, err
	}
	return img, err
}

func imageToData(img image.Image, imageName string) ([]byte, error) {
	buf := new(bytes.Buffer)
	var imageExtension = strings.Split(imageName, ".")[1]
	var err error
	switch imageExtension {
	case "png":
		err = png.Encode(buf, img)
	case "jpg":
		err = jpeg.Encode(buf, img, nil)
	case "jpeg":
		err = jpeg.Encode(buf, img, nil)
	default:
		img = nil
	}
	if err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), err
}

func imageToPNG(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	var err error
	err = png.Encode(buf, img)
	return buf.Bytes(), err
}

func Resize(img image.Image) image.Image {
	r := resize.Resize(160, 0, img, resize.Lanczos3)
	return r
}
