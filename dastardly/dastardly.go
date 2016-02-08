package dastardly

import (
	"github.com/nmuth/dastardly-go/dastardly/screen"
	"github.com/nsf/termbox-go"
)

// import "github.com/nmuth/dastardly-go/dastardly/player"

/*
func Start() {
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorGreen,
		Ch: '.',
	})

	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))

	player := player.NewPlayer(level)
	level.AddEntity(player)

	game := tl.NewGame()
	game.Screen().SetLevel(level)
	game.Screen().SetFps(30)
	game.Start()
}
*/

var showMsg bool

func draw(scr *screen.Screen) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	scr.Clear(termbox.ColorDefault, termbox.ColorRed)

	scr.SetCell(19, 19, '@', termbox.ColorBlack, termbox.ColorWhite)

	// scr.DrawPanel(0, 0, 10, 10)

	if showMsg {
		scr.DrawString("Hello, world!", 0, 19)
	}

	scr.Blit(10, 10)
	termbox.Flush()
}

func Start() {
	mainScreen := screen.NewScreen(20, 20)

	showMsg = false

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
