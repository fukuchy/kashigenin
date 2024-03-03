package kashigenin

import (
	"image/color"
	_ "image/png"
	"log"
	"os"
	"strconv"

	"github.com/fukuchy/kashigenin/src/image"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	Status_home = iota
	Status_alive
	Status_dead
)

var mplusNormalFont font.Face

type Game struct {
	Status       int
	Score        int
	ScreenWidth  int
	ScreenHeight int
	Player       Player
	Passers      Passers
}

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	Init_img()
}

func NewGame() (*Game, error) {
	g := &Game{
		Status:       0,
		Score:        0,
		ScreenWidth:  960,
		ScreenHeight: 840,
		Player:       *NewPlayer(),
		Passers:      *NewPassers(),
	}
	return g, nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenWidth, g.ScreenHeight
}

func (g *Game) Update() error {
	switch g.Status {
	case 0:
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.Status = 1
		}
	case 1:
		g.Move()
		if g.pop_passers() {
			g.Status = 2
		}
	case 2:
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.Score = 0
			g.Passers.Counter = 0
			g.Passers.Yokero_flag = false
			g.Passers.P_list = nil
			g.Status = 1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 245, 228, 0xff})
	image.Draw(screen, Img_haikei, 0, 0, 0)
	switch g.Status {
	case 0:
		image.Draw(screen, Img_title, 180, 120, 0)
		image.Draw(screen, Img_setsumei, 0, 720, 0)
	case 1:
		g.Player_Draw(screen)
		g.PassersDraw(screen)
		if g.Passers.Yokero_flag {
			image.Draw(screen, Img_yokero, 180, 180, 0)
		}
		text.Draw(screen, "Score: "+strconv.Itoa(g.Score), mplusNormalFont, 0, 30, color.White)
		image.Draw(screen, Img_setsumei, 0, 720, 0)
	case 2:
		image.Draw(screen, Img_gameover, 250, 180, 0)
		text.Draw(screen, "Score: "+strconv.Itoa(g.Score), mplusNormalFont, 450, 450, color.Black)
	}
}
