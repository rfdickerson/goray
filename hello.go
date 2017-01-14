package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type vec4 struct {
	X, Y, Z, W float64
}

func (v *vec4) norm() vec4 {
	mag := math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2) + math.Pow(v.W, 2)
	return vec4{v.X / mag, v.Y / mag, v.Z / mag, v.W / mag}
}

func (v *vec4) Add(b *vec4) {
	v.X += b.X
	v.Y += b.Y
	v.Z += b.Z
	v.W += b.W
}

type ray struct {
	origin    vec4
	direction vec4
}

type collidable interface {
	doesHit(r ray) float64
}

type sphere struct {
	origin vec4
	radius float64
}

func (s *sphere) DoesHit(r *ray) float64 {
	return 20
}

func main() {
	fmt.Print("Raytracing...\n")

	s := sphere{origin: vec4{0, 0, 0, 1}, radius: 0.3}
	r := ray{origin: vec4{0, 0, 0, 1}, direction: vec4{0, 0, 1, 1}}

	depth := s.DoesHit(&r)

	fmt.Printf("Collided at %f", depth)

	newImage := image.NewRGBA(image.Rect(0, 0, 300, 200))

	for i := 0; i < 300; i++ {
		for j := 0; j < 200; j++ {
			r := 255 * float32(i) / float32(300)
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
