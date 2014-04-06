package api

import (
	"github.com/jacobmoe/gorg"
	"github.com/martini-contrib/render"
	"path/filepath"
)

func PostIndex(render render.Render) {
	inPath, _ := filepath.Abs("test_files/test.org")

	tree := gorg.TreeFromFile(inPath)

	render.JSON(200, tree)
}
