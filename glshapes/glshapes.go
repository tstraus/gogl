package glshapes

import (
	"../glutil"
)

type Vec2f struct {
	X float32
	Y float32
}

// Shape is a container of the count and the form of the polygons
type Shape struct {
	PolygonCount int32
	Vao uint32
}

func NewQuadrilateral(topLeft, topRight, bottomLeft, bottomRight Vec2f) *Shape {
	s := new(Shape)

	polygons := []float32 {
		topLeft.X, topLeft.Y, 0, // top left polYgon
		bottomLeft.X, bottomLeft.Y, 0,
		topRight.X, topRight.Y, 0,
		topRight.X, topRight.Y, 0, // bottom right polYgon
		bottomLeft.X, bottomLeft.Y, 0,
		bottomRight.X, bottomRight.Y, 0,
	}

	s.Vao = glutil.MakeVao(polygons)
	s.PolygonCount = int32(len(polygons) / 3)

	return s
}

func NewRectangle(bottomLeft, topRight Vec2f) *Shape {
	return NewQuadrilateral(Vec2f{bottomLeft.X, topRight.Y}, topRight, bottomLeft, Vec2f{topRight.X, bottomLeft.Y})
}