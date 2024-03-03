package kashigenin

import (
	"math"
)

// 四角形の構造体
type Rectangle struct {
	x      float64
	y      float64
	width  float64
	height float64
}

// 人と傘の構造体
type Person struct {
	body *Rectangle
	kasa *Rectangle
	num  int
}

// 四角形の中心座標を返す関数
func (rec Rectangle) recmid() (float64, float64) {
	return rec.x + (rec.width / 2), rec.y + (rec.height / 2)
}

// ２ つの四角形の当たり判定
func rechit(rec1 Rectangle, rec2 Rectangle) bool {
	// それぞれの四角形の中心座標を求める
	midx1, midy1 := rec1.recmid()
	midx2, midy2 := rec2.recmid()
	// ２ つの中心座標間の距離を計算する
	distance_x := math.Abs(midx1 - midx2)
	distance_y := math.Abs(midy1 - midy2)
	// ２ つの四角形の横幅、縦幅の半分を足した距離を求める
	halfWidth := (rec1.width + rec2.width) / 2
	halfHeight := (rec1.height + rec2.height) / 2

	// 中心座標間の距離が ２ つの四角形の横幅、縦幅の半分を足した距離より短い場合、その四角形は重なっているので true を返す
	if distance_x < halfWidth && distance_y < halfHeight {
		return true
	}
	return false
}

// 人の当たり判定
func perhit(per1 Person, per2 Person) bool {
	// 人と人、人と傘、傘と傘のどれかが重なっている場合、人同士が接しているとして true を返す
	if rechit(*per1.body, *per2.body) || rechit(*per1.body, *per2.kasa) || rechit(*per1.kasa, *per2.body) || rechit(*per1.kasa, *per2.kasa) {
		return true
	}
	return false
}
