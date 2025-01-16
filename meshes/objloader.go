package meshes

import "bufio"
import "os"
import "strings"
import "strconv"

import lm "derf.space/gldemo/linearmath"

type Vertex struct {
	Vertice lm.Vec3
	TexCoord lm.Vec2
	Normal lm.Vec3
}

type Triangle struct {
	V1 int32
	V2 int32
	V3 int32
}

type MeshObj struct {
	Vertices []Vertex
	Triangles []Triangle
}

type Face struct {
	tris int32
	v int32
	vt int32
	vn int32
}

type FaceTriangle struct {
	v1 Face
	v2 Face
	v3 Face
}

func ObjLoader(objMesh *MeshObj, path string) {
	var v []lm.Vec3
	var vt []lm.Vec2
	var vn []lm.Vec3
	//var f []Face
	var f []FaceTriangle

	var fp *os.File
	var err error
	var scanner *bufio.Scanner

	fp, err = os.Open(path)

	if err != nil {
		panic(err)
	}

	scanner = bufio.NewScanner(fp)

	scanner.Split(bufio.ScanLines)

	var index int32 = 0

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		var line []string

		line = strings.Split(scanner.Text(), " ")

		if(line[0] == "v") {
			var x, y, z float64
			x, err = strconv.ParseFloat(line[1], 32)
			if err != nil {
				panic(err)
			}
			y, err = strconv.ParseFloat(line[2], 32)
			if err != nil {
				panic(err)
			}
			z, err = strconv.ParseFloat(line[3], 32)
			if err != nil {
				panic(err)
			}
			var temp lm.Vec3 = lm.Vec3{float32(x), float32(y), float32(z)}
			v = append(v, temp)
		} else if(line[0] == "vt") {
			var x, y float64
			x, err = strconv.ParseFloat(line[1], 32)
			if err != nil {
				panic(err)
			}
			y, err = strconv.ParseFloat(line[2], 32)
			if err != nil {
				panic(err)
			}
			var temp = lm.Vec2{float32(x), float32(y)}
			vt = append(vt, temp)

		} else if(line[0] == "vn") {
			var x, y, z float64
			x, err = strconv.ParseFloat(line[1], 32)
			if err != nil {
				panic(err)
			}
			y, err = strconv.ParseFloat(line[2], 32)
			if err != nil {
				panic(err)
			}
			z, err = strconv.ParseFloat(line[3], 32)
			if err != nil {
				panic(err)
			}
			var temp = lm.Vec3{float32(x), float32(y), float32(z)}
			vn = append(vn, temp)
		} else if(line[0] == "f") {
			var temp FaceTriangle

			// v1

			var args []string = strings.Split(line[1], "/")

			var v, vt, vn int64
			v, err = strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				panic(err)
			}
			vt, err = strconv.ParseInt(args[1], 10, 32)
			if err != nil {
				panic(err)
			}
			vn, err = strconv.ParseInt(args[2], 10, 32)
			if err != nil {
				panic(err)
			}

			temp.v1.v = int32(v) - 1
			temp.v1.vt = int32(vt) - 1
			temp.v1.vn = int32(vn) - 1
			temp.v1.tris = index
			index += 1

			// v2
			args = strings.Split(line[2], "/")
			v, err = strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				panic(err)
			}
			vt, err = strconv.ParseInt(args[1], 10, 32)
			if err != nil {
				panic(err)
			}
			vn, err = strconv.ParseInt(args[2], 10, 32)
			if err != nil {
				panic(err)
			}

			temp.v2.v = int32(v) - 1
			temp.v2.vt = int32(vt) - 1
			temp.v2.vn = int32(vn) - 1
			temp.v2.tris = index

			index += 1
			// v3
			args = strings.Split(line[3], "/")
			v, err = strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				panic(err)
			}
			vt, err = strconv.ParseInt(args[1], 10, 32)
			if err != nil {
				panic(err)
			}
			vn, err = strconv.ParseInt(args[2], 10, 32)
			if err != nil {
				panic(err)
			}

			temp.v3.v = int32(v) - 1
			temp.v3.vt = int32(vt) - 1
			temp.v3.vn = int32(vn) - 1
			temp.v3.tris = index
			index += 1

			f = append(f, temp)
		}
	}

	fp.Close()

	// Actually Create
	for i := 0; i < len(f); i++ {
		var temp Triangle
		var tempVertex Vertex

		// v1
		temp.V1 = f[i].v1.tris
		tempVertex.Vertice = v[f[i].v1.v]
		tempVertex.TexCoord = vt[f[i].v1.vt]
		tempVertex.Normal = vn[f[i].v1.vn]

		objMesh.Vertices = append(objMesh.Vertices, tempVertex)

		// v2
		temp.V2 = f[i].v2.tris
		tempVertex.Vertice = v[f[i].v2.v]
		tempVertex.TexCoord = vt[f[i].v2.vt]
		tempVertex.Normal = vn[f[i].v2.vn]

		objMesh.Vertices = append(objMesh.Vertices, tempVertex)

		// v3
		temp.V3 = f[i].v3.tris
		tempVertex.Vertice = v[f[i].v3.v]
		tempVertex.TexCoord = vt[f[i].v3.vt]
		tempVertex.Normal = vn[f[i].v3.vn]

		objMesh.Vertices = append(objMesh.Vertices, tempVertex)
		
		objMesh.Triangles = append(objMesh.Triangles, temp)
	}
}