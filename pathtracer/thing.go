package pathtracer

import "github.com/rfdickerson/goray/vector"

// Thing is something that can be added to a scene
type Thing interface {
	Intersect(r *Ray) (*Intersection, bool)
	Normal(p vector.Vec4) vector.Vec4
	Surface() Surface
}
