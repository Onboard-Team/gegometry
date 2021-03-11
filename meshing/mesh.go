package mesh

import (
	"fmt"
	"geGoMetry/predicates"
	"geGoMetry/r3"
	"sort"
	"strconv"
	"strings"
)

type tetra struct {
	ia, ib, ic, id int // point indices
	faces          []face
}
type face struct {
	pa, pb, pc int
}
type edge struct {
	start, end int
}

type tetrahedrisation map[string]tetra

type halfEdge struct {
	twin   *halfEdge
	next   *halfEdge
	vertex r3.Vector
	edge   int
	face   int
}

type Meshing struct {
	vertices         []r3.Vector
	indices          []uint32
	triangulationMap map[int][]tetra
}

type faceCheck struct {
	check   bool
	tetraID int
	faceID  int
}

func (e edge) isEqual(e1 edge) bool {
	if e.start == e1.start {
		if e.end == e1.end {
			return true
		}
	}
	return false
}

func (f face) isPartOfMesh() bool {

	found := (f.pa >= 0) && (f.pa <= 3)
	if found {
		return false
	}
	found = (f.pb >= 0) && (f.pb <= 3)
	if found {
		return false
	}

	found = (f.pc >= 0) && (f.pc <= 3)
	if found {
		return false
	}

	return true
}

func (f face) hash() string {
	indexList := []int{f.pa, f.pb, f.pc}
	sort.Ints(indexList)
	hashList := []string{strconv.Itoa(indexList[0]), strconv.Itoa(indexList[1]), strconv.Itoa(indexList[2])}
	hash := strings.Join(hashList, ".")
	return hash
}

func (t tetra) getNearestPointToCentroid(points []r3.Vector) r3.Vector {

	pa, pb, pc, pd := points[t.ia], points[t.ib], points[t.ic], points[t.id]

	centroid := findCentroid(pa, pb, pc, pd)

	startPoint := getNearestPoint(points, centroid)
	return startPoint

}

func (t *tetrahedrisation) init(points *[]r3.Vector) {
	//fmt.Println("Bazinga!", points)

	initialTetra := englobe(points)

	//fmt.Println("new points:", points)

	t.addTetra(initialTetra)
}

func (t *tetrahedrisation) addTetra(newTetra tetra) {
	hash := tetraHash(newTetra)
	(*t)[hash] = newTetra
}

func tetraHash(t tetra) string {
	return string(strconv.Itoa(t.ia) + "." + strconv.Itoa(t.ib) + "." + strconv.Itoa(t.ic) + "." + strconv.Itoa(t.id))
}

func englobe(points *[]r3.Vector) tetra {
	h := helper{}
	h.findExtremas((*points)...)
	h.computeBoundingBox()

	// add englobing  tetra endices at top of point stack
	p1, p2, p3, p4 := h.computerInitialTetraPoints()
	(*points) = append([]r3.Vector{p1, p2, p3, p4}, (*points)...)

	// in glTF we compute in terms of faces and not in terms of
	// edges.
	// we need to add edges instead of just faces or
	//
	tetraFaces := []face{
		face{0, 1, 2},
		face{1, 2, 3},
		face{2, 3, 0},
		face{3, 0, 1},
	}

	initialTetra := tetra{0, 1, 2, 3, tetraFaces}
	return initialTetra
}

func addFace(meshIndices []int, f face) []int {
	meshIndices = append(meshIndices, f.pa-4, f.pb-4, f.pc-4)
	return meshIndices
}

//addPoint adds a point to the tessellation
func (m Meshing) addPoint(point r3.Vector) {

}

func connect(f face, pointIndex int) tetra {
	newTetraFaces := []face{
		face{f.pa, f.pb, f.pc},
		face{f.pb, f.pc, pointIndex},
		face{f.pc, pointIndex, f.pa},
		face{pointIndex, f.pa, f.pb},
	}

	return tetra{f.pa, f.pb, f.pc, pointIndex, newTetraFaces}
}

func bowyerWatson(points ...r3.Vector) []int {
	t := tetrahedrisation{}

	// the points you'll triangulate
	triangulationPoints := points

	t.init(&points)

	for pointIndex, point := range triangulationPoints {
		//fmt.Println("--------------------------")
		//fmt.Println("adding this point:", point)
		//fmt.Println("tetrahedrisation: ", t)
		badTetras := []tetra{}
		polyhedron := []face{}
		faceSet := make(map[string]faceCheck, 0)
		for _, tetra := range t {
			if predicates.InSphereFast(points[tetra.ia], points[tetra.ib], points[tetra.ic], points[tetra.id], point) >= 0 {
				badTetras = append(badTetras, tetra)
			}
		}
		//fmt.Println(badTetras)

		// getting faces to connect to the new point:
		for badTetraIndex, badTetra := range badTetras {
			for faceIndex, face := range badTetra.faces {
				if _, ok := faceSet[face.hash()]; !ok {
					faceSet[face.hash()] = faceCheck{false, badTetraIndex, faceIndex}
				} else {
					faceSet[face.hash()] = faceCheck{true, -1, -1}
				}
			}
		}
		for _, faceCheck := range faceSet {
			if faceCheck.check == false {
				badTetra := badTetras[faceCheck.tetraID]
				polyhedron = append(polyhedron, badTetra.faces[faceCheck.faceID])
			}
		}
		//fmt.Println("polyhedron", polyhedron)

		// removing the badtetras
		for _, badTetra := range badTetras {
			badTetraHash := tetraHash(badTetra)
			delete(t, badTetraHash)
		}
		//	connecting the good ones
		for _, face := range polyhedron {
			t.addTetra(connect(face, pointIndex+4))
		}
	}
	//removing the faces we don't need and creating the mesh:
	meshIndices := []int{}
	for _, tetra := range t {

		for _, face := range tetra.faces {
			if face.isPartOfMesh() {
				meshIndices = addFace(meshIndices, face)
			}
		}
	}
	fmt.Println(meshIndices)

	return meshIndices

}
