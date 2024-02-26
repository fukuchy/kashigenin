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
var press_flag bool = false
var Down_flag bool = false
var press_count int

type Input struct {}

func (i *Input) Update () {
    if press_flag == false {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			press_flag = true
			Down_flag = false
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			press_flag = true
			Down_flag = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			if Playerx < 720 {
				Playerx += 60
			}
			press_flag = true
			Down_flag = false

		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			if Playerx > 120 {
				Playerx -= 60
			}
			press_flag = true
			Down_flag = false

		}
	} else {
		press_count += 1
	}
	if press_count == CountLimit {
		press_count = 0
		press_flag = false
	}
}

