package glw



import "github.com/go-gl/gl/v4.1-compatibility/gl"

type VertexArray struct {
	id uint32
	attributes map[string]uint32
}

func (this *VertexArray) Init() {
	gl.GenVertexArrays(1, &this.id)
}

func (this *VertexArray) Release() {
	gl.DeleteVertexArrays(1, &this.id)
}

func (this *VertexArray) Bind() {
	gl.BindVertexArray(this.id)
}

func (this *VertexArray) Unbind() {
	gl.BindVertexArray(0)
}

func (this *VertexArray) CreateAttribute(name string, id uint32) {
	this.attributes[name] = id
}

func (this *VertexArray) Enable(name string) {
	gl.EnableVertexAttribArray(this.attributes[name])
}

func (this *VertexArray) Disable(name string) {
	gl.DisableVertexAttribArray(this.attributes[name])
}

func (this *VertexArray) Pointer(name string, size int32, valueType uint32) {
	gl.VertexAttribPointer(this.attributes[name], size, valueType, false, 0, nil)
}

func (this *VertexArray) GetID() uint32 {
	return this.id
}