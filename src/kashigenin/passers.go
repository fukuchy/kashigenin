package kashigenin

import (
	"crypto/rand"
	"math/big"

	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
)

// 通行人の構造体
type Passers struct {
	P_list      []Person
	Counter     int
	Yokero_flag bool
}

// 通行人の初期化関数
func NewPassers() *Passers {
	p := &Passers{
		P_list:      []Person{},
		Counter:     0,
		Yokero_flag: false,
	}
	return p
}

// 通行人の描画関数
func (g *Game) PassersDraw(screen *ebiten.Image) {
	for _, passer := range g.Passers.P_list {
		// 通行人の種類ごとに画像を描画
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

// 通行人とプレイヤーの当たり判定を行う関数
func (g *Game) PasserHit() bool {
	for i := len(g.Passers.P_list) - 1; i >= 0; i-- {
		// もし、プレイヤーが通行人と接していたら、true を返す
		if perhit(g.Player.player, g.Passers.P_list[i]) {
			return true
		}
	}
	return false
}

// 通行人の追加から削除までを行う関数
func (g *Game) pop_passers() bool {
	g.Passers.Counter++

	// Passers_list からランダムで通行人を選択してリストに追加
	if g.Passers.Counter == 200 {
		list_num, err := rand.Int(rand.Reader, big.NewInt(int64(len(Passers_list))))
		if err != nil {
			panic(err)
		}
		g.Passers.P_list = append(g.Passers.P_list, Passers_list[int(list_num.Int64())]...)
	}
	// 「避けろ！！」を表示する
	if g.Passers.Counter == 400 {
		g.Passers.Yokero_flag = true
	}
	// 当たり判定を有効化
	if g.Passers.Counter >= 400 {
		// もし通行人とプレイヤーがぶつかっていたら true を返す
		if g.PasserHit() {
			return true
		}
	}
	if g.Passers.Counter == 500 {
		// 当たり判定を無効化
		g.Passers.Yokero_flag = false
		// もし通行人とぶつかっていなかったら、表示されていた人数分 Score を加算
		if !g.PasserHit() {
			g.Score += len(g.Passers.P_list)
		}
		// 通行人をリストから削除
		g.Passers.P_list = nil
		// カウンターを初期化
		g.Passers.Counter = 0
	}
	return false
}
