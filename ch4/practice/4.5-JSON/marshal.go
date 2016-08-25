package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type movie struct {
	Title  string
	Year   int `json:"released"`
	Actors []string
}

var movies = []movie{
	{Title: "Casablanca", Year: 1942, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Actors: []string{"Paul Newman"}},
}

func main() {
	// Marshal JSON
	data, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatalf("1 %s", err)
	}
	fmt.Printf("%s\n", data)

	// Write to file
	f, err := os.Create("tmp.json")
	if err != nil {
		log.Fatalf("2 %s", err)
	}
	w := bufio.NewWriter(f)
	w.WriteString(string(data))
	w.Flush()
	f.Close()

	// Read file
	f2, err := os.Open("tmp.json")

	b2, err := ioutil.ReadAll(f2)
	if err != nil {
		log.Fatalf("3 %s", err)
	}

	// Unmarshal
	var titles []struct{ Title string }

	if err := json.Unmarshal(b2, &titles); err != nil {
		log.Fatalf("4 %s", err)
	}

	fmt.Printf("%v", titles)
}
