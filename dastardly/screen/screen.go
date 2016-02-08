package screen

import tb "github.com/nsf/termbox-go"

type Screen struct {
	buffer []tb.Cell
	width  int
	height int
}

func NewScreen(width, height int) *Screen {
	return &Screen{
		buffer: make([]tb.Cell, width*height, width*height),
		width:  width,
		height: height,
	}
}

func (scr *Screen) Size() (int, int) {
	return scr.width, scr.height
}

func (scr *Screen) CellAt(x, y int) tb.Cell {
	return scr.buffer[(y*scr.width)+x]
}

func (scr *Screen) Blit(x, y int) {
	termWidth, _ := tb.Size()
	startOffset := (y * termWidth) + x
	cells := tb.CellBuffer()

	for i := 0; i < scr.height; i++ {
		tbRangeStart := startOffset + (i * termWidth)
		tbRangeEnd := tbRangeStart + scr.width
		scrRangeStart := i * scr.width
		scrRangeEnd := scrRangeStart + scr.width

		copy(cells[tbRangeStart:tbRangeEnd], scr.buffer[scrRangeStart:scrRangeEnd])
	}
}

func (scr *Screen) Clear(fg, bg tb.Attribute) {
	for i := range scr.buffer {
		c := &scr.buffer[i]

		c.Ch = ' '
		c.Fg = fg
		c.Bg = bg
	}
}

func (scr *Screen) SetCell(x, y int, rn rune, fg, bg tb.Attribute) {
	idx := (y * scr.width) + x

	scr.buffer[idx] = tb.Cell{Ch: rn, Fg: fg, Bg: bg}
}

func (scr *Screen) DrawString(str string, x, y int) {
	for idx, rn := range str {
		scr.SetCell(x+idx, y, rn, tb.ColorWhite, tb.ColorDefault)
	}
}

func (scr *Screen) DrawPanel(x, y, w, h int) {
	fg := tb.ColorWhite
	bg := tb.ColorDefault

	// top border
	for i := 0; i < w; i++ {
		scr.SetCell(x+i, y, '#', fg, bg)
	}

	// bottom border
	for i := 0; i < w; i++ {
		scr.SetCell(x+i, y+h, '#', fg, bg)
	}

	// left border
	for i := 0; i < h; i++ {
		scr.SetCell(x, y+i, '#', fg, bg)
	}

	// right border
	for i := 0; i < h; i++ {
		scr.SetCell(x+w, y+i, '#', fg, bg)
	}

	// bottom right corner
	scr.SetCell(x+w, y+h, '#', fg, bg)
}
