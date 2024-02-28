package kashigenin

import (
	"image/color"
	_ "image/png"

	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
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
	// pop_passers()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 245, 228, 0xff})
	image.Draw(screen, Img_haikei, 1.0, 120, 120, 0)
	Player_Draw(screen)
	PassersDraw(screen)
}
