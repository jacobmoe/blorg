Blorg
=====

_Blog-Org_ 

A simple [Backbone](http://backbonejs.org/) site served by a Go / [Martini](https://github.com/go-martini/martini) back-end. 
All content is managed with an emacs [org-mode](http://orgmode.org/) file.

## .org file 

### Structure

```
* page 1
** post 1
*** a date
the body of a post with a date.
** post 2
* page 2
** post 1
the body of a post without a date.
* page 3
the body of a page without post titles.
```

For now, content deeper than 3 asterisks is ignored. 

### Markdown

Standard markdown - implemented with [blackfriday](https://github.com/russross/blackfriday) - can be used in a post's body. Some exceptions:

#### Code blocks

The org-mode code block is used rather than the usual markdown syntax. 
Code is nested between #+BEGIN_SRC and #+END_SRC.

```
#+BEGIN_SRC go

package example

import "fmt"

func saySomething() {
  fmt.Println("org-mode rulez!")
}
 
#+END_SRC
```

#### Tables

If the first character of a line is a pipe, it will be interpreted as a table, both by org-mode and blorg. Org-mode tables are super cool. 

```
|   col1 |      col2 |  col3 |
|--------+-----------+-------+
|    111 |       111 |   111 |
|    111 |       111 |   111 |
```

Becomes:

<table>
  <tr>
    <th>col1</th>
    <th>col2</th>
    <th>col3</th>
  </tr>
  <tr>
    <td>111</td>
    <td>111</td>
    <td>111</td>
  </tr>
  <tr>
    <td>111</td>
    <td>111</td>
    <td>111</td>
  </tr>
</table>

## Bugs/Todo

- When there is a single page (one single-asterisk headline), pages become the two-asterisk healines nested under it. Could be a feature?
- If the last line of the file is a table, that table is ignored. This is an issue in `appendSections`, which needs to be cleaned up anyway. 
- Figure out a publishing system. Maybe use Dropbox.
- Cache parsed .org file in a .json file on publish. 
- Should generalize posts to allow deeper org-mode structures.
- The Backbone `Page` view should be broken up into collections of `Posts`.
- Test!
- Styles!
