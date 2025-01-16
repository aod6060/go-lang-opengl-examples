package blank

import sdl "github.com/veandco/go-sdl2/sdl"
import "github.com/go-gl/gl/v4.1-compatibility/gl"

type DemoApp struct {

}

func (this *DemoApp) Init() {
	gl.Enable(gl.DEPTH_TEST)
}

func (this *DemoApp) HandleEvent(e sdl.Event) {

}

func (this *DemoApp) Update(delta float32) {

}

func (this *DemoApp) Render() {
	gl.ClearColor(0.0, 0.0, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (this *DemoApp) Release() {
}

func (this *DemoApp) ToCaption() string {
	return "(blank)"
}