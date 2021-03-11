package r3

type Line struct {
	Vector    Vector
	Direction Vector
}

type Plane3D struct {
	Normal Vector
	Vector Vector
}

func (line Line) IntersectPlane(plane Plane3D) Vector {
	if Dot(plane.Normal, Normalize(line.Direction)) == 0 {
		panic("plane and line are parallel. Therefore, no intersection")
	}

	t := (Dot(plane.Normal, plane.Vector) - Dot(plane.Normal, line.Vector)) / (Dot(plane.Normal, Normalize(line.Direction)))
	return Add(line.Vector, Normalize(line.Direction).Scaled(t))
}
