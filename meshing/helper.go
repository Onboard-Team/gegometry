package mesh

import (
	"fmt"
	"geGoMetry/r3"
	"math"
)

type helper struct {
	min, max r3.Vector
}

//UPDATE : https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values

type distancePair struct {
	vector   r3.Vector
	distance float64
}

type distancePairList []distancePair

func (p distancePairList) Len() int           { return len(p) }
func (p distancePairList) Less(i, j int) bool { return p[i].distance < p[j].distance }
func (p distancePairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func getNearestPoint(vertices []r3.Vector, centroid r3.Vector) r3.Vector {

	minDistance := math.MaxFloat64

	nearestPoint := r3.Vector{}

	for _, vertex := range vertices {

		distance := r3.Distance(vertex, centroid)
		fmt.Println("distance", distance, " point:", vertex, " nearestPoint:", nearestPoint)
		if distance <= minDistance {

			minDistance = distance
			nearestPoint = vertex
		}

	}
	return nearestPoint
}

func (h *helper) findExtremas(points ...r3.Vector) {
	h.min = r3.Vector{X: math.MaxFloat64, Y: math.MaxFloat64, Z: math.MaxFloat64}
	h.max = r3.Vector{X: -math.MaxFloat64, Y: -math.MaxFloat64, Z: -math.MaxFloat64}

	for _, point := range points {
		if h.max.X < point.X {
			h.max.X = point.X
		}
		if h.max.Y < point.Y {
			h.max.Y = point.Y
		}
		if h.max.Z < point.Z {
			h.max.Z = point.Z
		}

		if h.min.X > point.X {
			h.min.X = point.X
		}
		if h.min.Y > point.Y {
			h.min.Y = point.Y
		}
		if h.min.Z > point.Z {
			h.min.Z = point.Z
		}
	}
	//fmt.Println("h.max:", h.max, "\nh.min:", h.min)
}

func (h helper) computeBoundingBox() {
	xmin, ymin, zmin := h.min.X, h.min.Y, h.min.Z
	xmax, ymax, zmax := h.max.X, h.max.Y, h.max.Z

	//bounding box
	p1 := r3.Vector{X: xmin, Y: ymin, Z: zmin}
	p2 := r3.Vector{X: xmax, Y: ymin, Z: zmin}
	p3 := r3.Vector{X: xmin, Y: ymax, Z: zmin}
	p4 := r3.Vector{X: xmin, Y: ymin, Z: zmax}
	p5 := r3.Vector{X: xmax, Y: ymax, Z: zmin}
	p6 := r3.Vector{X: xmin, Y: ymax, Z: zmax}
	p7 := r3.Vector{X: xmax, Y: ymin, Z: zmax}
	p8 := r3.Vector{X: xmax, Y: ymax, Z: zmax}

	boundingBox := make([]r3.Vector, 0)
	boundingBox = append(boundingBox, p1, p2, p3, p4, p5, p6, p7, p8)
	//fmt.Println(boundingBox)

}

func (h helper) computerInitialTetraPoints() (r3.Vector, r3.Vector, r3.Vector, r3.Vector) {
	xmin, ymin, zmin := h.min.X, h.min.Y, h.min.Z
	xmax, ymax, zmax := h.max.X, h.max.Y, h.max.Z

	xtetra := xmin + 3*(xmax-xmin)
	ytetra := ymin + 3*(ymax-ymin)
	ztetra := zmin + 3*(zmax-zmin)

	//Bounding Tetrahedron
	p1 := r3.Vector{X: xmin, Y: ymin, Z: zmin}
	p2 := r3.Vector{X: xtetra, Y: ymin, Z: zmin}
	p3 := r3.Vector{X: xmin, Y: ytetra, Z: zmin}
	p4 := r3.Vector{X: xmin, Y: ymin, Z: ztetra}

	return p1, p2, p3, p4
}
