package kashigenin

import (
// "github/fukuchy/kashigenin/src/image"
)

var passer Person = Person{
	body: &Rectangle{
		x:      420,
		y:      360,
		width:  120,
		height: 360,
	},
	kasa: &Rectangle{
		x:      420,
		y:      330,
		width:  240,
		height: 30,
	},
	mae: true, // 仮置き
}

func NewPasser() *Person {
	// body := &Rectangle{
	// 	x:      420,
	// 	y:      360,
	// 	width:  120,
	// 	height: 360,
	// }
	// kasa := &Rectangle{
	// 	x:      420,
	// 	y:      330,
	// 	width:  240,
	// 	height: 30,
	// }
	passer := &Person{
		body: &Rectangle{
			x:      420,
			y:      360,
			width:  120,
			height: 360,
		},
		kasa: &Rectangle{
			x:      420,
			y:      330,
			width:  240,
			height: 30,
		},
		mae: true, // 仮置き
	}
	return passer
}
