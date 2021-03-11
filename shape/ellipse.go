package shape

import (
	"geGoMetry/r3"
	"math"
)

// Ellipse : constructs a UV Ellipse.
type Ellipse struct {
	center      r3.Vector
	semiAxis1   float64
	semiAxis2   float64
	sectorCount int
}

// Construct : constructs the ellipse
func (e Ellipse) Construct() Mesh {
	ellipse := Mesh{}

	e.generateVertices(&ellipse)
	e.generateIndices(&ellipse)

	return ellipse
}

func (e Ellipse) generateVertices(circle *Mesh) {
	sectorStep := 2 * math.Pi / float64(e.sectorCount)

	circle.Vertices = append(circle.Vertices, e.center)

	for i := 0; i < e.sectorCount; i++ {

		sectorAngle := math.Pi/2 - float64(i)*sectorStep // starting from pi/2 to -pi/2

		x := e.semiAxis1 * math.Cos(sectorAngle)
		y := e.semiAxis2 * math.Sin(sectorAngle)
		z := 0.

		vertex := r3.Vector{x, y, z}
		vertex.Add(e.center)

		circle.Vertices = append(circle.Vertices, vertex)

	}
}

func (e Ellipse) generateIndices(circle *Mesh) {
	lenVertices := len(circle.Vertices)
	for i := 1; i < lenVertices; i++ {
		circle.Indices = append(circle.Indices, 0, uint32(i), uint32(i+1))
		circle.Indices = append(circle.Indices, uint32(i+1), uint32(i), 0)

	}
	circle.Indices = append(circle.Indices, 0, uint32(lenVertices-1), 1)
	circle.Indices = append(circle.Indices, 1, uint32(lenVertices-1), 0)

}
