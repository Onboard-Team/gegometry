package r3

// Plane definition of the plane given the normal and the distance
type Plane struct {
	Normal   Vector
	Distance float64
}

func newPlane(normal Vector, origin Vector) Plane {
	return Plane{Normal: normal, Distance: -Dot(normal, origin)}
}
func (p Plane) calculateNormal(a, b, c Vector) Vector {
	U := createVector(b, a)
	V := createVector(c, a)

	return Vector{
		U.Y*V.Z - U.Z*V.Y,
		U.Z*V.X - U.X*V.Z,
		U.X*V.Y - U.Y*V.X,
	}
}

func main() {
	normal := Vector{1, 1, 0}
	origin := Vector{0, 0, 0}

	newPlane(normal, origin)
}
