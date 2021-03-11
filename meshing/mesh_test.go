package mesh

import (
	"fmt"
	"geGoMetry/r3"
	"testing"
)

func TestSort(t *testing.T) {
	pa := r3.Vector{X: 0, Y: 0, Z: 0}
	pb := r3.Vector{X: 0, Y: 999, Z: 0}
	pc := r3.Vector{X: 1, Y: 17, Z: 0}
	pd := r3.Vector{X: 1, Y: 1, Z: 0}
	pe := r3.Vector{X: 0, Y: 0, Z: 12}
	pf := r3.Vector{X: -123, Y: 1, Z: 15}
	pg := r3.Vector{X: 1, Y: 123, Z: 1}
	ph := r3.Vector{X: 1, Y: 1, Z: 1}
	points := []r3.Vector{pa, pb, pc, pd, pe, pf, pg, ph}
	centroid := r3.Vector{X: 1, Y: 123, Z: 1}

	nearestPoint := getNearestPoint(points, centroid)
	fmt.Println("nearest Point :", nearestPoint)
}
func TestEnglobe(t *testing.T) {
	pa := r3.Vector{X: 0, Y: 0, Z: 0}
	pb := r3.Vector{X: 0, Y: 999, Z: 0}
	pc := r3.Vector{X: 1, Y: 17, Z: 0}
	pd := r3.Vector{X: 1, Y: 1, Z: 0}
	pe := r3.Vector{X: 0, Y: 0, Z: 12}
	pf := r3.Vector{X: -123, Y: 1, Z: 15}
	pg := r3.Vector{X: 1, Y: 123, Z: 1}
	ph := r3.Vector{X: 1, Y: 1, Z: 1}

	points := []r3.Vector{pa, pb, pc, pd, pe, pf, pg, ph}

	tetra := englobe(&points)
	v1, v2, v3, v4 := points[tetra.ia], points[tetra.ib], points[tetra.ic], points[tetra.id]
	for _, p := range points {
		if vectorInTetrahedron(v1, v2, v3, v4, p) == false {
			fmt.Println("The point ", p, " is not inside The initial Tetra")
		}
	}
	fmt.Println("created initial tetra")
}

func TestBowyerWatson(t *testing.T) {
	fmt.Println("_____tessellation test")

	pa := r3.Vector{X: 0, Y: 0, Z: 0}
	pb := r3.Vector{X: 0, Y: 5, Z: 0}
	pc := r3.Vector{X: 1, Y: 7, Z: 0}
	pd := r3.Vector{X: 1, Y: 1, Z: 0}
	pe := r3.Vector{X: 0, Y: 0, Z: 2}
	pf := r3.Vector{X: -3, Y: 1, Z: 5}
	pg := r3.Vector{X: 1, Y: 3, Z: 1}
	ph := r3.Vector{X: 1, Y: 1, Z: 1}
	points := []r3.Vector{pa, pb, pc, pd, pe, pf, pg, ph}

	bowyerWatson(points...)
	//outputFile := "../testfiles/exports/flip14.glb"
	//glTF.Export(mesh, outputFile)

}
