// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	const tpl = `
		<!DOCTYPE html>
		<html>
			<head>
				<meta charset="UTF-8">
			</head>
			<style>
			table, th, td {
				border: 1px solid black;
			}
			</style>
			<body>
				<table>
				{{ range $key, $value := . }}
				<tr>
					<td>{{ $key }}</td><td>{{ $value }}</td>
				</tr>
				{{ end }}
				</table>
			</body>
		</html>`

	check := func(err error) {
		if err != nil {
			w.WriteHeader(http.StatusTeapot)
			return
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	err = t.Execute(w, db)
	check(err)

}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if item == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad request item: %q price: %q\n", item, price)
	} else if price, ok := db[item]; ok {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "item already in database item: %q price: %q\n", item, price)
	} else {
		db[item] = price
		fmt.Fprintf(w, "item added to database item: %q price: %q\n", item, price)
	}

}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad request: item: %q price: %q\n", item, price)
	} else if price, ok := db[item]; ok {
		fmt.Fprintf(w, "item: %q price: %q\n", item, price)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad request: item: %q price: %q\n", item, price)
	} else if price, ok := db[item]; ok {
		db[item] = price
		fmt.Fprintf(w, "item updated in database: item: %q price: %q\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad request: item: %q \n", item)
	} else if price, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "item deleted from database: item: %q price: %q\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}

}
