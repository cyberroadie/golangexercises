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
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 2048, 2048
		subpixels              = 2
	)

	valueArray := [][]color.Color{}

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		row := []color.Color{}
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			row = append(row, mandelbrot(z))
		}
		valueArray = append(valueArray, row)
	}

	img := image.NewRGBA(image.Rect(0, 0, width/2, height/2))
	for py := 0; py < height; py = py + 2 {
		for px := 0; px < width; px = px + 2 {

			ac := avarageColor(valueArray[py][px], valueArray[py+1][px+0], valueArray[py+0][px+1], valueArray[py+1][px+1])

			img.Set(px/2, py/2, ac)
		}
	}

	file, _ := os.Create("mandlebrot.png")
	defer file.Close()
	w := bufio.NewWriter(file)
	err := png.Encode(w, img) // NOTE: ignoring errors
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}

func avarageColor(color1, color2, color3, color4 color.Color) color.Color {
	r1, g1, b1, _ := color1.RGBA()
	r2, g2, b2, _ := color2.RGBA()
	r3, g3, b3, _ := color3.RGBA()
	r4, g4, b4, _ := color4.RGBA()

	ac := color.RGBA{
		uint8((r1 + r2 + r3 + r4) / 4),
		uint8((g1 + g2 + g3 + g4) / 4),
		uint8((b1 + b2 + b3 + b4) / 4),
		255,
	}

	return ac
}

func mandelbrot(z complex128) color.Color {
	const iterations = 255
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		var r, g, b uint8
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			t := 255 - contrast*n
			if t < 85 {
				r = 255 - t
				g = t
				b = r - g
				if b < 0 {
					b = -b
				}
			} else if t > 85 && t < 170 {
				g = t
				b = 255 - t
				r = g - b
				if r < 0 {
					r = -r
				}
			} else {
				r = 255 - t
				b = 170 + t
				g = r - b
				if g < 0 {
					g = -g
				}
			}

			return color.RGBA{r, g, b, 255}
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
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
