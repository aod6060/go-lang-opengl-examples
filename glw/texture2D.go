package glw

import "unsafe"
import "github.com/go-gl/gl/v4.1-compatibility/gl"
import sdl "github.com/veandco/go-sdl2/sdl"
import img "github.com/veandco/go-sdl2/img"

type Texture2D struct {
	id uint32
	width int32
	height int32
}

func (this *Texture2D) Init() {
	gl.GenTextures(1, &this.id)
}

func (this *Texture2D) Release() {
	gl.DeleteTextures(1, &this.id)
}

func (this *Texture2D) Bind(activeTex uint32) {
	gl.ActiveTexture(activeTex)
	gl.BindTexture(gl.TEXTURE_2D, this.id)
}

func (this *Texture2D) Unbind(activeTex uint32) {
	gl.ActiveTexture(activeTex)
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (this *Texture2D) TexImage2D(level int32, interalFormat int32, width int32, height int32, format uint32, valueType uint32, pixels unsafe.Pointer) {
	gl.TexImage2D(
		gl.TEXTURE_2D,
		level,
		interalFormat,
		width,
		height,
		0,
		format,
		valueType,
		pixels)
}

func (this *Texture2D) TexParameter(pname uint32, param int32) {
	gl.TexParameteri(gl.TEXTURE_2D, pname, param)
}

func (this *Texture2D) GenerateMipmap() {
	gl.GenerateMipmap(gl.TEXTURE_2D)
}

func (this *Texture2D) GetID() uint32 {
	return this.id
}

func (this *Texture2D) GetWidth() int32 {
	return this.width
}

func (this *Texture2D) GetHeight() int32 {
	return this.height
}

func CreateTexture2DFromFile(out *Texture2D, path string) {
	var surf *sdl.Surface
	var err error

	surf, err = img.Load(path)

	if err != nil {
		panic(err)
	}

	out.width = surf.W
	out.height = surf.H

	out.Init()

	out.Bind(gl.TEXTURE0)

	out.TexImage2D(
		0,
		gl.RGBA,
		out.width,
		out.height,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(surf.Pixels()))

	out.TexParameter(gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	out.TexParameter(gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)

	out.GenerateMipmap()

	out.Unbind(gl.TEXTURE0)

	surf.Free()
}