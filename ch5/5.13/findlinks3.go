// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"os"

	"net/http"

	"gopl.io/ch5/links"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(urls []string) func(url string) []string {
	// create map of domains to compare against
	origUrls := make(map[string]bool)

	for _, us := range urls {
		u, _ := url.Parse(us)
		origUrls[u.Host] = true
	}

	return func(us string) []string {
		u, _ := url.Parse(us)
		if origUrls[u.Host] {
			fmt.Printf("Saving: %s\n", us)
			savePage(u)
		} else {
			fmt.Printf("%s\n", us)
		}

		list, err := links.Extract(us)
		if err != nil {
			log.Print(err)
		}

		return list
	}
}

//!-crawl

func savePage(u *url.URL) {
	fmt.Printf("%s", u.RequestURI())
	response, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	path := u.Host + u.Path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.Create(path + "content.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

}

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	c := crawl(os.Args[1:])

	breadthFirst(c, os.Args[1:])
}

//!-main
