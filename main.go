package main

//import "fmt"
import "os"
import "derf.space/gldemo/app"

import dtest "derf.space/gldemo/demo/test"
import "derf.space/gldemo/demo/blank"
import "derf.space/gldemo/demo/triangle"
import "derf.space/gldemo/demo/color_triangle"
import "derf.space/gldemo/demo/index_buffer"
import "derf.space/gldemo/demo/textures"
import "derf.space/gldemo/demo/staticmesh"
import "derf.space/gldemo/demo/glw_test"

func main() {
	var conf app.Config

	conf.Width = 1280
	conf.Height = 720


	if len(os.Args) > 1{
		if os.Args[1] == "blank" {
			conf.App = &blank.DemoApp{}
		} else if os.Args[1] == "triangle" {
			conf.App = &triangle.DemoApp{}
		} else if os.Args[1] == "color-triangle" {
			conf.App = &color_triangle.DemoApp{}
		} else if os.Args[1] == "index-buffer" {
			conf.App = &index_buffer.DemoApp{}
		} else if os.Args[1] == "textures" {
			conf.App = &textures.DemoApp{}
		} else if os.Args[1] == "staticmesh" {
			conf.App = &staticmesh.DemoApp{}
		} else if os.Args[1] == "glw-test" {
			conf.App = &glw_test.DemoApp{}
		} else {
			conf.App = &dtest.DemoApp{}
		}
	} else {
		conf.App = &dtest.DemoApp{}
	}

	conf.Caption = "OpenGL Demo " + conf.App.ToCaption()

	app.Init(&conf)
	app.Update()
	app.Release()
}