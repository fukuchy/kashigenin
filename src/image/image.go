package image

import (
    "math"
    "github.com/hajimehoshi/ebiten/v2"
)

func Draw(screen *ebiten.Image, img *ebiten.Image, coefficient float64, x_coodinates float64, y_coodinates float64, angle float64) {
    w := img.Bounds().Dx()
    h := img.Bounds().Dy()
    var sw, sh float64 = float64(w) * coefficient, float64(h) * coefficient
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Scale(coefficient, coefficient)
    op.GeoM.Translate(-float64(sw)/2, -float64(sh)/2)
    op.GeoM.Rotate(angle / 180 * math.Pi)
    op.GeoM.Translate(x_coodinates+(float64(sw)/2), y_coodinates+(float64(sh)/2))
    screen.DrawImage(img, op)
}
