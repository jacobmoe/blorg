package helper

import (
	"github.com/jacobmoe/blorg/controller/api"
	"github.com/jacobmoe/gorg"
	"path/filepath"
	"regexp"
	"strings"
)

func Application() map[string]interface{} {

	helper := make(map[string]interface{})

	helper["siteTitle"] = func() string {
		ext := filepath.Ext(api.OrgFileName)
		return strings.Replace(api.OrgFileName, ext, "", -1)
	}

	helper["sectionTitles"] = func() []string {
		return sectionTitles()
	}

	helper["slugify"] = func(title string) string {
		return slugify(title)
	}

	helper["plus1"] = func(i int) int {
		return i + 1
	}

	helper["titleize"] = func(s string) string {
		r := regexp.MustCompile("[_-]")

		return r.ReplaceAllLiteralString(strings.Title(s), " ")
	}

	helper["toUpper"] = func(s string) string {
		return strings.ToUpper(s)
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
