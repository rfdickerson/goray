package pathtracer

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/rfdickerson/goray/vector"
)

var sceneObjects []Thing

// Intersection is the intersection information of ray object collision
type Intersection struct {
	thing Thing
	r     Ray
	dist  float64
}

func castRay(r *Ray) color.RGBA {
	return color.RGBA{255, 0, 255, 255}
}

func main() {
	fmt.Print("Raytracing...\n")

	s := Sphere{origin: vector.Vec4{0, 0, 0, 1}, radius: 0.3}
	ray := Ray{origin: vector.Vec4{0, 0, -5, 1}, direction: vector.Vec4{0, 0, 1, 1}}

	i, hits := s.Intersect(&ray)
	if hits == true {
		fmt.Printf("Ray collided at %f", i.dist)
	}

	newImage := image.NewRGBA(image.Rect(0, 0, 300, 200))

	for i := 0; i < 300; i++ {
		for j := 0; j < 200; j++ {
			// r := 255 * float32(i) / float32(300)
			// newPixel := color.RGBA{uint8(r), 255, 0, 255}
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
