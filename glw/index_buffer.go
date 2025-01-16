package glw


import "unsafe"
import "github.com/go-gl/gl/v4.1-compatibility/gl"


type IndexBuffer struct {
	id uint32
	list []uint32
}

func (this *IndexBuffer) Init() {
	gl.GenBuffers(1, &this.id)
}

func (this *IndexBuffer) Release() {
	gl.GenBuffers(1, &this.id)
}

func (this *IndexBuffer) Update() {
	this.Bind()
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(this.list) * int(unsafe.Sizeof(float32(0))), gl.Ptr(this.list), gl.DYNAMIC_DRAW)
	this.Unbind()
}

func (this *IndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, this.id)
}

func (this *IndexBuffer) Unbind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

func (this *IndexBuffer) Add1(x uint32) {
	this.list = append(this.list, x)
}

func (this *IndexBuffer) Add2(x uint32, y uint32) {
	this.list = append(this.list, x)
	this.list = append(this.list, y)
}

func (this *IndexBuffer) Add3(x uint32, y uint32, z uint32) {
	this.list = append(this.list, x)
	this.list = append(this.list, y)
	this.list = append(this.list, z)
}

func (this *IndexBuffer) Add4(x uint32, y uint32, z uint32, w uint32) {
	this.list = append(this.list, x)
	this.list = append(this.list, y)
	this.list = append(this.list, z)
	this.list = append(this.list, w)
}

func (this *IndexBuffer) Clear() {
	this.list = this.list[:0]
}

func (this *IndexBuffer) GetID() uint32 {
	return this.id
}