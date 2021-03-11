package triangulate2

import (
	"fmt"
	"geGoMetry/r3"
	"math"
)

type Triangulation struct {
	Vectors    []r3.Vector
	ConvexHull []r3.Vector
	Triangles  []int
	Halfedges  []int
}

// Triangulate returns a Delaunay triangulation of the provided Vectors.
func Triangulate(Vectors []r3.Vector) (*Triangulation, error) {
	t := newTriangulator(Vectors)
	err := t.triangulate()
	return &Triangulation{Vectors, t.convexHull(), t.triangles, t.halfedges}, err
}

func (t *Triangulation) area() float64 {
	var result float64
	Vectors := t.Vectors
	ts := t.Triangles
	for i := 0; i < len(ts); i += 3 {
		p0 := Vectors[ts[i+0]]
		p1 := Vectors[ts[i+1]]
		p2 := Vectors[ts[i+2]]
		result += area(p0, p1, p2)
	}
	return result / 2
}

// Validate performs several sanity checks on the Triangulation to check for
// potential errors. Returns nil if no issues were found. You normally
// shouldn't need to call this function but it can be useful for debugging.
func (t *Triangulation) Validate() error {
	// verify halfedges
	for i1, i2 := range t.Halfedges {
		if i1 != -1 && t.Halfedges[i1] != i2 {
			return fmt.Errorf("invalid halfedge connection")
		}
		if i2 != -1 && t.Halfedges[i2] != i1 {
			return fmt.Errorf("invalid halfedge connection")
		}
	}

	// verify convex hull area 	 sum of triangle areas
	hull1 := t.ConvexHull
	hull2 := ConvexHull(t.Vectors)
	area1 := polygonArea(hull1)
	area2 := polygonArea(hull2)
	area3 := t.area()
	if math.Abs(area1-area2) > 1e-9 || math.Abs(area1-area3) > 1e-9 {
		return fmt.Errorf("hull areas disagree: %f, %f, %f", area1, area2, area3)
	}

	// verify convex hull perimeter
	perimeter1 := polygonPerimeter(hull1)
	perimeter2 := polygonPerimeter(hull2)
	if math.Abs(perimeter1-perimeter2) > 1e-9 {
		return fmt.Errorf("hull perimeters disagree: %f, %f", perimeter1, perimeter2)
	}

	return nil
}
