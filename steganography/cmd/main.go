package main

import (
	"log"
	"os"
)

func main() {
	// myImage := image.NewRGBA(image.Rect(0, 0, 100, 200))
	// for p := 0; p < 100*200; p++ {
	// 	pixelOffset := 4 * p
	// 	myImage.Pix[0+pixelOffset] = uint8(rand.Intn(256))
	// 	myImage.Pix[1+pixelOffset] = uint8(rand.Intn(256))
	// 	myImage.Pix[2+pixelOffset] = uint8(rand.Intn(256))
	// 	myImage.Pix[3+pixelOffset] = uint8(rand.Intn(256))
	// }

	// outputFile, err := os.Create("test.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// jpeg.Encode(outputFile, myImage, nil)
	// err = outputFile.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// zip.Zip()
	firstFile, err := os.Open("test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer firstFile.Close()
	secondFile, err := os.Open("test.zip")
}
