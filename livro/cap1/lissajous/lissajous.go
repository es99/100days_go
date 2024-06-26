// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	// "os"
)

var palette = []color.Color{color.Black, color.RGBA{R: 0, G: 255, B: 0, A: 255}, color.RGBA{R: 0, G: 0, B: 255, A: 255}, color.RGBA{R: 255, G: 0, B: 0, A: 255}}

/*
const (
	redIndex   = 0
	blackIndex = 1 // first color in palette
	greenIndex = 2 // next color in palette
	blueIndex  = 3
)


func main() {
	lissajous(os.Stdout)
}
*/

func Lissajous(out io.Writer, ciclos int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	cycles := float64(ciclos)    // number of complete x oscillator revolutions
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		colorIndex := rand.Intn(4) + 1
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(colorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
