package staticmesh

import "fmt"
import "io/ioutil"
import "strings"
import "unsafe"

/*
import "bufio"
import "os"
*/

import sdl "github.com/veandco/go-sdl2/sdl"
import img "github.com/veandco/go-sdl2/img"

import "github.com/go-gl/gl/v4.1-compatibility/gl"

import "derf.space/gldemo/app"
import lm "derf.space/gldemo/linearmath"
import tr "derf.space/gldemo/transform"

import "derf.space/gldemo/meshes"

type DemoApp struct {
	// Shader
	vertexShader uint32
	fragmentShader uint32
	// Program
	program uint32
	// Vertex Array
	vertexArray uint32
	// Uniforms
	uProj int32
	uView int32
	uModel int32
	uTex0 int32
	// Attributes
	aVertices uint32
	//aColors uint32
	aTexCoords uint32
	aNormal uint32

	// Vertex Buffer for Triangle

	verticesList []float32
	verticesVertexBuffer uint32

	// Added New
	/*
	colorList []float32
	colorVertexBuffer uint32
	*/

	texCoordList []float32
	texCoordVertexBuffer uint32

	normalList []float32
	normalVertexBuffer uint32

	indexList []uint32
	indexIndexBuffer uint32

	logoTexture uint32
	logoWidth int32
	logoHeight int32

	yrot float32

	cubeMesh meshes.MeshObj
}

func (this *DemoApp) Init() {
	gl.Enable(gl.DEPTH_TEST)

	// Setting Vertices
	this.aVertices = 0
	this.aTexCoords = 1
	this.aNormal = 2

	this.yrot = 32.0

	var err error
	// Shader
	this.vertexShader, err = createShader(gl.VERTEX_SHADER, "data/shaders/staticmesh/main.vert.glsl")
	if err != nil {
		panic(err)
	}

	this.fragmentShader, err = createShader(gl.FRAGMENT_SHADER, "data/shaders/staticmesh/main.frag.glsl")
	if err != nil {
		panic(err)
	}

	// Program
	this.program, err = createProgram(this.vertexShader, this.fragmentShader)
	if err != nil {
		panic(err)
	}


	// Vertex Array
	gl.GenVertexArrays(1, &this.vertexArray)

	// Setup Program
	gl.UseProgram(this.program)
	// Uniform
	this.uProj = gl.GetUniformLocation(this.program, gl.Str("proj\x00"))
	this.uView = gl.GetUniformLocation(this.program, gl.Str("view\x00"))
	this.uModel = gl.GetUniformLocation(this.program, gl.Str("model\x00"))
	this.uTex0 = gl.GetUniformLocation(this.program, gl.Str("tex0\x00"))
	gl.Uniform1i(this.uTex0, 0)

	// Vertex Array
	gl.BindVertexArray(this.vertexArray)
	gl.EnableVertexAttribArray(this.aVertices)
	gl.EnableVertexAttribArray(this.aTexCoords)
	gl.EnableVertexAttribArray(this.aNormal)
	gl.BindVertexArray(0)

	gl.UseProgram(0)

	meshes.ObjLoader(&this.cubeMesh, "data/mesh/cube.obj")

	gl.GenBuffers(1, &this.verticesVertexBuffer)
	gl.GenBuffers(1, &this.texCoordVertexBuffer)
	gl.GenBuffers(1, &this.normalVertexBuffer)

	for i:= 0; i < len(this.cubeMesh.Vertices); i++ {
		this.verticesList = addFloat3f(
			this.verticesList, 
			this.cubeMesh.Vertices[i].Vertice.X, 
			this.cubeMesh.Vertices[i].Vertice.Y, 
			this.cubeMesh.Vertices[i].Vertice.Z)
		this.texCoordList = addFloat2f(
			this.texCoordList,
			this.cubeMesh.Vertices[i].TexCoord.X,
			this.cubeMesh.Vertices[i].TexCoord.Y)
		this.normalList = addFloat3f(
			this.normalList,
			this.cubeMesh.Vertices[i].Normal.X,
			this.cubeMesh.Vertices[i].Normal.Y,
			this.cubeMesh.Vertices[i].Normal.Z)
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, this.verticesVertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(this.verticesList) * int(unsafe.Sizeof(float32(0))), gl.Ptr(this.verticesList), gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, this.texCoordVertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(this.texCoordList) * int(unsafe.Sizeof(float32(0))), gl.Ptr(this.texCoordList), gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, this.normalVertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(this.normalList) * int(unsafe.Sizeof(float32(0))), gl.Ptr(this.normalList), gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.GenBuffers(1, &this.indexIndexBuffer)

	for i := 0; i < len(this.cubeMesh.Triangles); i++ {
		this.indexList = addUint3(
			this.indexList,
			uint32(this.cubeMesh.Triangles[i].V1),
			uint32(this.cubeMesh.Triangles[i].V2),
			uint32(this.cubeMesh.Triangles[i].V3))
	}

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, this.indexIndexBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(this.indexList) * int(unsafe.Sizeof(uint32(0))), gl.Ptr(this.indexList), gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

	var surf *sdl.Surface

	surf, err = img.Load("data/textures/logo.png")

	if err != nil {
		panic(err)
	}

	this.logoWidth = surf.W
	this.logoHeight = surf.H

	gl.GenTextures(1, &this.logoTexture)

	gl.BindTexture(gl.TEXTURE_2D, this.logoTexture)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		this.logoWidth,
		this.logoHeight,
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(surf.Pixels()))
	
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)

	gl.GenerateMipmap(gl.TEXTURE_2D)

	gl.BindTexture(gl.TEXTURE_2D, 0)

	surf.Free()
}

