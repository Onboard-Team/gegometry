package shape

import (
	_ "fmt"
	"geGoMetry/r3"
	"math"
)

// Sphere : constructs a UV sphere.
type Sphere struct {
	center      r3.Vector
	radius      float64
	sectorCount int
	stackCount  int
}

//Construct : constructs the sphere
func (sd Sphere) Construct() Mesh {

	sphere := Mesh{}

	sd.generateVertices(&sphere)
	sd.generateIndices(&sphere)

	return sphere

}

func (sd Sphere) generateVertices(sphere *Mesh) {

	var x, y, z, xy float64           //	vertex position
	var nx, ny, nz, lengthInv float64 //	vertex normal
	var s, t float64                  //	vertex texCoord

	sectorStep := 2 * math.Pi / float64(sd.sectorCount)
	stackStep := math.Pi / float64(sd.stackCount)

	lengthInv = 1.0 / float64(sd.radius)

	for i := 0; i < sd.stackCount; i++ {

		stackAngle := math.Pi/2 - float64(i)*stackStep // starting from pi/2 to -pi/2
		xy = sd.radius * math.Cos(stackAngle)          // r * cos(u)
		z = sd.radius * math.Sin(stackAngle)           // r * sin(u)

		// add (sectorCount+1) vertices per stack
		// the first and last vertices have same position and normal, but different tex coords
		for j := 0; j <= sd.sectorCount; j++ {

			sectorAngle := float64(j) * sectorStep // starting from 0 to 2pi

			// vertex position (x, y, z)
			x = xy * math.Cos(sectorAngle) // r * cos(u) * cos(v)
			y = xy * math.Sin(sectorAngle) // r * cos(u) * sin(v)

			Vector := r3.Vector{x, y, z}
			sphere.Vertices = append(sphere.Vertices, Vector)

			// normalized vertex normal (nx, ny, nz)
			nx = x * lengthInv
			ny = y * lengthInv
			nz = z * lengthInv

			normal := [3]float64{nx, ny, nz}
			sphere.Normals = append(sphere.Normals, normal)

			// vertex tex coord (s, t) range between [0, 1]
			s = float64(j / sd.sectorCount)
			t = float64(i / sd.stackCount)

			texCoord := [2]float64{s, t}
			sphere.TexCoords = append(sphere.TexCoords, texCoord)
		}
	}
}

func (sd Sphere) generateIndices(sphere *Mesh) {

	var k1, k2 uint32

	//	TOP
	k1 = 0
	k2 = uint32(sd.sectorCount + 1)

	//fmt.Println("TOP")
	for j := 0; j < sd.sectorCount; j, k1, k2 = j+1, k1+1, k2+1 {
		//fmt.Println("A",k1+1,"B", k2, "C", k2+1)
		// k1+1 => k2 => k2+1
		sphere.Indices = append(sphere.Indices, k1+1, k2, k2+1)
	}

	//	BOTTOM
	// k1 = uint32((sd.stackCount-1) * (sd.sectorCount + 1))
	// k2 = 0 //k1+uint32(sd.sectorCount )

	// for j := 0; j < sd.sectorCount; j ,k1= j+1, k1+1{

	//     // k1+1 => k2 => k2+1
	//     indices = append(indices, k1, k2, k1+1);

	// }

	//	MIDDLE
	for i := 1; i < sd.stackCount-1; i++ {

		k1 = uint32(i * (sd.sectorCount + 1)) // beginning of current stack
		k2 = k1 + uint32(sd.sectorCount+1)    // beginning of next stack

		for j := 0; j < sd.sectorCount; j, k1, k2 = j+1, k1+1, k2+1 {

			// 2 triangles per sector excluding first and last stacks
			// k1 => k2 => k1+1
			sphere.Indices = append(sphere.Indices, k1, k2, k1+1)

			// k1+1 => k2 => k2+1
			sphere.Indices = append(sphere.Indices, k1+1, k2, k2+1)

		}

	}

}
