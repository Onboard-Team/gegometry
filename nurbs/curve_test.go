package nurbs

import (
	"fmt"
	"testing"

	"github.com/ungerik/go3d/float64/vec3"
)

func TestCurvePoint(t *testing.T) {
	p1 := vec3.T{0, 0, 0}
	p2 := vec3.T{1, 0, 0}
	p3 := vec3.T{0, 1, 0}
	p4 := vec3.T{1, 1, 0}
	p5 := vec3.T{0, 0, 1}
	p6 := vec3.T{1, 0, 1}
	p7 := vec3.T{0, 1, 1}
	p8 := vec3.T{1, 1, 1}
	p9 := vec3.T{2, 2, 2}

	controlPoints := []vec3.T{p1, p2, p3, p4, p5, p6, p7, p8, p9}
	/* we have :
	---m+1 knots : 13
	---n+1 controlPoints: 9
	---p   degree: 3
	------related by : m=n+p+1
	*/
	//FAIL:fmt.Println(curvePoint(8, 3, knots, controlPoints, 0.5))
	//FAIL:fmt.Println(curvePoint(8, 3, knots, controlPoints, 1.5))
	//knots := []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1}

	knots := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	fmt.Println(curvePoint(8, 3, knots, controlPoints, 4.5))
}
