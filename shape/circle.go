package shape

import (
	"geGoMetry/r3"
	"math"
)

type Circle struct {
	center      r3.Vector
	radius      float64
	sectorCount int
}

// Construct : constructs the circle.
func (c Circle) Construct() Mesh {
	circle := Mesh{}

	c.generateVertices(&circle)
	c.generateIndices(&circle)

	return circle
}

func (c Circle) generateVertices(circle *Mesh) {
	sectorStep := 2 * math.Pi / float64(c.sectorCount)

	circle.Vertices = append(circle.Vertices, c.center)

	for i := 0; i < c.sectorCount; i++ {

		sectorAngle := math.Pi/2 - float64(i)*sectorStep // starting from pi/2 to -pi/2

		x := c.radius * math.Cos(sectorAngle)
		y := c.radius * math.Sin(sectorAngle)
		z := 0.

		vertex := r3.Vector{x, y, z}
		vertex.Add(c.center)

		circle.Vertices = append(circle.Vertices, vertex)

	}
}

func (c Circle) generateIndices(circle *Mesh) {
	lenVertices := len(circle.Vertices)
	for i := 1; i < lenVertices; i++ {
		circle.Indices = append(circle.Indices, 0, uint32(i), uint32(i+1))
		circle.Indices = append(circle.Indices, uint32(i+1), uint32(i), 0)

	}
	circle.Indices = append(circle.Indices, 0, uint32(lenVertices-1), 1)
	circle.Indices = append(circle.Indices, 1, uint32(lenVertices-1), 0)

}
