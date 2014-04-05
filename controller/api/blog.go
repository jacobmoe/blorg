package api

import (
	"github.com/jacobmoe/gorg"
	"github.com/martini-contrib/render"
	"io/ioutil"
	"path/filepath"
)

func BlogShow(render render.Render) {
	inPath, _ := filepath.Abs("test_files/test.org")
	outPath, _ := filepath.Abs("test_files/test.json")

	gorg.OrgToJsonFile(inPath, outPath)

	json, _ := ioutil.ReadFile(outPath)

	render.JSON(200, string(json))
}
