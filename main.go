package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

//Toy interface to load a few image files.
func main() {

	argsWithoutProg := os.Args[1:]

	var images = make([]image.Image, 4)
	var gsimgs = make([]image.Gray16, 4)

	for i, path := range argsWithoutProg {
		var img, err = getImageFromFilePath(path)
		if err != nil {
			panic(err)
		} else {
			images[i] = img
			gsimgs[i] = *(imageToGrayscale(img))
		}
	}

	fmt.Println(images)

}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, format, err := image.Decode(f)

	fmt.Printf("Image loaded from %s with format %s\n", filePath, format)
	return image, err
}

func imageToGrayscale(img image.Image) *image.Gray16 {
	// Converting image to grayscale
	grayImg := image.NewGray16(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}
	return grayImg
}
