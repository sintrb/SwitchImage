package main

import (
	"code.google.com/p/graphics-go/graphics"
	"fmt"
	"image"
	// "image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"net/http"
)

func FitSize(src image.Image, w, h int) (image.Image, error) {
	dst := image.NewRGBA(image.Rect(0, 0, w, h))

	db := dst.Bounds()
	sb := src.Bounds()
	rx := float64(sb.Dx()) / float64(db.Dx())
	ry := float64(sb.Dy()) / float64(db.Dy())
	var b image.Rectangle
	if rx > ry {
		b = image.Rect(0, 0, db.Dx(), int(float64(sb.Dy())/rx))
	} else {
		b = image.Rect(0, 0, int(float64(sb.Dx())/ry), db.Dy())
	}
	sx := (db.Dx() - b.Dx()) / 2
	sy := (db.Dy() - b.Dy()) / 2
	ndb := image.Rect(sx, sy, db.Dx()-sx, db.Dy()-sy)
	ms := image.NewRGBA(image.Rect(0, 0, ndb.Dx(), ndb.Dy()))
	graphics.Scale(ms, src)

	draw.Draw(dst, ndb, ms, image.ZP, draw.Src)
	return dst, nil
}

func GetImage(url string) (string, image.Image) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var img image.Image
	ctype := resp.Header.Get("Content-Type")
	if ctype == "image/jpeg" {
		img, err = jpeg.Decode(resp.Body)
	} else if ctype == "image/png" {
		img, err = png.Decode(resp.Body)
	} else {
		panic("Content-Type Error: " + ctype)
	}
	img, _ = FitSize(img, 360, 200)
	return ctype, img
}

func onerr(w http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		w.Write([]byte("Error"))
		fmt.Println(err)
	}
}

func tranimg(w http.ResponseWriter, r *http.Request) {
	defer onerr(w, r)

	r.ParseForm()
	url := r.Form.Get("url")
	if url != "" {
		fmt.Println("image: " + url)
		_, m := GetImage(url)
		w.Header().Add("Content-Type", "image/png")
		png.Encode(w, m)
	} else {
		w.Write([]byte("image url is necessary, example: ?url=http://img3.douban.com/lpic/s6037735.jpg&width=360&height=200"))
	}
}

func main() {
	http.HandleFunc("/", tranimg)

	http.ListenAndServe(":9999", nil)
	fmt.Printf("ok\n")
}
