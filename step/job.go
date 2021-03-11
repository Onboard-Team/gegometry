package step

import "fmt"

type bSplineSurfaceWithKnots struct {
	name              string
	uDegree           string
	vDegree           string
	controlPointsList [][]string
	surfaceForm       string /* should be transformed to enum */
	uClosed           bool
	vClosed           bool
	selfIntersect     bool
	uMultiplicities   []string
	vMultiplicities   []string
	uKnots            []string
	vKnots            []string
}

func (nurb bSplineSurfaceWithKnots) print() {

	fmt.Println("name:", nurb.name)
	fmt.Println("uDegree:", nurb.uDegree)
	fmt.Println("vDegree:", nurb.vDegree)
	fmt.Println("controlPointsList:", nurb.controlPointsList)
	fmt.Println("surfaceForm:", nurb.surfaceForm)
	fmt.Println("uClosed:", nurb.uClosed)
	fmt.Println("vClosed", nurb.vClosed)
	fmt.Println("selfIntersect:", nurb.selfIntersect)
	fmt.Println("uMultiplicities:", nurb.uMultiplicities)
	fmt.Println("vMultiplicities:", nurb.vMultiplicities)
	fmt.Println("uKnots:", nurb.uKnots)
	fmt.Println("vKnots:", nurb.vKnots)

}
