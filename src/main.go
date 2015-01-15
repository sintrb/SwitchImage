package main

import (
	"fmt"
	"image"
	"image/color"
	// "image/draw"
	"image/jpeg"
	"os"
)

func main() {
	f1, err := os.Open("../imgs/s6037735.jpg")
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	m1, err := jpeg.Decode(f1)
	if err != nil {
		panic(err)
	}

	m := image.NewRGBA(image.Rect(0, 0, 360, 200))
	m.Set(100, 100, color.RGBA{255, 0, 0, 255})
	// draw.Draw(m, m.Bounds(), src, sp, op)
	f2, err := os.Create("../imgs/out.jpg")
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	err = jpeg.Encode(f2, m, &jpeg.Options{90})
	println(m1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ok\n")
}
