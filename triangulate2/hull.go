package triangulate2

import (
	"geGoMetry/r3"
	"sort"
)

func cross2D(p, a, b r3.Vector) float64 {
	return (a.X-p.X)*(b.Y-p.Y) - (a.Y-p.Y)*(b.X-p.X)
}

// ConvexHull returns the convex hull of the provided Vectors.
func ConvexHull(Vectors []r3.Vector) []r3.Vector {
	// copy Vectors
	VectorsCopy := make([]r3.Vector, len(Vectors))
	copy(VectorsCopy, Vectors)
	Vectors = VectorsCopy

	// sort Vectors
	sort.Slice(Vectors, func(i, j int) bool {
		a := Vectors[i]
		b := Vectors[j]
		if a.X != b.X {
			return a.X < b.X
		}
		return a.Y < b.Y
	})

	// filter nearly-duplicate Vectors
	distinctVectors := Vectors[:0]
	for i, p := range Vectors {
		if i > 0 && r3.SquaredDistance(p, Vectors[i-1]) < eps {
			continue
		}
		distinctVectors = append(distinctVectors, p)
	}
	Vectors = distinctVectors

	// find upper and lower portions
	var U, L []r3.Vector
	for _, p := range Vectors {
		for len(U) > 1 && cross2D(U[len(U)-2], U[len(U)-1], p) > 0 {
			U = U[:len(U)-1]
		}
		for len(L) > 1 && cross2D(L[len(L)-2], L[len(L)-1], p) < 0 {
			L = L[:len(L)-1]
		}
		U = append(U, p)
		L = append(L, p)
	}

	// reverse upper portion
	for i, j := 0, len(U)-1; i < j; i, j = i+1, j-1 {
		U[i], U[j] = U[j], U[i]
	}

	// construct complete hull
	if len(U) > 0 {
		U = U[:len(U)-1]
	}
	if len(L) > 0 {
		L = L[:len(L)-1]
	}
	return append(L, U...)
}
