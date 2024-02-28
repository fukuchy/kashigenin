package image

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

func Draw(screen *ebiten.Image, img *ebiten.Image, scale float64, dx float64, dy float64, radius float64) {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	var sw, sh float64 = float64(w) * scale, float64(h) * scale
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(-float64(sw)/2, -float64(sh)/2)
	op.GeoM.Rotate(radius / 180 * math.Pi)
	op.GeoM.Translate(dx+(float64(sw)/2), dy+(float64(sh)/2))
	screen.DrawImage(img, op)
}
