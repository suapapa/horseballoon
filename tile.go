// Copyright 2020 Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

	tileMap map[int]*ebiten.Image
)

func init() {
	r, err := os.Open("resource/tiles.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(r)
	if err != nil {
		log.Fatal(err)
	}
	tilesImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	tileMap = make(map[int]*ebiten.Image)
}

func tile(t int) *ebiten.Image {
	if _, ok := tileMap[t]; !ok {
		sx := (t % tileXNum) * tileSize
		sy := (t / tileXNum) * tileSize

		newTile := tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)
		tileMap[t] = newTile
	}

	return tileMap[t]
}

func drawTile(screen *ebiten.Image, x, y int, t int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))

	err := screen.DrawImage(tile(t), op)
	chk(err)
}
