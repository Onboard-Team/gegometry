package shape

import (
	"geGoMetry/r3"
	"math"
)

// Torus : constructs a UV torus.
type Torus struct {
	center      r3.Vector
	radius      float64 // radius going from the center to the tube revolution circle
	tubeRadius  float64 // radius of the tube
	sectorCount int
	stackCount  int
}

// Construct : constructs the torus
func (td Torus) Construct() Mesh {

	torus := Mesh{}

	td.generateVertices(&torus)
	td.generateIndices(&torus)

	return torus

}

func (td Torus) generateVertices(torus *Mesh) {

	var x, y, z, xy float64           //	vertex position
	var nx, ny, nz, lengthInv float64 //	vertex normal
	var s, t float64                  //	vertex texCoord

	sectorStep := 2 * math.Pi / float64(td.sectorCount)
	stackStep := 2 * math.Pi / float64(td.stackCount)

	lengthInv = 1.0 / float64(td.radius)

	for i := 0; i < td.stackCount; i++ {

		stackAngle := math.Pi/2 - float64(i)*stackStep      // starting from pi/2 to -pi/2
		xy = td.radius + td.tubeRadius*math.Cos(stackAngle) // R + r * cos(u)
		z = td.tubeRadius * math.Sin(stackAngle)            // r * sin(u)

		// add (sectorCount+1) vertices per stack
		// the first and last vertices have same position and normal, but different tex coords
		for j := 0; j <= td.sectorCount; j++ {

			sectorAngle := float64(j) * sectorStep // starting from 0 to 2pi

			// vertex position (x, y, z)
			x = xy * math.Cos(sectorAngle) // (R + r * cos(u)) * cos(v)
			y = xy * math.Sin(sectorAngle) // (R + r * cos(u)) * sin(v)

			Vector := r3.Vector{x, y, z}
			torus.Vertices = append(torus.Vertices, Vector)

			// normalized vertex normal (nx, ny, nz)
			nx = x * lengthInv
			ny = y * lengthInv
			nz = z * lengthInv

			normal := [3]float64{nx, ny, nz}
			torus.Normals = append(torus.Normals, normal)

			// vertex tex coord (s, t) range between [0, 1]
			s = float64(j / td.sectorCount)
			t = float64(i / td.stackCount)

			texCoord := [2]float64{s, t}
			torus.TexCoords = append(torus.TexCoords, texCoord)
		}
	}
}

func (td Torus) generateIndices(torus *Mesh) {

	var k1, k2 uint32

	//	TOP
	k1 = 0
	k2 = uint32(td.sectorCount + 1)

	//fmt.Println("TOP")
	for j := 0; j < td.sectorCount; j, k1, k2 = j+1, k1+1, k2+1 {
		//fmt.Println("A",k1+1,"B", k2, "C", k2+1)
		// k1+1 => k2 => k2+1
		torus.Indices = append(torus.Indices, k1+1, k2, k2+1)
	}

	//	BOTTOM
	// k1 = uint32((td.stackCount - 1) * (td.sectorCount + 1))
	// k2 = 0 //k1+uint32(td.sectorCount )

	// for j := 0; j < td.sectorCount; j, k1 = j+1, k1+1 {

	// 	// k1+1 => k2 => k2+1
	// 	torus.Indices = append(torus.Indices, k1, k2, k1+1)

	// }

	//	MIDDLE
	for i := 1; i < td.stackCount-1; i++ {

		k1 = uint32(i * (td.sectorCount + 1)) // beginning of current stack
		k2 = k1 + uint32(td.sectorCount+1)    // beginning of next stack

		for j := 0; j < td.sectorCount; j, k1, k2 = j+1, k1+1, k2+1 {

			// 2 triangles per sector excluding first and last stacks
			// k1 => k2 => k1+1
			torus.Indices = append(torus.Indices, k1, k2, k1+1)

			// k1+1 => k2 => k2+1
			torus.Indices = append(torus.Indices, k1+1, k2, k2+1)

		}

	}

}
