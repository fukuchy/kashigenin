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

var (
	press_flag,
	up_flag,
	kashige_flag_left,
	kashige_flag_right bool
	press_count int
	Player      Person = *NewPlayer()
)

func NewPlayer() *Person {
	player := &Person{
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
	return player
}

func Move() {
	if !press_flag {
		if !kashige_flag_left && !kashige_flag_right {
			if ebiten.IsKeyPressed(ebiten.KeyUp) {
				if Player.kasa.y > 210 {
					Player.kasa.y -= 120
				}
				if Player.kasa.y == 210 {
					up_flag = true
				}
				press_flag = true
			}
			if ebiten.IsKeyPressed(ebiten.KeyDown) {
				if Player.kasa.y < 330 {
					Player.kasa.y += 120
				}
				if up_flag {
					up_flag = false
				}
				press_flag = true
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			if Player.body.x < RightLimit {
				Player.body.x += 60
				Player.kasa.x += 60
			}
			press_flag = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			if Player.body.x > LeftLimit {
				Player.body.x -= 60
				Player.kasa.x -= 60
			}
			press_flag = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyQ) {
			switch {
			case Player.kasa.y >= 330 && Player.kasa.y < 390:
				Player.kasa.y += 60
				Player.kasa.x -= 60
				Player.kasa.width -= 60
				kashige_flag_left = true
				press_flag = true
			case kashige_flag_right:
				Player.kasa.y -= 60
				Player.kasa.width += 60
				kashige_flag_right = false
				press_flag = true
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyE) {
			switch {
			case Player.kasa.y >= 330 && Player.kasa.y < 390:
				Player.kasa.y += 60
				Player.kasa.width -= 60
				kashige_flag_right = true
				press_flag = true
			case kashige_flag_left:
				Player.kasa.y -= 60
				Player.kasa.x += 60
				Player.kasa.width += 60
				kashige_flag_left = false
				press_flag = true
			}
		}
	} else {
		press_count += 1
	}
	if press_count == CountLimit {
		press_count = 0
		press_flag = false
	}
}

func Player_Draw(screen *ebiten.Image) {
	switch {
	case kashige_flag_left:
		image.Draw_kasa(screen, Img_kasa_tsuka, Player.body.x, Player.kasa.y-30, -30)
		image.Draw(screen, Img_player, Player.body.x, Player.body.y+15, 0)
		image.Draw_kasa(screen, Img_kasa, Player.body.x, Player.kasa.y-30, -30)
	case kashige_flag_right:
		image.Draw_kasa(screen, Img_kasa_tsuka, Player.body.x-120, Player.kasa.y-30, 30)
		image.Draw(screen, Img_player, Player.body.x, Player.body.y+15, 0)
		image.Draw_kasa(screen, Img_kasa, Player.body.x-120, Player.kasa.y-30, 30)
	default:
		image.Draw_kasa(screen, Img_kasa_tsuka, Player.kasa.x, Player.kasa.y, 0)
		image.Draw(screen, Img_player, Player.body.x, Player.body.y+15, 0)
		image.Draw_kasa(screen, Img_kasa, Player.kasa.x, Player.kasa.y, 0)
	}
}
