package kashigenin

import (
	"fmt"

	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
)

var counter int

var passers []Person

func PassersDraw(screen *ebiten.Image) {
	for _, passer := range passers {
		switch passer.num {
		case 0:
			image.Draw(screen, Img_passer0, 1.0, passer.body.x, passer.body.y, 0)
		case 1:
			image.Draw(screen, Img_passer1, 1.0, passer.body.x, passer.body.y, 0)
		case 2:
			image.Draw(screen, Img_passer1, 1.0, passer.body.x, passer.body.y, 0)
		case 3:
			image.Draw(screen, Img_passer1, 1.0, passer.body.x, passer.body.y, 0)
		}

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
		passers = append(passers, Passers_list[1]...)
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
