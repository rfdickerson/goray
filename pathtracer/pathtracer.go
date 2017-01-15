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

			m := vector.MultiplyS(i.dist, &r.direction)
			intersectionPoint := vector.Add(&r.origin, &m)
			normal := i.thing.Normal(intersectionPoint)

			fmt.Printf("normal is %f, %f, %f\n", normal.X, normal.Y, normal.Z)

			return color.RGBA{uint8(normal.X * 255), uint8(normal.Y * 255), uint8(normal.Z * 255), 255}
		}
	}

	return color.RGBA{255, 0, 255, 255}
}

// StartPathtracing pathtracing
func StartPathtracing() {
	fmt.Print("Raytracing...\n")

	s := Sphere{origin: vector.NewVector(0, 0, 0), radius: 1.0}
	sceneObjects[0] = s

	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

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
