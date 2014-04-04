package controller

import (
	"github.com/codegangsta/martini"
	"github.com/jacobmoe/gorg"
	"io/ioutil"
	"path/filepath"
)

func Blog() string {
	inPath, _ := filepath.Abs("test/test.org")
	outPath, _ := filepath.Abs("test/test.html")

	gorg.OrgToHtmlFile(inPath, outPath)

	html, _ := ioutil.ReadFile(outPath)

	return string(html)
}

func Name(args martini.Params) string {
	return "<h1>" + args["name"] + "</h1>"
}
