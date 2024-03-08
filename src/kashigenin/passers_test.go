package kashigenin_test

import (
	"log"
	"testing"

	k "github.com/fukuchy/kashigenin/src/kashigenin"
)

func Test_passerHit(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		game k.Game
		want bool
	}{
		{name: "notHit,left", game: *setGame(500, "", 5), want: false},
		{name: "notHit,right", game: *setGame(0, "", 4), want: false},
		{name: "notHit,midle", game: *setGame(360, "", 6), want: false},
		{name: "notHit,rightKashige", game: *setGame(360, "right", 1), want: false},
		{name: "notHit,leftKashige", game: *setGame(360, "left", 2), want: false},
		{name: "hit,left", game: *setGame(360, "", 3), want: true},
		{name: "hit,right", game: *setGame(600, "", 3), want: true},
		{name: "hit,mid", game: *setGame(480, "", 3), want: true},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			// もし想定と結果が違ったらエラー表示
			result := c.game.PasserHit()
			if c.want != result {
				t.Errorf("Results do not match")
			}
		})
	}
}

// テスト用の Game 初期化関数
// Player の x 座標、傘の傾げの向き、通行人の種類を選択する
func setGame(x float64, kashige string, P_index int) *k.Game {

	game, err := k.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	game.Player = *k.NewPlayer_v(x, 360, 120, 360, kashige)
	// P_index で Passers_list の対応した通行人のセットを P_list に格納する
	game.Passers.P_list = k.Passers_list[P_index]
	return game
}
