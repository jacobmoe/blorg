package controller

import (
	"github.com/go-martini/martini"
	"github.com/jacobmoe/gorg"
	"github.com/martini-contrib/render"
	"io/ioutil"
	"path/filepath"
)

func Home(render render.Render) {
	render.HTML(200, "home", "THE BLORG")
}

func Blog() string {
	inPath, _ := filepath.Abs("test/test.org")
	outPath, _ := filepath.Abs("test/test.json")

	gorg.OrgToHtmlFile(inPath, outPath)

	html, _ := ioutil.ReadFile(outPath)

	return string(html)
}

func Name(args martini.Params) string {
	return "<h1>" + args["name"] + "</h1>"
}
