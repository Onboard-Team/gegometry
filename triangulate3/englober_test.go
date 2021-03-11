package triangulate3

import (
	"fmt"
	"geGoMetry/r3"
	"geGoMetry/shape"
	"testing"
)

//	testing ComputeBoundingBoxForVectors method
func TestComputeBoundingBoxForVectors(t *testing.T) {
	a := r3.Vector{0, 0, 0}
	b := r3.Vector{0, 1, 0}
	c := r3.Vector{1, 0, 0}
	d := r3.Vector{1, 1, 0}
	e := r3.Vector{0, 9, 1}
	f := r3.Vector{0, 1, 1}
	g := r3.Vector{1, 0, 8}
	h := r3.Vector{1, 1, 1}

	Vectors := []r3.Vector{a, b, c, d, e, f, g, h}
	minExtremity, maxExtremity := computeBoundingBoxForVectors(Vectors)
	fmt.Println(minExtremity, maxExtremity)
}

func testVectorsInsideTetrahedron(Vectors, tetrahedronVertices []r3.Vector) {

	for _, Vector := range Vectors {
		flag := VectorInTetrahedron(tetrahedronVertices[0], tetrahedronVertices[1], tetrahedronVertices[2], tetrahedronVertices[3], Vector)
		if !flag {
			panic("WRAPPING MISSES VectorS")
		}
	}
	fmt.Println("TEST PASSED")
}

func TestCreateEnglobingTetrahedronForVectors(t *testing.T) shape.Mesh {
	a := r3.Vector{0, 0, 0}
	b := r3.Vector{0, 1, 0}
	c := r3.Vector{1, 0, 0}
	d := r3.Vector{1, 1, 0}
	e := r3.Vector{0, 0, 1}
	f := r3.Vector{0, 1, 1}
	g := r3.Vector{1, 0, 1}
	h := r3.Vector{1, 1, 1}

	Vectors := []r3.Vector{a, b, c, d, e, f, g, h}

	englobingTetrahedron := constructEnglobingTetra(Vectors)
	testVectorsInsideTetrahedron(Vectors, englobingTetrahedron.Vertices)
	fmt.Println(englobingTetrahedron)
	return englobingTetrahedron

}

func TestCreateInitialTetra(t *testing.T) {
	a := r3.Vector{0, 0, 0}
	b := r3.Vector{0, 1, 0}
	c := r3.Vector{1, 0, 0}
	d := r3.Vector{1, 1, 0}
	e := r3.Vector{0, 0, 1}
	f := r3.Vector{0, 1, 1}
	g := r3.Vector{1, 0, 1}
	h := r3.Vector{1, 1, 1}

	Vectors := []r3.Vector{a, b, c, d, e, f, g, h}

	initialTetrahedron := createInitialTetra(Vectors)

	testVectorsInsideTetrahedron(Vectors, initialTetrahedron)

	fmt.Println(initialTetrahedron)

}
