package kashigenin

import (
	"github.com/hajimehoshi/ebiten/v2"
    "github.com/fukuchy/kashigenin/src/image"
)

const (
	CountLimit   = 5
	LeftLimit    = 120
	RightLimit   = 720
)
var Press_flag bool = false
var Up_flag bool = false
var Press_count int
var Player Person = *NewPlayer()

type Rectangle struct {
    x float64
    y float64
    width float64
    height float64
}
type Person struct {
    body *Rectangle
    kasa *Rectangle
}

func NewPlayer() *Person {
    player_body := &Rectangle{
        x: 420,
        y: 360,
        width: 120,
        height: 360,
    }
    player_kasa := &Rectangle{
        x: 420,
        y: 330,
        width: 240,
        height: 30,
    }
    Player := &Person{
        body: player_body,
        kasa: player_kasa,
    }
    return Player
}



func  Move () {
    if Press_flag == false {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
            if Player.kasa.y > 210 {
                Player.kasa.y -= 120
            }
            if Player.body.y > 360 {
                Player.body.y -= 120
            }
            if Player.kasa.y == 210 {
                Up_flag = true
            }
			Press_flag = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
            if Player.kasa.y < 450 {
                Player.kasa.y += 120
            }
            if Up_flag == true {
                Up_flag = false
            } else if Player.body.y < 480 {
                Player.body.y += 120
            }
			Press_flag = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			if Player.body.x < 720 {
				Player.body.x += 60
                Player.kasa.x += 60
			}
			Press_flag = true

		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			if Player.body.x > 120 {
				Player.body.x -= 60
                Player.kasa.x -= 60
			}
			Press_flag = true

		}
	} else {
		Press_count += 1
	}
	if Press_count == CountLimit {
		Press_count = 0
		Press_flag = false
	}
}

func Player_Draw (screen *ebiten.Image) {
    if Player.body.y == 480 {
        image.Draw(screen, Img_playerDown, 1.0, Player.body.x, Player.body.y, 0)
        image.Draw(screen, Img_kasa, 1.0, Player.kasa.x, Player.kasa.y, 0)
    } else {
        image.Draw(screen, Img_player, 1.0, Player.body.x, Player.body.y, 0)
        image.Draw(screen, Img_kasa, 1.0, Player.kasa.x, Player.kasa.y, 0)
    }
}
