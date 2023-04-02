// Lissajous генерирует анимированные GIF из случайных
// фигур Лиссажу.
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
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // Первый цвет палитры
	blackIndex = 1 // Сдежующий цвет палитры
	CYCLES     = 5
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	if _, ok := r.Form["cycles"]; !ok {
		lissajous(w, CYCLES)
	} else {
		cycles, err := strconv.Atoi(r.Form["cycles"][0])
		if err != nil {
			lissajous(w, CYCLES)
		}

		lissajous(w, float64(cycles))
	}

}

func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.01 // Угловое разрешение
		size    = 100  // Канва изображения охватывает [size..+size]
		nframes = 64   // Количество кадров анимации
		delay   = 8    // Задержка между кадрами (единица - 10мс)
	)

	rand.Seed(time.Now().UnixNano())
	freq := rand.Float64() * 3.0 // Относительная частота колебаний
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	if err := gif.EncodeAll(out, &anim); err != nil {
		log.Fatal(err)
	}
}
