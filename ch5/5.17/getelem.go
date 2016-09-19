package main

import (
	"log"
	"net/http"
	"os"

	"fmt"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalf("%s", err)
	}

	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("%s", err)
	}

	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	for _, i := range images {
		fmt.Printf("%v\n", i)
	}

	for _, h := range headings {
		fmt.Printf("%v\n", h)
	}

}

// ElementsByTagName find elements of specific type
func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	return elementsByTagName([]*html.Node{}, doc.FirstChild, names...)
}

func elementsByTagName(nodes []*html.Node, n *html.Node, names ...string) []*html.Node {
	if n == nil {
		return nodes
	}

	if n.Type == html.ElementNode {
		for _, ns := range names {
			if n.Data == ns {
				nodes = append(nodes, n)
				break
			}
		}
	}

	if n.NextSibling == nil {
		return elementsByTagName(nodes, n.FirstChild, names...)
	}

	return elementsByTagName(nodes, n.NextSibling, names...)
}
