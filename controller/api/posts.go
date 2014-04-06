package api

import (
	"github.com/jacobmoe/blorg/model"
	"github.com/jacobmoe/gorg"
	"github.com/martini-contrib/render"
	"path/filepath"
)

func PostIndex(render render.Render) {
	inPath, _ := filepath.Abs("test_files/test.org")

	tree := gorg.TreeFromFile(inPath)

	pages := model.GetPagesFromTree(tree, 0, 0)

	render.JSON(200, pages)
}
