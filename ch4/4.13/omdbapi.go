package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type OmDb struct {
	Title  string
	Poster string
}

const omURL = "http://www.omdbapi.com/"

func main() {
	q := url.QueryEscape(os.Args[1])
	resp, err := http.Get(omURL + "?t=" + q)
	if err != nil {
		log.Fatalf("%s", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("%s", err)
	}

	var result OmDb
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("%s", err)
	}

	resp2, err := http.Get(result.Poster)
	if err != nil {
		log.Fatalf("%s", err)
	}

	if resp2.StatusCode != http.StatusOK {
		log.Fatalf("%s", err)
	}

	f, err := os.Create(result.Title + ".jpg")
	if err != nil {
		log.Fatalf("%s", err)
	}

	_, err = io.Copy(f, resp2.Body)
	if err != nil {
		log.Fatalf("%s", err)
	}
	f.Close()
	fmt.Printf("Downloaded and saved file: " + result.Title + ".jpg")

}
