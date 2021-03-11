package triangulate3

import (
	"geGoMetry/r3"
	"geGoMetry/shape"
	"math"
)

func constructEnglobingTetra(Vectors []r3.Vector) shape.Mesh {
	offset := 1.0

	minExtremity, maxExtremity := computeBoundingBoxForVectors(Vectors)

	minExtremity.Add(r3.Vector{-offset, -offset, -offset})
	maxExtremity.Add(r3.Vector{offset, offset, offset})

	//	get bottom plane
	normal := r3.Vector{0, 1, 0}
	lowerPlane := r3.Plane3D{Normal: normal, Vector: minExtremity}
	top := getTopVectorOfTetrahedron(minExtremity, maxExtremity)
	//Line 1
	upperVector1 := r3.Vector{minExtremity.X, maxExtremity.Y, maxExtremity.Z}
	Direct1 := r3.Sub(top, upperVector1)
	line1 := r3.Line{Vector: upperVector1, Direction: Direct1}
	// first bottom Intersection Vector
	tetraPoint1 := line1.IntersectPlane(lowerPlane)

	//Line 2
	upperVector2 := r3.Vector{maxExtremity.X, maxExtremity.Y, maxExtremity.Z}
	Direct2 := r3.Sub(top, upperVector2)
	line2 := r3.Line{Vector: upperVector2, Direction: Direct2}
	// second bottom Intersection Vector
	tetraPoint2 := line2.IntersectPlane(lowerPlane)

	//Line 3
	upperVector3 := r3.Vector{maxExtremity.X, maxExtremity.Y, minExtremity.Z}
	Direct3 := r3.Sub(top, upperVector3)
	line3 := r3.Line{Vector: upperVector3, Direction: Direct3}
	// third bottom Intersection Vector
	tetraPoint3 := line3.IntersectPlane(lowerPlane)

	indices := []uint32{0, 1, 2,
		0, 2, 3,
		0, 3, 1,
		1, 2, 3}

	vertices := []r3.Vector{top, tetraPoint1, tetraPoint2, tetraPoint3}
	mesh := shape.Mesh{Vertices: vertices, Indices: indices}

	return mesh
}
func getTopVectorOfTetrahedron(minExtremity, maxExtremity r3.Vector) r3.Vector {
	x := minExtremity.X
	y := 2 * maxExtremity.Y
	z := minExtremity.Z

	return r3.Vector{x, y, z}
}

func computeBoundingBoxForVectors(Vectors []r3.Vector) (r3.Vector, r3.Vector) {
	minExtremity := r3.Vector{math.MaxFloat32, math.MaxFloat32, math.MaxFloat32}
	maxExtremity := r3.Vector{-math.MaxFloat32, -math.MaxFloat32, -math.MaxFloat32}
	for _, Vector := range Vectors {
		//Updating the minExtremity
		if Vector.X < minExtremity.X {
			minExtremity.X = Vector.X
		}
		if Vector.Y < minExtremity.Y {
			minExtremity.Y = Vector.Y
		}
		if Vector.Z < minExtremity.Z {
			minExtremity.Z = Vector.Z
		}

		//Updating the maxExtremity
		if Vector.X > maxExtremity.X {
			maxExtremity.X = Vector.X
		}
		if Vector.Y > maxExtremity.Y {
			maxExtremity.Y = Vector.Y
		}
		if Vector.Z > maxExtremity.Z {
			maxExtremity.Z = Vector.Z
		}
	}

	return minExtremity, maxExtremity
}
