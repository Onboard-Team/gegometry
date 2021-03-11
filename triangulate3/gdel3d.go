//	an implementation of the gDel3d research paper
//	intended to be working on both a cpu and a gpu
package triangulate3

import (
	"geGoMetry/r3"
	"math"
)

func createInitialTetra(Vectors []r3.Vector) []r3.Vector {
	v0, v1 := computeExtremitiesForAxeX(Vectors)
	// Find the furthest Vector from v0v1
	v2 := getFurthestFromTwoVectors(Vectors, v0, v1)
	// Find the furthest Vector from v0v1v2
	v3 := getFurthestFromthreeVectors(Vectors, v0, v1, v2)

	vertices := []r3.Vector{v0, v1, v2, v3}
	return vertices
}
func getFurthestFromTwoVectors(Vectors []r3.Vector, a, b r3.Vector) r3.Vector {
	max := 0.0
	var furthestVector r3.Vector
	for _, Vector := range Vectors {
		dist := Get2Ddist(a, b, Vector)
		if max < dist {
			furthestVector = Vector
			max = dist
		}
	}
	return furthestVector
}

func getFurthestFromthreeVectors(Vectors []r3.Vector, a, b, c r3.Vector) r3.Vector {
	max := 0.0
	var furthestVector r3.Vector
	for _, Vector := range Vectors {
		dist := Get3Ddist(a, b, c, Vector)
		if max > dist {
			furthestVector = Vector
			max = dist
		}
	}
	return furthestVector
}

func computeExtremitiesForAxeX(Vectors []r3.Vector) (r3.Vector, r3.Vector) {
	minExtremity := r3.Vector{math.MaxFloat32, math.MaxFloat32, math.MaxFloat32}
	maxExtremity := r3.Vector{-math.MaxFloat32, -math.MaxFloat32, -math.MaxFloat32}
	for _, Vector := range Vectors {
		//Updating the minExtremity
		if Vector.X < minExtremity.X {
			minExtremity = Vector
		}

		//Updating the maxExtremity
		if Vector.X > maxExtremity.X {
			maxExtremity = Vector
		}

	}

	return minExtremity, maxExtremity
}
