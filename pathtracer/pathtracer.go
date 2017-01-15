package pathtracer

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/rfdickerson/goray/vector"
)

var sceneObjects [1]Thing

const width = 300
const height = 300

func castRay(r *Ray) color.RGBA {

	for _, s := range sceneObjects {
		i, hits := s.Intersect(r)
		if hits == true {
			// fmt.Printf("Ray collided at %f", i.dist)

			normal := i.dist
			return color.RGBA{uint8(normal * 255), 0, 255, 255}
		}
	}

	return color.RGBA{255, 0, 255, 255}
}

// StartPathtracing pathtracing
func StartPathtracing() {
	fmt.Print("Raytracing...\n")

	s := Sphere{origin: vector.NewVector(0, 0, 0), radius: 1.0}
	sceneObjects[0] = s

	newImage := image.NewRGBA(image.Rect(0, 0, 300, 200))

	origin := vector.NewVector(0, 0, -10)

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {

			direction := vector.Vec4{X: -1 + 2*(float64(i)/width), Y: -1 + 2*(float64(j)/height), Z: 1, W: 0}
			direction = direction.Norm()

			ray := Ray{origin: origin, direction: direction}

			newPixel := castRay(&ray)
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
