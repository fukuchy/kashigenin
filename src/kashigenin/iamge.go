package kashigenin

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
