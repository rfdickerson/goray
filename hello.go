package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	fmt.Print("Raytracing...")

	newImage := image.NewRGBA(image.Rect(0, 0, 300, 200))

	for i := 0; i < 300; i++ {
		for j := 0; j < 200; j++ {
			r := 255 * float64(i) / float64(300)
			newPixel := color.RGBA{uint8(r), 255, 0, 255}
			newImage.SetRGBA(i, j, newPixel)
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, newImage)
}
