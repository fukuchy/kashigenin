package kashigenin_test

import (
	"log"
	"testing"

	k "github.com/fukuchy/kashigenin/src/kashigenin"
)

func Test_passerCollision(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		game k.Game
		want bool
	}{
		{name: "notCollision,left", game: *setGame(500, "", 5), want: false},
		{name: "notCollision,right", game: *setGame(0, "", 4), want: false},
		{name: "notCollision,midle", game: *setGame(360, "", 6), want: false},
		{name: "notCollision,rightKashige", game: *setGame(360, "right", 1), want: false},
		{name: "notCollision,leftKashige", game: *setGame(360, "left", 2), want: false},
		{name: "collision,left", game: *setGame(360, "", 3), want: true},
		{name: "collision,right", game: *setGame(600, "", 3), want: true},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
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
