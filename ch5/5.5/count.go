package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"net/http"

	"log"

	"golang.org/x/net/html"
	"gopkg.in/urfave/cli.v1"
)

type webpage struct {
	url       string
	imgCount  int
	wordCount int
}

var wp webpage

func main() {

	app := cli.NewApp()
	app.Action = countWordsAndImages
	app.Name = "count"
	app.Usage = " count images and words in a webpage"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "url, u",
			Usage:       "url to web page for counting images and words",
			Destination: &wp.url,
		},
	}
	app.Run(os.Args)

}

func countWordsAndImages(c *cli.Context) {
	resp, err := http.Get(wp.url)
	if err != nil {
		log.Fatalf("Error getting web page %s", err)
	}
	rootNode, err := html.Parse(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Fatalf("Error parsing html page: %s", err)
	}

	wp = visit(wp, rootNode.FirstChild)

	fmt.Println(wp)
}

func visit(wp webpage, n *html.Node) webpage {
	if n == nil {
		return wp
	}
	if n.Type == html.TextNode {

		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			wp.wordCount++
		}
	} else if n.Type == html.ElementNode {
		if n.Data == "img" {
			wp.imgCount++
		}
	}

	if n.NextSibling == nil {
		return visit(wp, n.FirstChild)
	}
	return visit(wp, n.NextSibling)
}
