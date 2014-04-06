package model

import "github.com/jacobmoe/gorg"

type Page struct {
	Title string  `json:"title"`
	Posts []*Post `json:"posts"`
}

// gorg made no assumptions about the structure of the org-mode file
// here, we are assuming:
// - every subtree contains at most 3 levels
// - the page name is a level-1 headline (*)
// - the title of the post is a level-2 headline (**)
// - the date of the post is a level-3 headline (***)
// - the body of the post is contained under the date
func GetPagesFromTree(tree *gorg.Tree, from int, to int) []*Page {
	var pages []*Page

	page := &Page{}
	post := &Post{}

	for _, subtree := range tree.Subtrees {

		for _, node := range subtree.Nodes {

			switch node.Position {
			case 1:
				page = &Page{Title: subtree.Nodes[0].Headline}
			case 2:
				post.Title = node.Headline
			case 3:
				post.Date = node.Headline
				post.Section = node.Section

				page.Posts = append(page.Posts, post)
				pages = append(pages, page)
			}

		}
	}

	return pages
}
