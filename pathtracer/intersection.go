package pathtracer

// Intersection is the intersection information of ray object collision
type Intersection struct {
	thing Thing
	r     Ray
	dist  float64
}
