package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
    "github.com/fukuchy/kashigenin/src/kashigenin"
)

func main() {
    game, err := kashigenin.NewGame()
    if err != nil {
        log.Fatal(err)
    }
	ebiten.SetWindowSize(kashigenin.ScreenWidth, kashigenin.ScreenHeight)
	ebiten.SetWindowTitle("必殺！ カシゲ人")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
