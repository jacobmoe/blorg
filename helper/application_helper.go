package helper

import "github.com/jacobmoe/gorg"
import "path/filepath"

func Application() map[string]interface{} {

	helper := make(map[string]interface{})

	helper["sectionTitles"] = func() []string {
		return sectionTitles()
	}

	return helper
}

func sectionTitles() []string {
	inPath, _ := filepath.Abs("test_files/test.org")
	tree := gorg.TreeFromFile(inPath)
	var sectionTitles []string

	for _, subtree := range tree.Subtrees {
		sectionTitles = append(sectionTitles, subtree.Nodes[0].Headline)
	}

	return sectionTitles
}
