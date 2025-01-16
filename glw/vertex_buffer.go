package glw

import "unsafe"
import "github.com/go-gl/gl/v4.1-compatibility/gl"

type VertexBuffer struct {
	id uint32
	list []float32
}

func (this *VertexBuffer) Init() {
	gl.GenBuffers(1, &this.id)
}

func (this *VertexBuffer) Release() {
	gl.GenBuffers(1, &this.id)
}

func (this *VertexBuffer) Update() {
	this.Bind()
	gl.BufferData(gl.ARRAY_BUFFER, len(this.list) * int(unsafe.Sizeof(float32(0))), gl.Ptr(this.list), gl.DYNAMIC_DRAW)
	this.Unbind()
}

func (this *VertexBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, this.id)
}

func (this *VertexBuffer) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (this *VertexBuffer) Add1(x float32) {
	this.list = append(this.list, x)
}

func (this *VertexBuffer) Add2(x float32, y float32) {
	this.list = append(this.list, x)
	this.list = append(this.list, y)
}

func (this *VertexBuffer) Add3(x float32, y float32, z float32) {
	this.list = append(this.list, x)
	this.list = append(this.list, y)
	this.list = append(this.list, z)
}

func (this *VertexBuffer) Add4(x float32, y float32, z float32, w float32) {
	this.list = append(this.list, x)
	this.list = append(this.list, y)
	this.list = append(this.list, z)
	this.list = append(this.list, w)
}

func (this *VertexBuffer) Clear() {
	this.list = this.list[:0]
}

func (this *VertexBuffer) GetID() uint32 {
	return this.id
}