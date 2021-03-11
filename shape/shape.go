package shape

import (
	"geGoMetry/r3"
)

// ConstructSphere constructs a sphere given the center and
// the radius.
func ConstructSphere(center r3.Vector, radius float64) Mesh {

	sectorCount := 300
	stackCount := 300

	s := Sphere{center, radius, sectorCount, stackCount}
	sphere := s.Construct()

	return sphere
}

// ConstructCircle constructs a circle given the center and
// the radius
func ConstructCircle(center r3.Vector, radius float64) Mesh {
	sectorCount := 100

	c := Circle{center, radius, sectorCount}
	circle := c.Construct()
	return circle
}

// ConstructTorus constructs a torus given the center and
// the rediee
func ConstructTorus(center r3.Vector, radius, tubeRadius float64) Mesh {
	sectorCount := 100
	stackCount := 100
	t := Torus{center, radius, tubeRadius, sectorCount, stackCount}
	torus := t.Construct()

	return torus
}

// ConstructEllipse constructs an ellipse.
func ConstructEllipse(center r3.Vector, semiAxis1, semiAxis2 float64) Mesh {
	sectorCount := 100
	t := Ellipse{center, semiAxis1, semiAxis2, sectorCount}
	torus := t.Construct()

	return torus
}
