package main

import (
	"os"
	"maze/public"
	"image/png"
	"github.com/nickdavies/go-astar/astar"
	"image"
	"image/color"
)

func main() {
	img, err := os.Open("./images/source-2.png")
	public.CheckErr(err)

	defer img.Close()

	i, err := png.Decode(img)

	public.CheckErr(err)

	b := i.Bounds()

	rows := b.Max.X
	cols := b.Max.Y

	upLeft := image.Point{0, 0}
	lowRight := image.Point{rows, cols}

	newImage := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	a := astar.NewAStar(rows, cols)
	p2p := astar.NewPointToPoint()

	cyan := color.RGBA{R: 255, A: 0xff}

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			p := i.At(x, y)

			r, _, _, _ := p.RGBA()

			if r != 65535 {
				a.FillTile(astar.Point{y, x}, -1)
				newImage.Set(x, y, color.Black)
			} else {
				newImage.Set(x, y, color.White)
			}
		}
	}

	source := []astar.Point{{Row: 20, Col: 354}}
	target := []astar.Point{{Row: 500, Col: 360}}

	path := a.FindPath(p2p, source, target)

	for path != nil {
		newImage.Set(path.Col, path.Row, cyan)
		path = path.Parent
	}

	newImage.Set(354, 20, cyan)
	newImage.Set(360, 966, cyan)

	f, _ := os.Create("images/result-2.png")
	png.Encode(f, newImage)
}
