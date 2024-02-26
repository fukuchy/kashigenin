package kashigenin

import (
    "log"
    "image/color"
    _ "image/png"
    "github.com/fukuchy/kashigenin/src/image"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)
const (
	ScreenWidth  = 960 
	ScreenHeight = 840
)

var Img_haikei *ebiten.Image
var Img_player *ebiten.Image
var Img_playerDown *ebiten.Image

func init () {
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
	if Down_flag == false {
		image.Draw(screen, Img_player, 1.0, Playerx, Playery, 0)
	} else {
		image.Draw(screen, Img_playerDown, 1.0, Playerx, Playery+120, 0)
	}
}



