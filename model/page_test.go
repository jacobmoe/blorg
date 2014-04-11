package model

import (
	"fmt"
	"github.com/jacobmoe/gorg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPagesFromTree(t *testing.T) {
	fmt.Println("==== TestModelPage PagesFromTree")

	tree := &gorg.Tree{
		Subtrees: []*gorg.Tree{
			&gorg.Tree{
				Subtrees: []*gorg.Tree{
					&gorg.Tree{
						Nodes: []*gorg.Node{
							&gorg.Node{
								Headline: "Post Title 1",
								Position: 2,
							},
							&gorg.Node{
								Headline: "January 1, 2001",
								Position: 3,
								Section: []string{
									"the first post",
								},
							},
						},
					},
					&gorg.Tree{
						Nodes: []*gorg.Node{
							&gorg.Node{
								Headline: "Post Title 2",
								Position: 2,
							},
							&gorg.Node{
								Headline: "January 2, 2002",
								Position: 3,
								Section: []string{
									"the section post",
								},
							},
						},
					},
				},
				Nodes: []*gorg.Node{
					&gorg.Node{
						Headline: "Section Title 1",
						Position: 1,
					},
				},
			},
		},
	}

	expectedPages := Page{
		Title: "Section Title 1",
		Posts: []*Post{
			&Post{
				Title:   "Post Title 1",
				Section: []string{"the first post"},
				Date:    "January 1, 2001",
			},
			&Post{
				Title:   "Post Title 2",
				Section: []string{"the second post"},
				Date:    "January 2, 2002",
			},
		},
	}

	pages := PagesFromTree(tree)

	assert.Equal(t, pages, expectedPages)
}
