package kashigenin

import (
	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	CountLimit = 5
	LeftLimit  = 120
	RightLimit = 720
)

var Press_flag bool = false
var Up_flag bool = false
var Kashige_flag_hidari = false
var Kashige_flag_migi = false
var Press_count int
var Player Person = *NewPlayer()

func NewPlayer() *Person {
	Player := &Person{
		body: &Rectangle{
			x:      420,
			y:      360,
			width:  120,
			height: 360,
		},
		kasa: &Rectangle{
			x:      420,
			y:      330,
			width:  240,
			height: 30,
		},
		num: 0,
	}
	return Player
}

func Move() {
	if !Press_flag {
		if !Kashige_flag_hidari && !Kashige_flag_migi {
			if ebiten.IsKeyPressed(ebiten.KeyUp) {
				if Player.kasa.y > 210 {
					Player.kasa.y -= 120
				}
				if Player.kasa.y == 210 {
					Up_flag = true
				}
				Press_flag = true
			}
			if ebiten.IsKeyPressed(ebiten.KeyDown) {
				if Player.kasa.y < 330 {
					Player.kasa.y += 120
				}
				if Up_flag {
					Up_flag = false
				}
				Press_flag = true
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			if Player.body.x < RightLimit {
				Player.body.x += 60
				Player.kasa.x += 60
			}
			Press_flag = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			if Player.body.x > LeftLimit {
				Player.body.x -= 60
				Player.kasa.x -= 60
			}
			Press_flag = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyQ) {
			if Player.kasa.y >= 330 && Player.kasa.y < 390 {
				Player.kasa.y += 60
				Player.kasa.x -= 60
				Player.kasa.width -= 60
				Kashige_flag_hidari = true
				Press_flag = true
			} else if Kashige_flag_migi {
				Player.kasa.y -= 60
				Player.kasa.width += 60
				Kashige_flag_migi = false
				Press_flag = true
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyE) {
			if Player.kasa.y >= 330 && Player.kasa.y < 390 {
				Player.kasa.y += 60
				Player.kasa.width -= 60
				Kashige_flag_migi = true
				Press_flag = true
			} else if Kashige_flag_hidari {
				Player.kasa.y -= 60
				Player.kasa.x += 60
				Player.kasa.width += 60
				Kashige_flag_hidari = false
				Press_flag = true
			}
		}
	} else {
		Press_count += 1
	}
	if Press_count == CountLimit {
		Press_count = 0
		Press_flag = false
	}
}

func Player_Draw(screen *ebiten.Image) {
	if Kashige_flag_hidari {
		image.Draw_kasa(screen, Img_kasa, 1.0, Player.body.x, Player.kasa.y-30, -30)
		image.Draw(screen, Img_player, 1.0, Player.body.x, Player.body.y, 0)
	} else if Kashige_flag_migi {
		image.Draw_kasa(screen, Img_kasa, 1.0, Player.body.x-120, Player.kasa.y-30, 30)
		image.Draw(screen, Img_player, 1.0, Player.body.x, Player.body.y, 0)
	} else {
		image.Draw_kasa(screen, Img_kasa, 1.0, Player.kasa.x, Player.kasa.y, 0)
		image.Draw(screen, Img_player, 1.0, Player.body.x, Player.body.y, 0)
	}
}
