package dastardly

import "github.com/nsf/termbox-go"

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

func drawString(str string, x, y int) {
	for idx, rn := range str {
		termbox.SetCell(x+idx, y, rn, termbox.ColorWhite, termbox.ColorBlack)
	}
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	termbox.SetCell(1, 1, '@', termbox.ColorBlack, termbox.ColorWhite)

	if showMsg {
		drawString("Hello, world!", 1, 4)
	}

	termbox.Flush()
}

func Start() {
	showMsg = false

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.HideCursor()

	draw()

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

		draw()
	}
}
