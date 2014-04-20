package helper

import "github.com/jacobmoe/gorg"
import "github.com/jacobmoe/blorg/controller/api"
import "path/filepath"
import "strings"
import "regexp"

func Application() map[string]interface{} {

	helper := make(map[string]interface{})

	helper["sectionTitles"] = func() []string {
		return sectionTitles()
	}

	helper["slugify"] = func(title string) string {
		return slugify(title)
	}

	helper["plus1"] = func(i int) int {
		return i + 1
	}

	return helper
}

func sectionTitles() []string {
	inPath, _ := filepath.Abs(api.OrgFilePath)
	tree := gorg.TreeFromFile(inPath)
	var sectionTitles []string

	for _, subtree := range tree.Subtrees {
		sectionTitles = append(sectionTitles, subtree.Nodes[0].Headline)
	}

	return sectionTitles
}

// string cleanup for slug
func slugify(s string) string {
	re := regexp.MustCompile("[^A-Za-z0-9 ]")

	s = re.ReplaceAllString(s, "")
	s = strings.Trim(s, " ")
	s = strings.Replace(s, " ", "-", -1)
	s = strings.ToLower(s)

	return s
}
