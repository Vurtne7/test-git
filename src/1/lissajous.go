package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{
	//color.RGBA{235, 207, 239, 1},
	color.RGBA{0x22, 0xCC, 0x33, 0xff},
	color.White,
	color.Black,
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w, r)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:9777", nil))
		return
	}

	lissajous(os.Stdout, nil)
}

func lissajous(out io.Writer, r *http.Request) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	var cycles float64 = 5
	if r != nil {
		if err := r.ParseForm(); err == nil {
			if intCycles, err := strconv.Atoi(r.Form.Get("cycles")); err == nil {
				cycles = float64(intCycles)
			}
		}
	}
	log.Printf("get cycles: %f \n", cycles)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 2)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
