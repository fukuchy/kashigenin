package image

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func Draw(screen *ebiten.Image, img *ebiten.Image, dx float64, dy float64, radius float64) {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(radius / 180 * math.Pi)
	op.GeoM.Translate(dx+(float64(w)/2), dy+(float64(h)/2))
	screen.DrawImage(img, op)
}
func Draw_kasa(screen *ebiten.Image, img *ebiten.Image, dx float64, dy float64, radius float64) {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h))
	op.GeoM.Rotate(radius / 180 * math.Pi)
	op.GeoM.Translate(dx+(float64(w)/2), dy+float64(h))
	screen.DrawImage(img, op)
}

func Load(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
