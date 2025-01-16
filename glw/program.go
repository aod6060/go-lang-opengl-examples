package glw

import "fmt"
import "strings"

import "github.com/go-gl/gl/v4.1-compatibility/gl"

type Program struct {
	id uint32
	shaders []*Shader
	vertexArray *VertexArray
	uniform *Uniform
}

func (this *Program) Init(shaders ...*Shader) error {
	this.id = gl.CreateProgram()

	for _, shader := range shaders {
		gl.AttachShader(this.id, shader.GetID())
		this.shaders = append(this.shaders, shader)
	}

	gl.LinkProgram(this.id)

	var status int32
	gl.GetProgramiv(this.id, gl.LINK_STATUS, &status)

	if status == gl.FALSE {
		var len int32
		gl.GetProgramiv(this.id, gl.INFO_LOG_LENGTH, &len)
		log := strings.Repeat("\x00", int(len + 1))
		gl.GetProgramInfoLog(this.id, len, nil, gl.Str(log))
		return fmt.Errorf("Failed to link program: %v", log)
	}

	this.vertexArray.Init()

	this.uniform.Init(this)

	return nil
}

func (this *Program) Release() {
	this.uniform.Release()
	this.vertexArray.Release()

	for _, shader := range this.shaders {
		gl.DetachShader(this.id, shader.GetID())
	}

	gl.DeleteProgram(this.id)
}

func (this *Program) Bind() {
	gl.UseProgram(this.id)
}

func (this *Program) Unbind() {
	gl.UseProgram(0)
}

func (this *Program) GetVertexArray() *VertexArray {
	return this.vertexArray
}

func (this *Program) GetUniform() *Uniform {
	return this.uniform
}

func (this *Program) GetID() uint32 {
	return this.id
}