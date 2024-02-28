package kashigenin

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Img_haikei *ebiten.Image
var Img_player *ebiten.Image
var Img_passer0 *ebiten.Image
var Img_passer1 *ebiten.Image
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
	Img_passer0, _, err = ebitenutil.NewImageFromFile("material/passer0.png")
	if err != nil {
		log.Fatal(err)
	}
	Img_passer1, _, err = ebitenutil.NewImageFromFile("material/passer1.png")
	if err != nil {
		log.Fatal(err)
	}
	Img_kasa, _, err = ebitenutil.NewImageFromFile("material/kasa.png")
	if err != nil {
		log.Fatal(err)
	}
}
