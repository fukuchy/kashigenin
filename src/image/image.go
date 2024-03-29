package image

import (
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// 画像の描画関数
func Draw(screen *ebiten.Image, img *ebiten.Image, dx float64, dy float64, radius float64) {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(radius / 180 * math.Pi)
	op.GeoM.Translate(dx+(float64(w)/2), dy+(float64(h)/2))
	screen.DrawImage(img, op)
}

// 傘の描画関数
func Draw_kasa(screen *ebiten.Image, img *ebiten.Image, dx float64, dy float64, radius float64) {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h))
	op.GeoM.Rotate(radius / 180 * math.Pi)
	op.GeoM.Translate(dx+(float64(w)/2), dy+float64(h))
	screen.DrawImage(img, op)
}

// 画像のロード関数
func Load(path string) *ebiten.Image {
	img, err := ebitenutil.NewImageFromURL(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
