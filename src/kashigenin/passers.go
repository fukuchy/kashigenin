package kashigenin

import (
	"fmt"

	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
)

var counter int

var passers_list = []Person{
	{
		body: &Rectangle{
			x:      660,
			y:      360,
			width:  120,
			height: 360,
		},
		kasa: &Rectangle{
			x:      660,
			y:      330,
			width:  240,
			height: 30,
		},
		num: 1,
	},
	{
		body: &Rectangle{
			x:      240,
			y:      360,
			width:  120,
			height: 360,
		},
		kasa: &Rectangle{
			x:      240,
			y:      330,
			width:  240,
			height: 30,
		},
		num: 2,
	},
}

var passers []Person

func PassersDraw(screen *ebiten.Image) {
	for _, passer := range passers {
		image.Draw(screen, Img_player, 1.0, passer.body.x, passer.body.y, 0)
		image.Draw(screen, Img_kasa, 1.0, passer.kasa.x, passer.kasa.y, 0)
	}
}

func PasserHit() {
	for i := len(passers) - 1; i >= 0; i-- {
		if perhit(Player, passers[i]) {
			fmt.Println("HIT!!")
			passers = append(passers[:i], passers[i+1:]...)
		}
	}
}
func pop_passers() {
	counter++
	if counter%200 == 0 {
		passers = append(passers, passers_list[0])
	}
	if counter%200 == 0 {
		passers = append(passers, passers_list[1])
	}
	if counter == 400 {
		fmt.Println("NOW!!!")
	}
	if counter >= 400 {

		PasserHit()
	}
	if counter == 500 {
		passers = nil
		counter = 0
	}
}
