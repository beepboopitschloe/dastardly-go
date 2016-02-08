package dastardly

import (
	"github.com/nmuth/dastardly-go/dastardly/screen"
	tb "github.com/nsf/termbox-go"
)

type Drawable interface {
	Draw(*screen.Screen)
}

type Entity struct {
	x, y   int
	Ch     rune
	Fg, Bg tb.Attribute
}

func (this *Entity) Position() (int, int) {
	return this.x, this.y
}

func (this *Entity) Move(dx, dy int) (int, int) {
	this.x += dx
	this.y += dy

	return this.x, this.y
}

func (this *Entity) Draw(scr *screen.Screen) {
	x, y := this.Position()

	scr.SetCell(x, y, this.Ch, this.Fg, this.Bg)
}
