package main

import (
	"image"
	"log"
	"os"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

const (
	tileSize = 16
	tileXNum = 7
)

var (
	tilesImage *ebiten.Image
)

func tile(t int) *ebiten.Image {
	if tilesImage == nil {
		r, err := os.Open("resource/tiles.png")
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(r)
		if err != nil {
			log.Fatal(err)
		}
		tilesImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	}

	sx := (t % tileXNum) * tileSize
	sy := (t / tileXNum) * tileSize

	return tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)
}

func drawTile(screen *ebiten.Image, x, y int, t int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))

	err := screen.DrawImage(tile(t), op)
	chk(err)
}
