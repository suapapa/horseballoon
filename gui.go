package main

import (
	"image/color"
	"log"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Game implements ebiten.Game interface.
type Game struct {
	Str   string
	Start time.Time

	sync.RWMutex
}

var game Game

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	g.Lock()
	defer g.Unlock()
	elapse := time.Now().Sub(g.Start)
	if elapse > 3*time.Second {
		g.Str = ""
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0xff, 0x00, 0b11111111})
	g.RLock()
	defer g.RUnlock()
	ebitenutil.DebugPrint(screen, g.Str)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func gui() {
	// game := &Game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("voice-translator")
	ebiten.SetMaxTPS(10)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
