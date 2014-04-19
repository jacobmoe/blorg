package model

import "github.com/jacobmoe/gorg"
import "github.com/russross/blackfriday"
import "regexp"
import "strings"

type Page struct {
	Title string  `json:"title"`
	Posts []*Post `json:"posts"`
}

// gorg made no assumptions about the structure of the org-mode file
// here, we are assuming:
// - the page name is a level-1 headline (*)
// - the title of the post is a level-2 headline (**)
// - the date of the post is a level-3 headline (***)
// - the body of the post is collected from all sections
func PagesFromTree(tree *gorg.Tree) []*Page {
	var pages []*Page

	for _, subtree := range tree.Subtrees {
		page := &Page{}
		post := &Post{}

		var sections []string

		for _, node := range subtree.Nodes {

			// if there is a single post, this is as deep as the tree gets
			switch node.Position {
			case 1:
				page = &Page{Title: subtree.Nodes[0].Headline}
			case 2:
				post.Title = node.Headline
			case 3:
				post.Date = node.Headline
			}

			sections = appendSections(sections, node.Section)

			for _, sst := range subtree.Subtrees {
				post := &Post{}

				for _, n := range sst.Nodes {
					switch n.Position {
					case 2:
						post.Title = n.Headline
					case 3:
						post.Date = n.Headline
					}

					sections = appendSections(sections, n.Section)
				}

				post.Section = sections
				sections = []string{}
				page.Posts = append(page.Posts, post)
			}
		}

		post.Section = sections
		page.Posts = append(page.Posts, post)
		pages = append(pages, page)
	}

	return pages
}

func appendSections(sections []string, nodeSection []string) []string {
	for _, s := range nodeSection {

		r, _ := regexp.Compile(`(?s)(\#\+BEGIN_SRC)\ ([a-zA-Z]*)(.*)(\#\+END_SRC)`)
		submatches := r.FindStringSubmatch(s)

		// if section is a code block that is structured correctly,
		// we will find a match plus exactly 4 sub-matches
		if len(submatches) == 5 {
			lang := strings.TrimSpace("lang-" + submatches[2])
			html := "<pre class=\"" + lang + "\"><code>" + strings.TrimSpace(submatches[3]) + "</code></pre>"
			sections = append(sections, html)
		} else {
			section := blackfriday.MarkdownCommon([]byte(s))
			sections = append(sections, string(section))
		}
	}

	return sections
}
