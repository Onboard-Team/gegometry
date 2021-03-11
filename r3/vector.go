package r3

import (
	"math"
)

// Vector data structure
type Vector struct {
	X, Y, Z float64
}

// Zero Vector
var Zero = Vector{}

func createVector(a, b Vector) Vector {
	return Vector{
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
	}
}

// Norm returns the vector's norm.
func (v Vector) Norm() float64 { return math.Sqrt(float64(Dot(v, v))) }

// Norm2 returns the square of the norm.
func (v Vector) Norm2() float64 { return Dot(v, v) }

// Normalize updates the vector to a unit vector in the same direction as v.
func (v *Vector) Normalize() *Vector {
	n2 := v.Norm2()
	if n2 == 0 {
		v.X, v.Y, v.Z = 0, 0, 0
	}
	v.Scale(1 / math.Sqrt(n2))
	return v
}

// Normalize returns a unit vector in the same direction as v.
func Normalize(v Vector) Vector {
	n2 := v.Norm2()
	if n2 == 0 {
		return Vector{0, 0, 0}
	}
	return v.Scaled(1 / math.Sqrt(n2))
}

// IsUnit returns whether this vector is of approximately unit length.
func (v Vector) IsUnit() bool {
	const epsilon = 5e-14
	return math.Abs(float64(v.Norm2()-1)) <= epsilon
}

// Abs returns the vector with nonnegative components.
func (v *Vector) Abs() {
	v.X = math.Abs(v.X)
	v.Y = math.Abs(v.Y)
	v.Z = math.Abs(v.Z)
}

func Abs(v Vector) Vector {
	return Vector{
		math.Abs(v.X), math.Abs(v.Y), math.Abs(v.Z)}

}

// Scale multiplies all element of the vector by f and returns vec.
func (v *Vector) Scale(f float64) *Vector {
	v.X = f * v.X
	v.Y = f * v.Y
	v.Z = f * v.Z
	return v
}

// Scaled returns a copy of vec with all elements multiplies by f.
func (v Vector) Scaled(f float64) Vector { return Vector{f * v.X, f * v.Y, f * v.Z} }

// Add updated the vector to the standard vector sum of v and ov.
func (v *Vector) Add(ov Vector) {
	v.X = v.X + ov.X
	v.Y = v.Y + ov.Y
	v.Z = v.Z + ov.Z

}

// Add returns the standard vector sum of v and ov
func Add(v, ov Vector) Vector {
	return Vector{v.X + ov.X, v.Y + ov.Y, v.Z + ov.Z}
}

// Sub updates the vector to the standard vector difference of v and ov.
func (v *Vector) Sub(ov Vector) {
	v.X = v.X - ov.X
	v.Y = v.Y - ov.Y
	v.Z = v.Z - ov.Z
}

// Sub returns the standard vector difference of v and ov.
func Sub(v, ov Vector) Vector {
	return Vector{v.X - ov.X, v.Y - ov.Y, v.Z - ov.Z}
}

// Dot returns the standard dot product of v and ov.
func Dot(v, ov Vector) float64 { return v.X*ov.X + v.Y*ov.Y + v.Z*ov.Z }

// Cross returns the cross product of
func Cross(v Vector, ov Vector) Vector {
	return Vector{v.Y*ov.Z - ov.Y*v.Z, v.Z*ov.X - ov.Z*v.X, v.X*ov.Y - ov.X*v.Y}
}

// CalculateDistanceToPlane returns the distance from vector v to plane p
func (v Vector) CalculateDistanceToPlane(p Plane) float64 { return Dot(p.Normal, v) + p.Normal.Norm2() }

func det2D(a, b Vector) float64 {
	return a.X*b.Y - a.Y*b.X
}

// GetInterpolationVector returns the interpolated vector between v and b given a certain divisor
func (v Vector) GetInterpolationVector(b Vector, divisor float64) Vector {

	divisionVector := Vector{
		v.X*divisor + b.X*(1-divisor),
		v.Y*divisor + b.Y*(1-divisor),
		v.Z*divisor + b.Z*(1-divisor),
	}

	return divisionVector
}

// Magnitude returns the magnitude of a vector
func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.SquaredMagnitude())
}

// SquaredMagnitude returns the squared magnitude of a vector
func (v Vector) SquaredMagnitude() float64 {
	return float64(Dot(v, v))
}

// Distance returns the distance between v and ov
func Distance(v, ov Vector) float64 {
	diff := Sub(v, ov)
	return diff.Norm()
}

// SquaredDistance returns the distance between v and b
func SquaredDistance(v, ov Vector) float64 {
	diff := Sub(v, ov)
	return diff.Norm2()
}

// Interpolate interpolates between a and b at t (0,1).
func Interpolate(a, b *Vector, t float64) Vector {
	t1 := 1 - t
	return Vector{
		a.X*t1 + b.X*t,
		a.Y*t1 + b.Y*t,
		a.Z*t1 + b.Z*t,
	}
}
