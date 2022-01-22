package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"image"
	_ "image/jpeg"
	_ "image/png"
)

var imgPath = flag.String("file", "./go-logo.png", "The image file to be read")

func main() {
	flag.Parse()
	fmt.Printf("Displaying data for file %s\n", *imgPath)
	reader, err := os.Open(*imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	img, imgType, err := image.Decode(reader)
	fmt.Printf("Image type is %s", imgType)
	if err != nil {
		fmt.Printf("Error opening the image file %s", *imgPath)
		os.Exit(-1)
	}
	rect := img.Bounds()
	for i := rect.Min.X; i < rect.Max.X; i++ {
		for j := rect.Min.Y; j < rect.Max.Y; j++ {
			r, g, b, a := img.At(i, j).RGBA()
			fmt.Printf("X: %d, Y: %d >> Red: %d, Green: %d, Blue: %d Alpha: %d, ", i, j, r, g, b, a)
		}
	}
}
