package kashigenin

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	CountLimit   = 5
	LeftLimit    = 120
	RightLimit   = 720
)
var Playerx float64 = 420
var Playery float64 = 360
var Press_flag bool = false
var Down_flag bool = false
var Press_count int


func  Move () {
    if Press_flag == false {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			Press_flag = true
			Down_flag = false
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			Press_flag = true
			Down_flag = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			if Playerx < 720 {
				Playerx += 60
			}
			Press_flag = true
			Down_flag = false

		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			if Playerx > 120 {
				Playerx -= 60
			}
			Press_flag = true
			Down_flag = false

		}
	} else {
		Press_count += 1
	}
	if Press_count == CountLimit {
		Press_count = 0
		Press_flag = false
	}
}

