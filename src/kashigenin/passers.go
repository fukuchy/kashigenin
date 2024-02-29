package kashigenin

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
)

var Counter int
var Yokero_flag bool

var Passers []Person

func PassersDraw(screen *ebiten.Image) {
	for _, passer := range Passers {
		switch passer.num {
		case 0:
			image.Draw(screen, Img_passer0, 1.0, passer.body.x, passer.body.y, 0)
			image.Draw(screen, Img_kasa_mae, 1.0, passer.kasa.x, passer.kasa.y, 0)
		case 1:
			image.Draw(screen, Img_passer1, 1.0, passer.body.x, passer.body.y, 0)
			image.Draw(screen, Img_kasa_mae, 1.0, passer.kasa.x, passer.kasa.y, 0)
		case 2:
			image.Draw(screen, Img_passer2, 1.0, passer.body.x, passer.body.y, 0)
			image.Draw(screen, Img_kasa_mae, 1.0, passer.kasa.x, passer.kasa.y, 0)
		case 3:
			image.Draw(screen, Img_passer3, 1.0, passer.body.x, passer.body.y, 0)
			image.Draw(screen, Img_kasa_mae, 1.0, passer.kasa.x, passer.kasa.y, 0)
		}

	}
}

func PasserHit() bool {
	for i := len(Passers) - 1; i >= 0; i-- {
		if perhit(Player, Passers[i]) {
			Passers = append(Passers[:i], Passers[i+1:]...)
			return true
		}
	}
	return false
}
func pop_passers() bool {
	Counter++

	if Counter == 200 {
		list_num, err := rand.Int(rand.Reader, big.NewInt(int64(len(Passers_list))))
		fmt.Println(list_num)
		if err != nil {
			panic(err)
		}
		// Passers = append(Passers, Passers_list[3]...)
		Passers = append(Passers, Passers_list[int(list_num.Int64())]...)
	}
	if Counter == 400 {
		Yokero_flag = true
	}
	if Counter >= 400 {

		if PasserHit() {
			return true
		}
	}
	if Counter == 500 {
		Yokero_flag = false
		if !PasserHit() {
			Score += len(Passers)
		}
		Passers = nil
		Counter = 0
	}
	return false
}
