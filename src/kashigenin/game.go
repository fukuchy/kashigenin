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

var (
	Img_haikei     = image.Load("https://raw.githubusercontent.com/fukuchy/kashigenin/main/material/haikei.png")
	Img_player     = image.Load("https://raw.githubusercontent.com/fukuchy/kashigenin/main/material/player.png")
	Img_passer0    = image.Load("https://raw.githubusercontent.com/fukuchy/kashigenin/main/material/passer0.png")
	Img_passer1    = image.Load("https://raw.githubusercontent.com/fukuchy/kashigenin/main/material/passer1.png")
	Img_passer2    = image.Load("https://raw.githubusercontent.com/fukuchy/kashigenin/main/material/passer2.png")
	Img_passer3    = image.Load("https://raw.githubusercontent.com/fukuchy/kashigenin/main/material/passer3.png")
	Img_kasa       = image.Load("https://raw.githubusercontent.com/fukuchy/kashigenin/main/material/kasa_top.png")
	Img_kasa_mae   = image.Load("https://raw.githubusercontent.com/fukuchy/kashigenin/main/material/kasa_all.png")
	Img_kasa_tsuka = image.Load("https://raw.githubusercontent.com/fukuchy/kashigenin/main/material/kasa_tsuka.png")
	Img_yokero     = image.Load("https://raw.githubusercontent.com/fukuchy/kashigenin/main/material/yokero.png")
	Img_title      = image.Load("https://raw.githubusercontent.com/fukuchy/kashigenin/main/material/title.png")
	Img_gameover   = image.Load("https://raw.githubusercontent.com/fukuchy/kashigenin/main/material/gameover.png")
	Img_setsumei   = image.Load("https://raw.githubusercontent.com/fukuchy/kashigenin/main/material/setsumei.png")
)

// ステータスを設定
const (
	Status_home = iota
	Status_alive
	Status_dead
)

var mplusNormalFont font.Face

// ゲーム全体の構造体
type Game struct {
	Status       int
	Score        int
	ScreenWidth  int
	ScreenHeight int
	Player       Player
	Passers      Passers
}

// フォントの設定と画像のロードを行う関数
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
	// image.Init_img()
}

// ゲームの初期化関数
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

// ステータスに合わせて適切な関数を呼び出す関数
func (g *Game) Update() error {
	switch g.Status {

	case 0:
		// ホーム画面で Space key を押したらゲームを開始する
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.Status = 1
		}
	case 1:
		// ゲームプレイ中はプレイヤーの操作を許可する
		g.Move()
		// 「避けろ！！」表示中に通行人と接触したらゲーム終了する．
		if g.pop_passers() {
			g.Status = 2
		}
	case 2:
		// ゲームオーバー時に Space key を押したら変数を初期化してゲームを開始する
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

// ゲーム全体の描画関数
func (g *Game) Draw(screen *ebiten.Image) {
	image.Draw(screen, Img_haikei, 0, 0, 0)
	// ステータスによって表示する画像を選択する
	switch g.Status {
	case 0:
		// ゲーム開始前にはタイトルと操作説明を表示
		image.Draw(screen, Img_title, 180, 120, 0)
		image.Draw(screen, Img_setsumei, 0, 720, 0)
	case 1:
		// ゲーム開始後はプレイヤーと通行人を表示する
		g.Player_Draw(screen)
		g.PassersDraw(screen)
		// Yokero_flag が true の間「避けろ!!」を表示する
		if g.Passers.Yokero_flag {
			image.Draw(screen, Img_yokero, 180, 180, 0)
		}
		// スコアと説明の表示
		text.Draw(screen, "Score: "+strconv.Itoa(g.Score), mplusNormalFont, 0, 30, color.White)
		image.Draw(screen, Img_setsumei, 0, 720, 0)
	case 2:
		// ゲームオーバー表示と最終スコア表示
		image.Draw(screen, Img_gameover, 250, 180, 0)
		text.Draw(screen, "Score: "+strconv.Itoa(g.Score), mplusNormalFont, 450, 450, color.Black)
	}
}
