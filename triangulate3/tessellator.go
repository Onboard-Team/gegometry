package triangulate3

import (
	"fmt"
	"geGoMetry/r3"
)

type Tessellator struct {
	Vectors []r3.Vector
	indices []r3.Vector
}

func (t Tessellator) tessellate(Vectors []r3.Vector) {

	triangulatedMesh := createInitialTetra(Vectors)

	_ = triangulatedMesh
	for _, Vector := range Vectors {
		
		fmt.Println(Vector)
	}

}
