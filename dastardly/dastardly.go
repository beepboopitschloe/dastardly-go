package dastardly

import (
	"github.com/nmuth/dastardly-go/dastardly/screen"
	"github.com/nsf/termbox-go"
)

var showMsg bool

func draw(scr *screen.Screen) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	scr.Clear(termbox.ColorDefault, termbox.ColorRed)

	scr.DrawBorder()

	scr.SetCell(2, 2, '@', termbox.ColorBlack, termbox.ColorWhite)

	if showMsg {
		scr.DrawString("Hello, world!", 2, 17)
	}

	scr.Blit(10, 10)
	termbox.Flush()
}

func Start() {
	mainScreen := screen.NewScreen(20, 20)

	showMsg = true

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.HideCursor()

	draw(mainScreen)

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeySpace:
				showMsg = !showMsg
			}
		}

		draw(mainScreen)
	}
}
