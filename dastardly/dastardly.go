package dastardly

import (
	"github.com/nmuth/dastardly-go/dastardly/screen"
	"github.com/nsf/termbox-go"
)

var showMsg bool

var player Entity

func draw(scr *screen.Screen) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	scr.Clear(termbox.ColorDefault, termbox.ColorRed)

	scr.DrawBorder()

	x, y := player.Position()
	scr.SetCell(x, y, '@', termbox.ColorBlack, termbox.ColorWhite)

	if showMsg {
		scr.DrawString("Hello, world!", 2, 17)
	}

	scr.Blit(10, 10)
	termbox.Flush()
}

func Start() {
	mainScreen := screen.NewScreen(20, 20)

	showMsg = true

	player = Entity{
		x:  0,
		y:  0,
		Ch: '@',
		Fg: termbox.ColorWhite,
		Bg: termbox.ColorDefault,
	}

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

			switch ev.Ch {
			case 'h':
				player.Move(-1, 0)
			case 'j':
				player.Move(0, 1)
			case 'k':
				player.Move(0, -1)
			case 'l':
				player.Move(1, 0)
			}
		}

		draw(mainScreen)
	}
}
