// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	d := depth(0)

	//!+call
	forEachNode(doc, startElement, endElement, d)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node, d func(int) int), depth func(int) int) {
	if pre != nil {
		pre(n, depth)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, depth)
	}

	if post != nil {
		post(n, depth)
	}
}

//!-forEachNode

//!+startend
func depth(l int) func(int) int {
	var d int
	return func(l int) int {
		d = d + l
		return d
	}
}

func startElement(n *html.Node, depth func(int) int) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth(1)*2-2, "", n.Data)
	}
}

func endElement(n *html.Node, depth func(int) int) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth(-1)*2, "", n.Data)
	}
}

//!-startend
