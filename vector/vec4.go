package vector

import (
	"math"
)

// Vec4 is a basic vector
type Vec4 struct {
	X, Y, Z, W float64
}

// NewVector returns a vector object
func NewVector(x, y, z float64) Vec4 {
	v := Vec4{X: x, Y: y, Z: z, W: 1}
	return v
}

// Norm - Normalize the vector
func (v *Vec4) Norm() Vec4 {
	mag := math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2) + math.Pow(v.W, 2)
	return Vec4{v.X / mag, v.Y / mag, v.Z / mag, v.W / mag}
}

// Add - adds two vectors
func Add(a, b *Vec4) Vec4 {
	return Vec4{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}
}

// Subtract - subtracts two vectors
func Subtract(a, b *Vec4) Vec4 {
	return Vec4{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}
}

// Dot - dot product of two vectors
func Dot(a, b *Vec4) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

// Multiply - multiplies a scalar by a vector
func MultiplyS(a float64, b *Vec4) Vec4 {
	return Vec4{a * b.X, a * b.Y, a * b.Z, a * b.W}
}

// Multiply - multiplies a scalar by a vector
func Multiply(a *Vec4, b *Vec4) Vec4 {
	return Vec4{a.X * b.X, a.Y * b.Y, a.Z * b.Z, a.W * b.W}
}
