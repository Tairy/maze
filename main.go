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
	img, err := os.Open("./images/source.png")
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

	cyan := color.RGBA{255, 0, 0, 0xff}

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

	source := []astar.Point{{Row: 16, Col: 534}}
	target := []astar.Point{{Row: 1066, Col: 560}}

	path := a.FindPath(p2p, source, target)

	for path != nil {
		newImage.Set(path.Col, path.Row, cyan)
		path = path.Parent
	}

	f, _ := os.Create("images/result.png")
	png.Encode(f, newImage)
}
