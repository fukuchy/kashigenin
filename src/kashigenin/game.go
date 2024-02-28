package kashigenin

import (
	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	_ "image/png"
	"log"
)

const (
	ScreenWidth  = 960
	ScreenHeight = 840
)

var Img_haikei *ebiten.Image
var Img_player *ebiten.Image
var Img_playerDown *ebiten.Image
var Img_kasa *ebiten.Image

func init() {
	var err error
	Img_haikei, _, err = ebitenutil.NewImageFromFile("material/haikei3.png")
	if err != nil {
		log.Fatal(err)
	}
	Img_player, _, err = ebitenutil.NewImageFromFile("material/player.png")
	if err != nil {
		log.Fatal(err)
	}
	Img_playerDown, _, err = ebitenutil.NewImageFromFile("material/playerDown.png")
	if err != nil {
		log.Fatal(err)
	}
	Img_kasa, _, err = ebitenutil.NewImageFromFile("material/kasa.png")
	if err != nil {
		log.Fatal(err)
	}
}

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
