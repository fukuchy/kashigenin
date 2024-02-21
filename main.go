package main

import (
	"image/color"
	"log"
    _ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/fukuchy/kashigenin/image"
)

const (
	ScreenWidth  = 1200
	ScreenHeight = 840
)

var img_haikei *ebiten.Image

func init() {
    var err error
    img_haikei, _, err = ebitenutil.NewImageFromFile("haikei.png")
    if err != nil {
        log.Fatal(err)
    }
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 245, 228, 0xff})
    image.Draw(screen, img_haikei, 1.0, 240, 120, 45)
    // op := &ebiten.DrawImageOptions{}
    // op.GeoM.Translate(240, 120)
    // screen.DrawImage(img_haikei, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