func (this *DemoApp) HandleEvent(e sdl.Event) {

}

func (this *DemoApp) Update(delta float32) {
	this.yrot += 64.0 * delta

	if(this.yrot > 360.0) {
		this.yrot -= 360.0
	}
}

func (this *DemoApp) Render() {
	gl.ClearColor(135.0 / 255.0, 206.0 / 255.0, 235.0 / 255.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.UseProgram(this.program)

	// Projection
	var cm []float32 = tr.Perspective(lm.ToRadian(45.0), app.GetAspect(), 1.0, 1024.0).ToArray()
	gl.UniformMatrix4fv(this.uProj, 1, false, &cm[0])

	// View (identity)
	cm = lm.CreateMat4Identity().ToArray()
	gl.UniformMatrix4fv(this.uView, 1, false, &cm[0])

	// Model
	cm = tr.RotateAxis(lm.ToRadian(this.yrot), lm.CreateVec3(1.0, 1.0, 0.0)).Mul(tr.Translate(lm.CreateVec3(0.0, 0.0, -3.0))).ToArray()
	gl.UniformMatrix4fv(this.uModel, 1, false, &cm[0])

	// Draw Triangle

	gl.BindTexture(gl.TEXTURE_2D, this.logoTexture)

	gl.BindVertexArray(this.vertexArray)

	gl.BindBuffer(gl.ARRAY_BUFFER, this.verticesVertexBuffer)
	gl.VertexAttribPointer(this.aVertices, 3, gl.FLOAT, false, 0, nil)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, this.texCoordVertexBuffer)
	gl.VertexAttribPointer(this.aTexCoords, 2, gl.FLOAT, false, 0, nil)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, this.normalVertexBuffer)
	gl.VertexAttribPointer(this.aNormal, 3, gl.FLOAT, false, 0, nil)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, this.indexIndexBuffer)
	gl.DrawElements(gl.TRIANGLES, int32(len(this.indexList)), gl.UNSIGNED_INT, nil)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

	gl.BindVertexArray(0)

	gl.BindTexture(gl.TEXTURE_2D, 0)

	gl.UseProgram(0)
}

func (this *DemoApp) Release() {
	gl.DeleteTextures(1, &this.logoTexture)

	gl.DeleteBuffers(1, &this.indexIndexBuffer)
	this.indexList = this.indexList[:0]

	gl.DeleteBuffers(1, &this.normalVertexBuffer)
	this.normalList = this.normalList[:0]

	gl.DeleteBuffers(1, &this.texCoordVertexBuffer)
	this.texCoordList = this.texCoordList[:0]

	gl.DeleteBuffers(1, &this.verticesVertexBuffer)
	this.verticesList = this.verticesList[:0]
	
	gl.DeleteVertexArrays(1, &this.vertexArray)
	
	deleteProgram(this.program, this.vertexShader, this.fragmentShader)
	
	gl.DeleteShader(this.vertexShader)
	
	gl.DeleteShader(this.fragmentShader)
}

func (this *DemoApp) ToCaption() string {
	return "(staticmesh)"
}


func getFileContents(path string) string {
	var content, err = ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func createShader(shaderType uint32, path string) (uint32, error) {
	var temp uint32 = gl.CreateShader(shaderType)

	var source = getFileContents(path)

	c_source, free := gl.Strs(source + "\x00")

	gl.ShaderSource(temp, 1, c_source, nil)
	free()

	gl.CompileShader(temp)

	var status int32
	gl.GetShaderiv(temp, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var len int32
		gl.GetShaderiv(temp, gl.INFO_LOG_LENGTH, &len)
		log := strings.Repeat("\x00", int(len+1))
		gl.GetShaderInfoLog(temp, len, nil, gl.Str(log))
		return 0, fmt.Errorf("Failed to compile %v: %v", source, log)
	}

	return temp, nil
}

func createProgram(shaders ...uint32) (uint32, error) {
	var temp uint32 = gl.CreateProgram()

	for _, shader := range shaders {
		gl.AttachShader(temp, shader)
	}

	gl.LinkProgram(temp)

	var status int32
	gl.GetProgramiv(temp, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var len int32
		gl.GetProgramiv(temp, gl.INFO_LOG_LENGTH, &len)
		log := strings.Repeat("\x00", int(len + 1))
		gl.GetProgramInfoLog(temp, len, nil, gl.Str(log))
		return 0, fmt.Errorf("Failed to link program: %v", log)
	}

	return temp, nil
}

func deleteProgram(program uint32, shaders ...uint32) {
	for _, shader := range shaders {
		gl.DetachShader(program, shader)
	}
	gl.DeleteProgram(program)
}

func addFloat2f(arr []float32, x float32, y float32) []float32 {
	arr = append(arr, x)
	arr = append(arr, y)
	return arr
}

func addFloat3f(arr []float32, x float32, y float32, z float32) []float32 {
	arr = append(arr, x)
	arr = append(arr, y)
	arr = append(arr, z)
	return arr
}

func addFloat4f(arr []float32, x float32, y float32, z float32, w float32) []float32 {
	arr = append(arr, x)
	arr = append(arr, y)
	arr = append(arr, z)
	arr = append(arr, w)
	return arr
}

func addUint3(arr []uint32, x uint32, y uint32, z uint32) []uint32 {
	arr = append(arr, x)
	arr = append(arr, y)
	arr = append(arr, z)
	return arr
}