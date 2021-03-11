package mesh

import "geGoMetry/r3"

//link to the paper :
//https://www.kiv.zcu.cz/site/documents/verejne/vyzkum/publikace/technicke-zpravy/2009/tr-2009-03.pdf?fbclid=IwAR34DUXzbkTk-YRgABKRGDHfsQ-W3so1A1JNG4yoWaXrskejxn13LEPpIUA

// tetra a, b, c, d
// indices abc bcd cda dab

// flip14 – replaces 1 tetrahedron by 4.
// It inserts a point into the triangulation.

func (t *Meshing) flip14(p r3.Vector) {
	t.vertices = append(t.vertices, p)
	tetra1Indices := []uint32{
		0, 1, 2,
		1, 2, 4,
		2, 4, 0,
		4, 0, 1,
	}
	tetra2Indices := []uint32{
		0, 1, 3,
		1, 3, 4,
		3, 4, 0,
		4, 0, 1,
	}
	tetra3Indices := []uint32{
		1, 2, 3,
		2, 3, 4,
		3, 4, 1,
		4, 1, 2,
	}
	tetra4Indices := []uint32{
		0, 2, 3,
		2, 3, 4,
		3, 4, 0,
		4, 0, 2,
	}
	duplicateInReverse(&tetra1Indices)
	duplicateInReverse(&tetra2Indices)
	duplicateInReverse(&tetra3Indices)
	duplicateInReverse(&tetra4Indices)

	t.indices = []uint32{}
	t.indices = append(t.indices, tetra1Indices...)
	t.indices = append(t.indices, tetra2Indices...)
	t.indices = append(t.indices, tetra3Indices...)
	t.indices = append(t.indices, tetra4Indices...)

}

// flip41 – replaces 4 tetrahedra by 1.
// It removes a point from the triangulation.
func (t Meshing) flip41(index, p r3.Vector) {

}

// flip23 – replaces 2 tetrahedra by 3.
// It destroys an inner face shared by the two flipped
// tetrahedra and replaces it by three new inner faces.
func flip32(t Meshing, p r3.Vector) {

}

// flip32 – replaces 3 tetrahedra by 2.
// It destroys three inner faces, each shared by twoof three flipped
// tetrahedra, and replaces them by one new inner face
func flip23(t Meshing, p r3.Vector) {

}
