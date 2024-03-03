package kashigenin

import (
	"crypto/rand"
	"math/big"

	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
)

var Counter int
var Yokero_flag bool

func (g *Game) PassersDraw(screen *ebiten.Image) {
	for _, passer := range g.Passers {
		switch passer.num {
		case 0:
			image.Draw(screen, Img_passer0, passer.body.x, passer.body.y, 0)
			image.Draw(screen, Img_kasa_mae, passer.kasa.x, passer.kasa.y, 0)
		case 1:
			image.Draw(screen, Img_passer1, passer.body.x, passer.body.y, 0)
			image.Draw(screen, Img_kasa_mae, passer.kasa.x, passer.kasa.y, 0)
		case 2:
			image.Draw(screen, Img_passer2, passer.body.x, passer.body.y, 0)
			image.Draw(screen, Img_kasa_mae, passer.kasa.x, passer.kasa.y, 0)
		case 3:
			image.Draw(screen, Img_passer3, passer.body.x, passer.body.y, 0)
			image.Draw(screen, Img_kasa_mae, passer.kasa.x, passer.kasa.y, 0)
		}

	}
}

func (g *Game) PasserHit() bool {
	for i := len(g.Passers) - 1; i >= 0; i-- {
		if perhit(g.Player, g.Passers[i]) {
			g.Passers = append(g.Passers[:i], g.Passers[i+1:]...)
			return true
		}
	}
	return false
}

func (g *Game) pop_passers() bool {
	Counter++

	if Counter == 200 {
		list_num, err := rand.Int(rand.Reader, big.NewInt(int64(len(Passers_list))))
		if err != nil {
			panic(err)
		}
		g.Passers = append(g.Passers, Passers_list[int(list_num.Int64())]...)
	}
	if Counter == 400 {
		Yokero_flag = true
	}
	if Counter >= 400 {
		if g.PasserHit() {
			return true
		}
	}
	if Counter == 500 {
		Yokero_flag = false
		if !g.PasserHit() {
			g.Score += len(g.Passers)
		}
		g.Passers = nil
		Counter = 0
	}
	return false
}
