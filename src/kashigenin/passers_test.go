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
		{name: "not kashige, left", game: *setGame(500, "", 5), want: false},
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

// Game struct の構造
// type Game struct {
// 	Status       int
// 	Score        int
// 	ScreenWidth  int
// 	ScreenHeight int
// 	Player       Player
// 	Passers      Passers
// }

func setGame(x float64, kashige string, P_index int) *k.Game {

	game, err := k.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	game.Player = *k.NewPlayer_v(x, 360, 120, 360, kashige)
	game.Passers.P_list = k.Passers_list[P_index]
	return game
}
