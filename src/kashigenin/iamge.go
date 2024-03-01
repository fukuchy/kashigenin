package kashigenin

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	Img_haikei,
	Img_player,
	Img_passer0,
	Img_passer1,
	Img_passer2,
	Img_passer3,
	Img_kasa,
	Img_kasa_mae,
	Img_kasa_tsuka,
	Img_yokero,
	Img_title,
	Img_gameover,
	Img_setsumei *ebiten.Image
)

func init_img() {
	var err error
	Img_haikei, _, err = ebitenutil.NewImageFromFile("material/haikei7.png")
	if err != nil {
		log.Fatal(err)
	}
	Img_player, _, err = ebitenutil.NewImageFromFile("material/player2.png")
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
	Img_passer2, _, err = ebitenutil.NewImageFromFile("material/passer2.png")
	if err != nil {
		log.Fatal(err)
	}
	Img_passer3, _, err = ebitenutil.NewImageFromFile("material/passer3.png")
	if err != nil {
		log.Fatal(err)
	}
	Img_kasa_mae, _, err = ebitenutil.NewImageFromFile("material/kasa4.png")
	if err != nil {
		log.Fatal(err)
	}
	Img_kasa, _, err = ebitenutil.NewImageFromFile("material/kasa5.png")
	if err != nil {
		log.Fatal(err)
	}
	Img_kasa_tsuka, _, err = ebitenutil.NewImageFromFile("material/kasa_tsuka.png")
	if err != nil {
		log.Fatal(err)
	}
	Img_yokero, _, err = ebitenutil.NewImageFromFile("material/yokero2.png")
	if err != nil {
		log.Fatal(err)
	}
	Img_title, _, err = ebitenutil.NewImageFromFile("material/title2.png")
	if err != nil {
		log.Fatal(err)
	}
	Img_gameover, _, err = ebitenutil.NewImageFromFile("material/gameover.png")
	if err != nil {
		log.Fatal(err)
	}
	Img_setsumei, _, err = ebitenutil.NewImageFromFile("material/setsumei.png")
	if err != nil {
		log.Fatal(err)
	}
}
