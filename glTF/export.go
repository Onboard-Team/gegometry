package glTF

import (
	"fmt"
	"geGoMetry/shape"

	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
)

func Export(mesh shape.Mesh, outputFile string) {

	fmt.Println("Exporting glTF")
	doc := gltf.NewDocument()

	model := transferModel(mesh)

	positionAccessor := modeler.WritePosition(doc, model.Positions)
	indicesAccessor := modeler.WriteIndices(doc, model.Indices)
	// NormalsAccessor := modeler.WriteNormal(doc, m.Normals)
	// TexCoordsAccessor := modeler.WriteTextureCoord(doc, m.TexCoords)

	//colors := make ([][3]uint8 , len(positions))
	colors := make([][3]uint8, 0)

	black := [3]uint8{0, 0, 0}
	white := [3]uint8{255, 255, 255}
	_ = black

	for i := 0; i < len(mesh.Vertices); i++ {
		colors = append(colors, white)
	}

	colorIndices := modeler.WriteColor(doc, colors)

	doc.Meshes = []*gltf.Mesh{{
		Name: "drawTriangle",
		Primitives: []*gltf.Primitive{{
			Indices: gltf.Index(indicesAccessor),
			Attributes: map[string]uint32{
				"POSITION": positionAccessor,
				"COLOR_0":  colorIndices,
				// "NORMAL":     NormalsAccessor,
				// "TEXCOORD_1": TexCoordsAccessor,
			},
		}},
	}}
	doc.Nodes = []*gltf.Node{{Name: "Root", Mesh: gltf.Index(0)}}
	doc.Scenes[0].Nodes = append(doc.Scenes[0].Nodes, 0)
	if err := gltf.SaveBinary(doc, outputFile); err != nil {
		panic(err)
	}
}
