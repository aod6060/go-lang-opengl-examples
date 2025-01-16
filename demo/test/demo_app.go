package test



//import "fmt"

import sdl "github.com/veandco/go-sdl2/sdl"
import "github.com/go-gl/gl/v4.1-compatibility/gl"


//import "derf.space/gldemo/input"

type DemoApp struct {
	// This is were any variable I'll be needing
}

func (this *DemoApp) Init() {
	gl.Enable(gl.DEPTH_TEST)
}

func (this *DemoApp) HandleEvent(e sdl.Event) {

}

func (this *DemoApp) Update(delta float32) {
	
}

func (this *DemoApp) Render() {
	gl.ClearColor(1.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	// Where other gl code will go :)
}

func (this *DemoApp) Release() {

}

func (this *DemoApp) ToCaption() string {
	return "(test)"
}