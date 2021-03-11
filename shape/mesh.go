package shape

import "geGoMetry/r3"

// Mesh : tessellated mesh for glTF exporter.
type Mesh struct {
	Vertices  []r3.Vector
	Normals   [][3]float64
	TexCoords [][2]float64
	Indices   []uint32
}
