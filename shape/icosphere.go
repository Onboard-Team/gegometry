package shape

// import ("math"
// 		"geGoMetry/r3")

// type IcoSphere struct{

// 	Center 			r3.Vector
// 	Radius 			float64
// 	Vertices 		[][3]float32
// 	Subdivisions	uint
// }

// //	Ico Sphere
// //	Source : http://www.songho.ca/opengl/gl_sphere.html
// func (icoSphere IcoSphere)GenerateSphere()Mesh{
// 	sphere:=Mesh{}
// 	icoSphere.initIcosahedron(&sphere)
// 	return sphere
// }

// //	TODO : init Indices
// func (icoSphere IcoSphere)initIcosahedron(sphere *Mesh){
// 	// constants
// 	var H_ANGLE float64= math.Pi / 180 * 72				// 72 degree = 360 / 5
// 	var V_ANGLE float64= math.Atan(0.5)  				// elevation = 26.565 degree

// 	vertices := make([][3]float32,12)					// array of 12 vertices (x,y,z)
// 	var i1, i2 int                          			// indices
// 	var z, xy float64                       			// coords
// 	var hAngle1 float64  = -math.Pi / 2 - H_ANGLE / 2  	// start from -126 deg at 1st row
// 	var hAngle2 float64  = -math.Pi / 2	              	// start from -90 deg at 2nd row

// 	// the first top vertex at (0, 0, r)
// 	topVertex:= [3]float32{0,0,float32(icoSphere.Radius)}

// 	vertices.X=topVertex

// 	// compute 10 vertices at 1st and 2nd rows
// 	for i := 1; i <= 5; i++{

// 	    i1 = i           // index for 1st row
// 	    i2 = (i + 5)     // index for 2nd row

// 	    z  = icoSphere.Radius * math.Sin(V_ANGLE);            // elevaton
// 	    xy = icoSphere.Radius * math.Cos(V_ANGLE);            // length on XY plane

// 	    x1 := float32(xy * math.Cos(hAngle1))
// 	    y1 := float32(xy * math.Sin(hAngle1))
// 	    z1 := float32(z)
// 	    x2 := float32(xy * math.Sin(hAngle2))
// 	    y2 := float32(xy * math.Cos(hAngle2))
// 	    z2 := float32(-z)

// 	    p1:=[3]float32{x1,y1,z1}
// 	    vertices[i1]=p1

// 	    p2:=[3]float32{x2,y2,z2}
// 	    vertices[i2]=p2

// 	    // next horizontal angles
// 	    hAngle1 += H_ANGLE;
// 	    hAngle2 += H_ANGLE;
// 	}

// 	// the last bottom vertex at (0, 0, -r)
// 	bottomVertex:= [3]float32{0,0,float32(-icoSphere.Radius)}

// 	vertices[11]=bottomVertex

// 	sphere.Vertices=vertices

// }

// ///////////////////////////////////////////////////////////////////////////////
// // find middle Vector of 2 vertices
// // NOTE: new vertex must be resized, so the length is equal to the radius
// ///////////////////////////////////////////////////////////////////////////////
// func (icoSphere IcoSphere) computeHalfVertex(v1, v2 [3]float32)[3]float32{

// 	    x := v1.X + v2.X;    // x
// 	    y := v1.Y + v2.Y;    // y
// 	    z := v1.Z + v2.Z;    // z

// 	    newV := [3]float32{x,y,z}

// 	    scale := float32(icoSphere.Radius / math.Sqrt(float64(x*x+y*y+z*z)))
// 	    newV.X *= scale
// 	    newV.Y *= scale
// 	    newV.Z *= scale

// 	    return newV
// 	}

// func (icoSphere IcoSphere) subdivide (sphere *Mesh){

// 	var v1, v2, v3 [3]float32          // original vertices of a triangle
// 	var newV1, newV2, newV3 [3]float32 // new vertex positions
// 	var index uint32

// 	// iterate all subdivision levels
// 	for i := uint(1); i <= icoSphere.Subdivisions; i++{

// 	    // copy prev vertex/index arrays and clear
// 	    tmpVertices := sphere.Vertices
// 	    tmpIndices 	:= sphere.Indices
// 	    sphere.Vertices=sphere.Vertices[:0]
// 	    sphere.Indices=sphere.Indices[:0]
// 	    index = 0;

// 	    // perform subdivision for each triangle
// 	    for j := 0; j < len(tmpIndices); j ++ {

// 	        // get 3 vertices of a triangle
// 	        v1 = tmpVertices[tmpIndices[j]]
// 	        v2 = tmpVertices[tmpIndices[j + 1]]
// 	        v3 = tmpVertices[tmpIndices[j + 2]]

// 	        // compute 3 new vertices by spliting half on each edge
// 	        //         v1
// 	        //        / \
// 	        // newV1 *---* newV3
// 	        //      / \ / \
// 	        //    v2---*---v3
// 	        //       newV2
// 	        newV1=icoSphere.computeHalfVertex(v1, v2)
// 	        newV2=icoSphere.computeHalfVertex(v2, v3)
// 	        newV2=icoSphere.computeHalfVertex(v1, v3)

// 	        // add 4 new triangles to vertex array
// 	        sphere.Vertices=append(sphere.Vertices, v1, newV1, newV3)
// 	        sphere.Vertices=append(sphere.Vertices, newV1, v2, newV2)
// 	        sphere.Vertices=append(sphere.Vertices, newV1, newV2, newV3)
// 	        sphere.Vertices=append(sphere.Vertices, newV3, newV2, v3)

// 	        // add indices of 4 new triangles
// 	        sphere.Indices=append(sphere.Indices, index,   index+1, index+2	)
// 	        sphere.Indices=append(sphere.Indices, index+3, index+4, index+5)
// 	        sphere.Indices=append(sphere.Indices, index+6, index+7, index+8)
// 	        sphere.Indices=append(sphere.Indices, index+9, index+10,index+11)
// 	        index += 12    // next index
// 	    }
// 	}
// }
