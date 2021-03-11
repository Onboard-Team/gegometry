package r3

import (
	"fmt"
	"testing"
)

func TestIntersection(t *testing.T) {
	point := Vector{1, 2, 3}
	direction := Vector{0, 1, 0}
	line := Line{Vector: point, Direction: direction}

	normal := Vector{0, 1, 0}
	planeVector := Vector{0, 0, 0}
	plane := Plane3D{normal, planeVector}

	//Should be {1,0,3}
	solution := Vector{1, 0, 3}
	intersectionVector := line.IntersectPlane(plane)
	fmt.Println(intersectionVector)
	if intersectionVector != solution {
		t.Errorf("Should be {1,0,3} got", intersectionVector)
	}

	//Another test
	point = Vector{1, 2, 3}
	direction = Vector{1, 1, 0}
	line = Line{Vector: point, Direction: direction}

	normal = Vector{0, 1, 0}
	planeVector = Vector{0, 0, 0}
	plane = Plane3D{normal, planeVector}

	//Should be {-1, 0, 3 }
	solution = Vector{-1, 0, 3}
	intersectionVector = line.IntersectPlane(plane)
	fmt.Println(intersectionVector)

	//Another test
	point = Vector{1, 2, 3}
	direction = Vector{1, 0, 0}
	line = Line{Vector: point, Direction: direction}

	normal = Vector{0, 1, 0}
	planeVector = Vector{0, 0, 0}
	plane = Plane3D{normal, planeVector}

	//Should panic
	intersectionVector = line.IntersectPlane(plane)
	fmt.Println(intersectionVector)

	if intersectionVector != solution {
		t.Errorf("Should be {1,0,3} got", intersectionVector)
	}

}
