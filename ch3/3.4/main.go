// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

const (
	width, height = 1200, 640           // canvas size in pixels
	cells         = 400                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	fill          = "#FF0000"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			graph(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	file, _ := os.Create("graph.svg")
	defer file.Close()
	w := bufio.NewWriter(file)
	graph(w)

}

func graph(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _, err := corner(i+1, j)
			if err != nil {
				continue
			}

			bx, by, _, err := corner(i, j)
			if err != nil {
				continue
			}

			cx, cy, _, err := corner(i, j+1)
			if err != nil {
				continue
			}

			dx, dy, fill, err := corner(i+1, j+1)
			if err != nil {
				continue
			}

			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: "+fill+"'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (float64, float64, string, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, err := f(x, y)
	if err != nil {
		return 0, 0, "", err
	}

	var fill string
	if z < 0 {
		fill = "#0000FF"
	} else {
		fill = "#FF0000"
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, fill, nil
}

func f(x, y float64) (z float64, err error) {
	r := math.Hypot(x, y) // distance from (0,0)
	z = math.Sin(r) / r
	if math.IsNaN(z) {
		err = errors.New("NaN")
	}

	return z, err

}

//!-
