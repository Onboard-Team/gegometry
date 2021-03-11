package glTF

import (
	"testing"

	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
)

func TestRectangle(t *testing.T) {
	doc := gltf.NewDocument()

	positionAccessor := modeler.WritePosition(doc, [][3]float32{{0, 0, 0}, {0, 10, 0},
		{0, 0, 10}, {0, 10, 10}})

	indicesAccessor := modeler.WriteIndices(doc, []uint8{0, 1, 2,
		3, 2, 1})
	colorIndices := modeler.WriteColor(doc, [][3]uint8{{50, 155, 255}, {0, 100, 200}, {255, 155, 50}, {255, 155, 50}})

	doc.Meshes = []*gltf.Mesh{{
		Name: "drawTriangle",
		Primitives: []*gltf.Primitive{{
			Indices: gltf.Index(indicesAccessor),
			Attributes: map[string]uint32{
				"POSITION": positionAccessor,
				"COLOR_0":  colorIndices,
			},
		}},
	}}
	doc.Nodes = []*gltf.Node{{Name: "Root", Mesh: gltf.Index(0)}}
	doc.Scenes[0].Nodes = append(doc.Scenes[0].Nodes, 0)
	if err := gltf.SaveBinary(doc, "./example.glb"); err != nil {
		panic(err)
	}
}

func TestTetrahedron(t *testing.T) {
	doc := gltf.NewDocument()

	positionAccessor := modeler.WritePosition(doc, [][3]float32{
		{0, 0, 1}, {0, 0, 1}, {5.3028762e-17, 0.8660254, -0.5},
		{0.8660254, 0, -0.5}, {-5.3028762e-17, -0.8660254, -0.5}, {-0.8660254, -0, -0.5}})

	indicesAccessor := modeler.WriteIndices(doc, []uint8{0, 2, 3,
		0, 3, 4,
		0, 4, 5,
		0, 5, 2,
	})
	colorIndices := modeler.WriteColor(doc, [][3]uint8{{50, 155, 255}, {0, 100, 200}, {50, 155, 255}, {0, 100, 200}, {255, 155, 50}, {255, 155, 50}})

	doc.Meshes = []*gltf.Mesh{{
		Name: "drawTriangle",
		Primitives: []*gltf.Primitive{{
			Indices: gltf.Index(indicesAccessor),
			Attributes: map[string]uint32{
				"POSITION": positionAccessor,
				"COLOR_0":  colorIndices,
			},
		}},
	}}
	doc.Nodes = []*gltf.Node{{Name: "Root", Mesh: gltf.Index(0)}}
	doc.Scenes[0].Nodes = append(doc.Scenes[0].Nodes, 0)
	if err := gltf.SaveBinary(doc, "./example.glb"); err != nil {
		panic(err)
	}
}
