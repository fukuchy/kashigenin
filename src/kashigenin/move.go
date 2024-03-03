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

// var (
// 	g.Player.press_flag,
// 	up_flag,
// 	kashige_flag_left,
// 	kashige_flag_right bool
// 	press_count int
// )

type Player struct {
	player             Person
	press_flag         bool
	up_flag            bool
	kashige_flag_left  bool
	kashige_flag_right bool
	press_count        int
}

func NewPlayer() *Player {
	player := &Player{
		player: Person{
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
		},
		press_flag:         false,
		up_flag:            false,
		kashige_flag_left:  false,
		kashige_flag_right: false,
		press_count:        0,
	}
	return player
}

func (g *Game) Move() {
	if !g.Player.press_flag {
		if !g.Player.kashige_flag_left && !g.Player.kashige_flag_right {
			if ebiten.IsKeyPressed(ebiten.KeyUp) {
				if g.Player.player.kasa.y > 210 {
					g.Player.player.kasa.y -= 120
				}
				if g.Player.player.kasa.y == 210 {
					g.Player.up_flag = true
				}
				g.Player.press_flag = true
			}
			if ebiten.IsKeyPressed(ebiten.KeyDown) {
				if g.Player.player.kasa.y < 330 {
					g.Player.player.kasa.y += 120
				}
				if g.Player.up_flag {
					g.Player.up_flag = false
				}
				g.Player.press_flag = true
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			if g.Player.player.body.x < RightLimit {
				g.Player.player.body.x += 60
				g.Player.player.kasa.x += 60
			}
			g.Player.press_flag = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			if g.Player.player.body.x > LeftLimit {
				g.Player.player.body.x -= 60
				g.Player.player.kasa.x -= 60
			}
			g.Player.press_flag = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyQ) {
			switch {
			case g.Player.player.kasa.y >= 330 && g.Player.player.kasa.y < 390:
				g.Player.player.kasa.y += 60
				g.Player.player.kasa.x -= 60
				g.Player.player.kasa.width -= 60
				g.Player.kashige_flag_left = true
				g.Player.press_flag = true
			case g.Player.kashige_flag_right:
				g.Player.player.kasa.y -= 60
				g.Player.player.kasa.width += 60
				g.Player.kashige_flag_right = false
				g.Player.press_flag = true
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyE) {
			switch {
			case g.Player.player.kasa.y >= 330 && g.Player.player.kasa.y < 390:
				g.Player.player.kasa.y += 60
				g.Player.player.kasa.width -= 60
				g.Player.kashige_flag_right = true
				g.Player.press_flag = true
			case g.Player.kashige_flag_left:
				g.Player.player.kasa.y -= 60
				g.Player.player.kasa.x += 60
				g.Player.player.kasa.width += 60
				g.Player.kashige_flag_left = false
				g.Player.press_flag = true
			}
		}
	} else {
		g.Player.press_count += 1
	}
	if g.Player.press_count == CountLimit {
		g.Player.press_count = 0
		g.Player.press_flag = false
	}
}

func (g *Game) Player_Draw(screen *ebiten.Image) {
	switch {
	case g.Player.kashige_flag_left:
		image.Draw_kasa(screen, Img_kasa_tsuka, g.Player.player.body.x, g.Player.player.kasa.y-30, -30)
		image.Draw(screen, Img_player, g.Player.player.body.x, g.Player.player.body.y+15, 0)
		image.Draw_kasa(screen, Img_kasa, g.Player.player.body.x, g.Player.player.kasa.y-30, -30)
	case g.Player.kashige_flag_right:
		image.Draw_kasa(screen, Img_kasa_tsuka, g.Player.player.body.x-120, g.Player.player.kasa.y-30, 30)
		image.Draw(screen, Img_player, g.Player.player.body.x, g.Player.player.body.y+15, 0)
		image.Draw_kasa(screen, Img_kasa, g.Player.player.body.x-120, g.Player.player.kasa.y-30, 30)
	default:
		image.Draw_kasa(screen, Img_kasa_tsuka, g.Player.player.kasa.x, g.Player.player.kasa.y, 0)
		image.Draw(screen, Img_player, g.Player.player.body.x, g.Player.player.body.y+15, 0)
		image.Draw_kasa(screen, Img_kasa, g.Player.player.kasa.x, g.Player.player.kasa.y, 0)
	}
}
