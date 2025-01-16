package glw

import "fmt"
import "io/ioutil"
import "strings"
import "github.com/go-gl/gl/v4.1-compatibility/gl"

type Shader struct {
	id uint32
}


func getFileContents(path string) string {
	var content, err = ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func (this *Shader) Init(shaderType uint32, path string) error {
	this.id = gl.CreateShader(shaderType)

	var source string = getFileContents(path)

	c_source, free := gl.Strs(source + "\x00")

	gl.ShaderSource(this.id, 1, c_source, nil)
	free()

	gl.CompileShader(this.id)

	var status int32

	gl.GetShaderiv(this.id, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var len int32
		gl.GetShaderiv(this.id, gl.INFO_LOG_LENGTH, &len)
		log := strings.Repeat("\x00", int(len+1))
		gl.GetShaderInfoLog(this.id, len, nil, gl.Str(log))
		return fmt.Errorf("Failed to compile %v: %v", source, log)
	}

	return nil
}

func (this *Shader) GetID() uint32 {
	return this.id
}

func (this *Shader) Release() {
	gl.DeleteShader(this.id)
}