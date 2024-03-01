package kashigenin

import (
	_ "image/png"

	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
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

func Init_img() {
	Img_haikei = image.Load("material/haikei7.png")
	Img_player = image.Load("material/player2.png")
	Img_passer0 = image.Load("material/passer0.png")
	Img_passer1 = image.Load("material/passer1.png")
	Img_passer2 = image.Load("material/passer2.png")
	Img_passer3 = image.Load("material/passer3.png")
	Img_kasa_mae = image.Load("material/kasa4.png")
	Img_kasa = image.Load("material/kasa5.png")
	Img_kasa_tsuka = image.Load("material/kasa_tsuka.png")
	Img_yokero = image.Load("material/yokero2.png")
	Img_title = image.Load("material/title2.png")
	Img_gameover = image.Load("material/gameover.png")
	Img_setsumei = image.Load("material/setsumei.png")
}
