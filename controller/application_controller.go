package controller

import (
	// "github.com/martini-contrib/render"
	"github.com/jacobmoe/render"
)

func Home(render render.Render) {
	render.HTML(200, "home", "")
}
