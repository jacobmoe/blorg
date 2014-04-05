package controller

import (
	"github.com/go-martini/martini"
	"github.com/jacobmoe/gorg"
	"github.com/martini-contrib/render"
	"html/template"
	"io/ioutil"
	"path/filepath"
)

func Home(render render.Render) {
	render.HTML(200, "home", "THE BLORG")
}

func Blog(render render.Render) {
	inPath, _ := filepath.Abs("test_files/test.org")
	outPath, _ := filepath.Abs("test_files/test.html")

	gorg.OrgToHtmlFile(inPath, outPath)

	html, _ := ioutil.ReadFile(outPath)

	render.HTML(200, "blog", template.HTML(string(html)))
}

func Name(args martini.Params) string {
	return "<h1>" + args["name"] + "</h1>"
}
