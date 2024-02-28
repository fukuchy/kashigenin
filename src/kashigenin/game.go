package kashigenin

import (
	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	_ "image/png"
)

const (
	ScreenWidth  = 960
	ScreenHeight = 840
)

type Game struct{}

func NewGame() (*Game, error) {
	g := &Game{}
	return g, nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	Move()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 245, 228, 0xff})
	image.Draw(screen, Img_haikei, 1.0, 120, 120, 0)
	Player_Draw(screen)
}
