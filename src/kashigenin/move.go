package kashigenin

import (
	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// ボタンのクールタイム
	CountLimit = 5
	// Player が移動範囲
	LeftLimit  = 120
	RightLimit = 720
)

// プレイヤーの構造体
type Player struct {
	player             Person
	press_flag         bool
	up_flag            bool
	kashige_flag_left  bool
	kashige_flag_right bool
	press_count        int
}

// プレイヤーの初期化関数
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

// テストコード用の座標指定プレイヤー初期化関数
func NewPlayer_v(x, y, width, height float64, kashige string) *Player {
	var player *Player
	// kashige 変数によって傘を傾げた状態で初期化
	switch kashige {
	case "right":
		player = &Player{
			player: Person{
				body: &Rectangle{x: x, y: y, width: width, height: height},
				kasa: &Rectangle{x: x, y: y + 30, width: 240 - 60, height: 30},
			},
		}
	case "left":
		player = &Player{
			player: Person{
				body: &Rectangle{x: x, y: y, width: width, height: height},
				kasa: &Rectangle{x: x - 60, y: y + 30, width: 240 - 60, height: 30},
			},
		}
	default:
		player = &Player{
			player: Person{
				body: &Rectangle{x: x, y: y, width: width, height: height},
				kasa: &Rectangle{x: x, y: y - 30, width: 240, height: 30},
			},
		}
	}
	return player
}

// プレイヤーの移動関数
func (g *Game) Move() {
	// press_flag が false のときのみ入力を受け付ける
	if !g.Player.press_flag {
		// 左右の移動
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
		// 傘の傾げ
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
	// ボタンの入力にクールタイムを設けるために、一度 press_flag が true になったら 5 カウント経ってから press_flag を false にする
	if g.Player.press_count == CountLimit {
		g.Player.press_count = 0
		g.Player.press_flag = false
	}
}

// プレイヤーの描画関数
func (g *Game) Player_Draw(screen *ebiten.Image) {
	// 傘の傾げによって描画する位置や角度を変更する
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
