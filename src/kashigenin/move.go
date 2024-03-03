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

func (g *Game) Move() {
	if !press_flag {
		if !kashige_flag_left && !kashige_flag_right {
			if ebiten.IsKeyPressed(ebiten.KeyUp) {
				if g.Player.kasa.y > 210 {
					g.Player.kasa.y -= 120
				}
				if g.Player.kasa.y == 210 {
					up_flag = true
				}
				press_flag = true
			}
			if ebiten.IsKeyPressed(ebiten.KeyDown) {
				if g.Player.kasa.y < 330 {
					g.Player.kasa.y += 120
				}
				if up_flag {
					up_flag = false
				}
				press_flag = true
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			if g.Player.body.x < RightLimit {
				g.Player.body.x += 60
				g.Player.kasa.x += 60
			}
			press_flag = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			if g.Player.body.x > LeftLimit {
				g.Player.body.x -= 60
				g.Player.kasa.x -= 60
			}
			press_flag = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyQ) {
			switch {
			case g.Player.kasa.y >= 330 && g.Player.kasa.y < 390:
				g.Player.kasa.y += 60
				g.Player.kasa.x -= 60
				g.Player.kasa.width -= 60
				kashige_flag_left = true
				press_flag = true
			case kashige_flag_right:
				g.Player.kasa.y -= 60
				g.Player.kasa.width += 60
				kashige_flag_right = false
				press_flag = true
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyE) {
			switch {
			case g.Player.kasa.y >= 330 && g.Player.kasa.y < 390:
				g.Player.kasa.y += 60
				g.Player.kasa.width -= 60
				kashige_flag_right = true
				press_flag = true
			case kashige_flag_left:
				g.Player.kasa.y -= 60
				g.Player.kasa.x += 60
				g.Player.kasa.width += 60
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

func (g *Game) Player_Draw(screen *ebiten.Image) {
	switch {
	case kashige_flag_left:
		image.Draw_kasa(screen, Img_kasa_tsuka, g.Player.body.x, g.Player.kasa.y-30, -30)
		image.Draw(screen, Img_player, g.Player.body.x, g.Player.body.y+15, 0)
		image.Draw_kasa(screen, Img_kasa, g.Player.body.x, g.Player.kasa.y-30, -30)
	case kashige_flag_right:
		image.Draw_kasa(screen, Img_kasa_tsuka, g.Player.body.x-120, g.Player.kasa.y-30, 30)
		image.Draw(screen, Img_player, g.Player.body.x, g.Player.body.y+15, 0)
		image.Draw_kasa(screen, Img_kasa, g.Player.body.x-120, g.Player.kasa.y-30, 30)
	default:
		image.Draw_kasa(screen, Img_kasa_tsuka, g.Player.kasa.x, g.Player.kasa.y, 0)
		image.Draw(screen, Img_player, g.Player.body.x, g.Player.body.y+15, 0)
		image.Draw_kasa(screen, Img_kasa, g.Player.kasa.x, g.Player.kasa.y, 0)
	}
}
