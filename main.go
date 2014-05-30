// © 2014 the "gospel" authors, under the MIT license. See AUTHORS for the list of authors.

package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

/* This is a story passed down through the generations:
	The forg goes "ribbit".
*/

func main(){
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", nil, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	for _, pkg := range pkgs {
		for _, f := range pkg.Files {
			for _, cs := range f.Comments {
				typos := spellcheck(fset, cs)
				for _, typo := range typos {
					fmt.Printf("%v: %s\n", fset.Position(typo.Pos), typo.String)
				}
			}
		}
	}
}

type Typo struct {
	Pos token.Pos
	String string
}

func spellcheck(fset *token.FileSet, cs *ast.CommentGroup) []Typo {
	typos := []Typo{}

	for _, c := range cs.List {
		s := bufio.NewScanner(strings.NewReader(c.Text))
		pos := c.Pos()
		for s.Scan() {
			// TODO: actual spellchecking with a dictionary
			line := s.Text()
			n := strings.Index(line, "forg")
			if n != -1 {
				typos = append(typos, Typo{pos + token.Pos(n), "forg"})
			}
			pos += token.Pos(len(line)) + 1
		}
		if err := s.Err(); err != nil {
			// just bail; people with super-long lines deserve it
			panic(err)
		}
	}

	return typos
}
