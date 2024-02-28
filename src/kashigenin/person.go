package kashigenin

import (
	"math"
)

type Rectangle struct {
	x      float64
	y      float64
	width  float64
	height float64
}

type Person struct {
	body *Rectangle
	kasa *Rectangle
	num  int
}

func (rec Rectangle) recmid() (float64, float64) {
	return rec.x + (rec.width / 2), rec.y + (rec.height / 2)
}

func rechit(rec1 Rectangle, rec2 Rectangle) bool {
	disx1, disy1 := rec1.recmid()
	disx2, disy2 := rec2.recmid()
	distance_x := math.Abs(disx1 - disx2)
	distance_y := math.Abs(disy1 - disy2)
	halfWidth := (rec1.width + rec2.width) / 2
	halfHeight := (rec1.height + rec2.height) / 2

	if distance_x < halfWidth && distance_y < halfHeight {
		return true
	}
	return false
}

func perhit(per1 Person, per2 Person) bool {
	if rechit(*per1.body, *per2.body) || rechit(*per1.body, *per2.kasa) || rechit(*per1.kasa, *per2.body) || rechit(*per1.kasa, *per2.kasa) {
		return true
	}
	return false
}
