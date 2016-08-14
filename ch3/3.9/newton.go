// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			x, err := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
			if err != nil {
				x = 2
			}
			y, err := strconv.ParseFloat(r.URL.Query().Get("y"), 64)
			if err != nil {
				y = 2
			}

			w.Header().Set("Content-Type", "image/png")
			pngimage(w, x, y)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	file, _ := os.Create("newton_color.png")
	defer file.Close()
	w := bufio.NewWriter(file)
	pngimage(w, 2, 2)

}

func pngimage(w io.Writer, xmax, ymax float64) {

	const (
		width, height = 2048, 2048
	)

	xmin := -xmax
	ymin := -ymax

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}

	err := png.Encode(w, img) // NOTE: ignoring errors
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7

	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return intToColor(i, 37)
		}
	}
	return color.Black
}

func intToColor(n, max uint8) color.Color {
	sic := 0xffffff / uint32(max)
	r := 255 - uint8((uint32(n)*sic)>>16)
	g := 255 - uint8(((uint32(n) * sic) >> 8))
	b := 255 - uint8((uint32(n) * sic))

	return color.RGBA{r, g, b, 255}
}
