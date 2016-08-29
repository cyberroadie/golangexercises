package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "count elements: %v\n", err)
		os.Exit(1)
	}
	for k, v := range visit(make(map[string]int), doc.FirstChild) {
		fmt.Printf("%s: %d\n", k, v)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(elem map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return elem
	}
	if n.Type == html.ElementNode {
		elem[n.Data]++
	}
	if n.NextSibling == nil {
		return visit(elem, n.FirstChild)
	}
	return visit(elem, n.NextSibling)
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
