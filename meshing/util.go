package mesh

import "geGoMetry/r3"

func findCentroid(p0, p1, p2, p3 r3.Vector) r3.Vector {
	x := (p0.X + p1.X + p2.X + p3.X) / 4.0
	y := (p0.Y + p1.Y + p2.Y + p3.Y) / 4.0
	z := (p0.Z + p1.Z + p2.Z + p3.Z) / 4.0
	return r3.Vector{X: x, Y: y, Z: z}
}

func sameSide(v1, v2, v3, v4, p r3.Vector) bool {
	normal := r3.Cross(r3.Sub(v2, v1), r3.Sub(v3, v1))
	dotV4 := r3.Dot(normal, r3.Sub(v4, v1))
	dotP := r3.Dot(normal, r3.Sub(p, v1))
	if dotV4 <= 0 && dotP <= 0 {
		return true
	}
	if dotV4 >= 0 && dotP >= 0 {
		return true
	}

	return false
}

func vectorInTetrahedron(v1, v2, v3, v4, p r3.Vector) bool {
	return sameSide(v1, v2, v3, v4, p) &&
		sameSide(v2, v3, v4, v1, p) &&
		sameSide(v3, v4, v1, v2, p) &&
		sameSide(v4, v1, v2, v3, p)
}

func duplicateInReverse(indices *[]uint32) {
	for i := len(*indices) - 1; i >= 0; i-- {
		*indices = append(*indices, (*indices)[i])
	}
}
