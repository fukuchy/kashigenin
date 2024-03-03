package kashigenin

import (
	_ "image/png"

	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
)

// ゲーム内で使用する画像の宣言とロード
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
	Img_haikei = image.Load("material/haikei.png")
	Img_player = image.Load("material/player.png")
	Img_passer0 = image.Load("material/passer0.png")
	Img_passer1 = image.Load("material/passer1.png")
	Img_passer2 = image.Load("material/passer2.png")
	Img_passer3 = image.Load("material/passer3.png")
	Img_kasa_mae = image.Load("material/kasa_all.png")
	Img_kasa = image.Load("material/kasa_top.png")
	Img_kasa_tsuka = image.Load("material/kasa_tsuka.png")
	Img_yokero = image.Load("material/yokero.png")
	Img_title = image.Load("material/title.png")
	Img_gameover = image.Load("material/gameover.png")
	Img_setsumei = image.Load("material/setsumei.png")
}
