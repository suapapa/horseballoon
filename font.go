package main

import (
	"image/color"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	penFont font.Face
)

func init() {
	fontBytes, err := ioutil.ReadFile("resource/NanumPen.ttf")
	chk(err)
	tt, err := truetype.Parse(fontBytes)
	chk(err)

	const dpi = 72
	penFont = truetype.NewFace(tt, &truetype.Options{
		Size:    40,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func drawTextBullon(screen *ebiten.Image, msg string) {
	if msg == "" {
		return
	}

	w, h := getSize(msg)
	wCnt := (w + tileSize - 1) / tileSize
	hCnt := (h + tileSize - 1) / tileSize
	bImg, err := ebiten.NewImage(
		wCnt*tileSize+tileSize*2, hCnt*tileSize+tileSize*3,
		ebiten.FilterDefault,
	)
	chk(err)

	// draw Top
	var fx, fy int
	drawTile(bImg, fx, fy, 8)
	fx += tileSize
	for i := 0; i < wCnt; i++ {
		drawTile(bImg, fx, fy, 9)
		fx += tileSize
	}
	drawTile(bImg, fx, fy, 10)
	fx = 0
	fy += 16

	// draw middle
	for j := 0; j < hCnt; j++ {
		drawTile(bImg, fx, fy, 15)
		fx += tileSize
		for i := 0; i < wCnt; i++ {
			drawTile(bImg, fx, fy, 16)
			fx += tileSize
		}
		drawTile(bImg, fx, fy, 17)
		fx = 0
		fy += 16
	}

	// draw bottom
	drawTile(bImg, fx, fy, 22)
	fx += tileSize
	for i := 0; i < wCnt; i++ {
		drawTile(bImg, fx, fy, 23)
		fx += tileSize
	}
	drawTile(bImg, fx, fy, 24)
	fx -= 16
	fy += 16

	// draw tail
	drawTile(bImg, fx, fy, 12)
	drawTile(bImg, fx, fy-16, 24)

	msgW, msgH := getSize(msg)

	bW, bH := bImg.Size()
	msgX := (bW - msgW) / 2
	msgY := (bH-msgH)/2 + 24
	text.Draw(bImg, msg, penFont, msgX, msgY, color.White)

	scrnW, scrnH := screen.Size()
	bX := scrnW - bW - 10
	bY := scrnH - bH - 10
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(bX), float64(bY))
	screen.DrawImage(bImg, op)
}

func getSize(msg string) (w, h int) {
	runes := []rune(msg)
	fx, fy := fixed.I(0), fixed.I(0)

	fy = penFont.Metrics().Height
	for _, r := range runes {
		x, _ := penFont.GlyphAdvance(r)
		fx += x
	}

	return fx.Round(), fy.Round()
}
