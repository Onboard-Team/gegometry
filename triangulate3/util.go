package triangulate3

import "geGoMetry/r3"

func SameSide(v1, v2, v3, v4, p r3.Vector) bool {
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

func VectorInTetrahedron(v1, v2, v3, v4, p r3.Vector) bool {
	return SameSide(v1, v2, v3, v4, p) &&
		SameSide(v2, v3, v4, v1, p) &&
		SameSide(v3, v4, v1, v2, p) &&
		SameSide(v4, v1, v2, v3, p)
}

func findCentroid(p0, p1, p2, p3 r3.Vector) r3.Vector {
	x := (p0.X + p1.X + p2.X + p3.X) / 4.0
	y := (p0.Y + p1.Y + p2.Y + p3.Y) / 4.0
	z := (p0.Z + p1.Z + p2.Z + p3.Z) / 4.0
	return r3.Vector{x, y, z}
}

func Get2Ddist(a, b, c r3.Vector) float64 {

	abx := b.X - a.X
	aby := b.Y - a.Y
	abz := b.Z - a.Z

	acx := c.X - a.X
	acy := c.Y - a.Y
	acz := c.Z - a.Z

	xy := abx*acy - aby*acx
	yz := aby*acz - abz*acy
	zx := abz*acx - abx*acz

	dist := xy*xy + yz*yz + zx*zx

	return dist
}

func Get3Ddist(a, b, c, d r3.Vector) float64 {
	abx := b.X - a.X
	aby := b.Y - a.Y
	abz := b.Z - a.Z
	acx := c.X - a.X
	acy := c.Y - a.Y
	acz := c.Z - a.Z

	adx := d.X - a.X
	ady := d.Y - a.Y
	adz := d.Z - a.Z

	bc := abx*acy - aby*acx
	cd := acx*ady - acy*adx
	db := adx*aby - ady*abx

	dist := abz*cd + acz*db + adz*bc

	return dist
}
