package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

/*
Exercise 1.12 --
Modify lissajous server to read parameter from the URL. for example http://localhost:8080/?cycles=20&size=999
and in that case we will use those values and pass them into the lissajous() method to change the way the gif looks
in the browser
*/

//http://localhost:8080/?cycles=20&size=999

var palette = []color.Color{color.White, color.RGBA{25, 250, 25, 25}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	cycles, err := strconv.Atoi(r.Form.Get("cycles"))
	if err != nil {
		fmt.Fprintf(w, "Missing %s !", err)
	}
	size, err := strconv.Atoi(r.Form.Get("size"))
	if err != nil {
		fmt.Fprintf(w, "Missing %s !", err)
	}
	// frames, err := strconv.Atoi(r.Form.Get("nframes"))
	// if err != nil {
	// 	fmt.Fprintf(w, "Missing %s !", err)
	// }
	lissajous(w, float64(cycles), uint(size))

}

func lissajous(out io.Writer, cycls float64, s uint) {
	cycles := cycls
	const (
		res     = 0.001
		size    = 100
		nFrames = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nFrames}
	phase := 0.0
	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
