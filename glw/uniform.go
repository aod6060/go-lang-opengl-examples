package glw


import "github.com/go-gl/gl/v4.1-compatibility/gl"

import lm "derf.space/gldemo/linearmath"

type Uniform struct {
	program *Program
	uniforms map[string]int32
}

func (this *Uniform) Init(program *Program) {
	this.program = program
}

func (this *Uniform) Release() {
	this.program = nil
}

func (this *Uniform) CreateUniform(name string) {
	this.uniforms[name] = gl.GetUniformLocation(this.program.id, gl.Str(name + "\x00"))
}

// Integer
func (this *Uniform) Set1i(name string, x int32) {
	gl.Uniform1i(this.uniforms[name], x)
}

func (this *Uniform) Set2i(name string, x int32, y int32) {
	gl.Uniform2i(this.uniforms[name], x, y)
}

func (this *Uniform) Set3i(name string, x int32, y int32, z int32) {
	gl.Uniform3i(this.uniforms[name], x, y, z)
}

func (this *Uniform) Set4i(name string, x int32, y int32, z int32, w int32) {
	gl.Uniform4i(this.uniforms[name], x, y, z, w)
}

// Floats
func (this *Uniform) Set1f(name string, x float32) {
	gl.Uniform1f(this.uniforms[name], x)
}

func (this *Uniform) Set2f(name string, x float32, y float32) {
	gl.Uniform2f(this.uniforms[name], x, y)
}

func (this *Uniform) Set3f(name string, x float32, y float32, z float32) {
	gl.Uniform3f(this.uniforms[name], x, y, z)
}

func (this *Uniform) Set4f(name string, x float32, y float32, z float32, w float32) {
	gl.Uniform4f(this.uniforms[name], x, y, z, w)
}

// Matrix
func (this *Uniform) SetMatrix2(name string, m *lm.Mat2) {
	gl.UniformMatrix2fv(this.uniforms[name], 1, false, &m.ToArray()[0])
}

func (this *Uniform) SetMatrix3(name string, m *lm.Mat3) {
	gl.UniformMatrix3fv(this.uniforms[name], 1, false, &m.ToArray()[0])
}

func (this *Uniform) SetMatrix4(name string, m *lm.Mat4) {
	gl.UniformMatrix4fv(this.uniforms[name], 1, false, &m.ToArray()[0])
}

