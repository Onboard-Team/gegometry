package modeler

import (
	"geGoMetry/glTF"
	"geGoMetry/r3"
	"geGoMetry/shape"

	"testing"
)

func TestConstructSphere(t *testing.T) {
	center := r3.Vector{0, 0, 0}
	radius := 1.0
	sphere := shape.ConstructSphere(center, radius)
	outputFile := "../testfiles/exports/sphere.glb"
	glTF.Export(sphere, outputFile)
}

func TestConstructCircle(t *testing.T) {
	center := r3.Vector{0, 0, 0}
	radius := 1.0
	circle := shape.ConstructCircle(center, radius)
	outputFile := "../testfiles/exports/circle.glb"
	glTF.Export(circle, outputFile)
}

func TestConstructTorus(t *testing.T) {
	center := r3.Vector{0, 0, 0}
	radius := 10.0
	tubeRadius := 4.0

	circle := shape.ConstructTorus(center, radius, tubeRadius)
	outputFile := "../testfiles/exports/torus.glb"
	glTF.Export(circle, outputFile)
}

func TestConstructEllipse(t *testing.T) {
	center := r3.Vector{0, 0, 0}
	semiAxis1 := 10.0
	semiAxis2 := 4.0

	circle := shape.ConstructEllipse(center, semiAxis1, semiAxis2)
	outputFile := "../testfiles/exports/ellipse.glb"
	glTF.Export(circle, outputFile)
}
