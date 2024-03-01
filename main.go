package main

import (
	"log"

	"github.com/fukuchy/kashigenin/src/kashigenin"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := kashigenin.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("必殺！ カシゲ人")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
