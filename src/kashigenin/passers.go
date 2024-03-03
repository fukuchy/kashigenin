package kashigenin

import (
	"crypto/rand"
	"math/big"

	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
)

type Passers struct {
	P_list      []Person
	Counter     int
	Yokero_flag bool
}

func NewPassers() *Passers {
	p := &Passers{
		P_list:      []Person{},
		Counter:     0,
		Yokero_flag: false,
	}
	return p
}

func (g *Game) PassersDraw(screen *ebiten.Image) {
	for _, passer := range g.Passers.P_list {
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
	for i := len(g.Passers.P_list) - 1; i >= 0; i-- {
		if perhit(g.Player, g.Passers.P_list[i]) {
			g.Passers.P_list = append(g.Passers.P_list[:i], g.Passers.P_list[i+1:]...)
			return true
		}
	}
	return false
}

func (g *Game) pop_passers() bool {
	g.Passers.Counter++

	if g.Passers.Counter == 200 {
		list_num, err := rand.Int(rand.Reader, big.NewInt(int64(len(Passers_list))))
		if err != nil {
			panic(err)
		}
		g.Passers.P_list = append(g.Passers.P_list, Passers_list[int(list_num.Int64())]...)
	}
	if g.Passers.Counter == 400 {
		g.Passers.Yokero_flag = true
	}
	if g.Passers.Counter >= 400 {
		if g.PasserHit() {
			return true
		}
	}
	if g.Passers.Counter == 500 {
		g.Passers.Yokero_flag = false
		if !g.PasserHit() {
			g.Score += len(g.Passers.P_list)
		}
		g.Passers.P_list = nil
		g.Passers.Counter = 0
	}
	return false
}
