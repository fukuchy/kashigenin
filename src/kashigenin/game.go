package kashigenin

import (
	"image/color"
	_ "image/png"
	"strconv"

	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 960
	ScreenHeight = 840
)
const (
	Status_home = iota
	Status_alive
	Status_dead
)

var Score int

type Game struct {
	Status int
}

func NewGame() (*Game, error) {
	g := &Game{
		Status: 0,
	}
	return g, nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	switch g.Status {
	case 0:
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.Status = 1
		}
	case 1:
		Move()
		if pop_passers() {
			g.Status = 2
		}
	case 2:
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			Score = 0
			Counter = 0
			Yokero_flag = false
			Passers = nil
			g.Status = 1
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 245, 228, 0xff})
	image.Draw(screen, Img_haikei, 1.0, 0, 0, 0)
	switch g.Status {
	case 0:
		image.Draw(screen, Img_title, 1.0, 180, 120, 0)
		image.Draw(screen, Img_setsumei, 1.0, 0, 720, 0)
	case 1:
		Player_Draw(screen)
		PassersDraw(screen)
		if Yokero_flag {
			image.Draw(screen, Img_yokero, 1.0, 180, 180, 0)
		}
		ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(Score))
		image.Draw(screen, Img_setsumei, 1.0, 0, 720, 0)
	case 2:
		image.Draw(screen, Img_gameover, 1.0, 180, 180, 0)
	}
}
