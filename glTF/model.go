package glTF

import "geGoMetry/shape"

type Model struct {
	Positions [][3]float32
	Indices   []uint32
}

func transferModel(mesh shape.Mesh) Model {

	//	transfer positions
	positions := make([][3]float32, 0)

	for _, v := range mesh.Vertices {
		p := [3]float32{
			float32(v.X),
			float32(v.Y),
			float32(v.Z)}
		positions = append(positions, p)
	}

	model := Model{Positions: positions, Indices: mesh.Indices}

	return model
}
