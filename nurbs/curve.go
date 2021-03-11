package nurbs

import (
	"fmt"

	"github.com/ungerik/go3d/float64/vec3"
)

/* curvePoint : computes a curve point */
/* Input:
-- n : last index in control points vector
-- p : degree of the polynomial
-- U : Knot Vector
-- P : list of the control points
-- u : the variable that lies on the knot span
*/
func curvePoint(n, p int, U []int, P []vec3.T, u float64) vec3.T {
	fmt.Println("creating curve point")
	span := findSpan(n, p, u, U)
	N := basisFuns(span, u, p, U)
	C := vec3.T{}
	for i := 0; i <= p; i++ {
		B := P[span-p+i].Scaled(N[i])
		C.Add(&B)
	}
	return C
}
