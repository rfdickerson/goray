package pathtracer

import "math"
import "github.com/rfdickerson/goray/vector"

// Sphere is a thing
type Sphere struct {
	origin vector.Vec4
	radius float64
}

func solveQuadratic(a, b, c float64) (float64, float64) {

	// get discriminant
	d := b*b - 4*a*c

	if d < 0 {
		return 0, 0
	}

	var q float64
	if b > 0 {
		q = -0.5*b + math.Sqrt(d)
	} else {
		q = -0.5*b - math.Sqrt(d)
	}

	x0 := q / a
	x1 := c / q

	return x0, x1
}

// Intersect - get intersection back of Sphere
func (s *Sphere) Intersect(r *Ray) (*Intersection, bool) {

	l := vector.Subtract(&r.origin, &s.origin)
	a := vector.Dot(&r.direction, &r.direction)
	b := 2 * vector.Dot(&r.direction, &l)
	c := vector.Dot(&l, &l) - s.radius*s.radius

	r0, r1 := solveQuadratic(a, b, c)

	t0 := math.Min(r0, r1)
	t1 := math.Max(r0, r1)

	if t0 < 0 && t1 < 0 {
		return nil, false
	}

	i := Intersection{thing: s, r: *r, dist: t0}

	return &i, true
}

// Normal - get normal of sphere
func (s *Sphere) Normal(p vector.Vec4) vector.Vec4 {
	return vector.Subtract(&s.origin, &p)
}

// Surface - gets surface back for the sphere
func (s *Sphere) Surface() Surface {
	return Surface{diffuse: vector.Vec4{255, 255, 0, 255}}
}
