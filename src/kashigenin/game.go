package kashigenin

import (
    "log"
    "image/color"
    _ "image/png"
    // "github.com/fukuchy/kashigenin/src/image"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)
const (
	ScreenWidth  = 1200
	ScreenHeight = 840
)
var img_haikei *ebiten.Image
var img_player *ebiten.Image
var img_playerDown *ebiten.Image
func init () {
    var err error
    img_haikei, _, err = ebitenutil.NewImageFromFile("github.com/fukuchy/kashigenin/material/haikei3.png")
    if err != nil {
        log.Fatal(err)
    }
    img_player, _, err = ebitenutil.NewImageFromFile("github.com/fukuchy/kashigenin/material/player.png")
    if err != nil {
        log.Fatal(err)
    }
    img_haikei, _, err = ebitenutil.NewImageFromFile("github.com/fukuchy/kashigenin/material/playerDown.png")
    if err != nil {
        log.Fatal(err)
    }

}
type Game struct{}

func NewGame() (*Game, error) {
    g := &Game{}
    return g, nil
}

func (g *Game) Update() error {
    return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 245, 228, 0xff})
	// image.Draw(screen, img_haikei, 1.0, 120, 120, 0)
	// if move.Down_flag == false {
	// 	image.Draw(screen, img_player, 1.0, move.Playerx, move.Playery, 0)
	// } else {
	// 	image.Draw(screen, img_playerDown, 1.0, move.Playerx, move.Playery+120, 0)

	// }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

