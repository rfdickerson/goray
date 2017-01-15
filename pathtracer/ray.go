package pathtracer

import "github.com/rfdickerson/goray/vector"

// Ray is a ray
type Ray struct {
	origin    vector.Vec4
	direction vector.Vec4
}
