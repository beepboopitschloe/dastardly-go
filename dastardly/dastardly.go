package dastardly

import (
	"github.com/nmuth/dastardly-go/screen"
	"github.com/nsf/termbox-go"
)

type Game struct {
	viewPanel,
	logPanel,
	infoPanel,
	otherPanel,
	smallPopupPanel,
	largePopupPanel *screen.Screen

	player *Entity
}

const SCREEN_WIDTH = 100
const SCREEN_HEIGHT = 40

func draw(game *Game) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	game.viewPanel.Clear(termbox.ColorDefault, termbox.ColorDefault)
	game.logPanel.Clear(termbox.ColorDefault, termbox.ColorDefault)
	game.infoPanel.Clear(termbox.ColorDefault, termbox.ColorDefault)
	game.otherPanel.Clear(termbox.ColorDefault, termbox.ColorDefault)

	game.viewPanel.DrawBorder()
	game.logPanel.DrawBorder()
	game.infoPanel.DrawBorder()
	game.otherPanel.DrawBorder()

	x, y := game.player.Position()
	game.viewPanel.SetCell(x, y, '@', termbox.ColorBlack, termbox.ColorWhite)

	game.viewPanel.Blit(40, 0)
	game.logPanel.Blit(40, 30)
	game.infoPanel.Blit(0, 0)
	game.otherPanel.Blit(0, 20)
	termbox.Flush()
}

func NewGame() *Game {
	return &Game{
		viewPanel:       screen.NewScreen(80, 30),
		logPanel:        screen.NewScreen(80, 10),
		otherPanel:      screen.NewScreen(40, 20),
		infoPanel:       screen.NewScreen(40, 20),
		smallPopupPanel: screen.NewScreen(3, 80),
		largePopupPanel: screen.NewScreen(80, 30),

		player: &Entity{Ch: '@', Fg: termbox.ColorWhite, Bg: termbox.ColorDefault},
	}
}

func Start() {
	game := NewGame()

	game.player.SetPosition(10, 10)

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.HideCursor()

	draw(game)

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			}

			switch ev.Ch {
			case 'h':
				game.player.Move(-1, 0)
			case 'j':
				game.player.Move(0, 1)
			case 'k':
				game.player.Move(0, -1)
			case 'l':
				game.player.Move(1, 0)
			}
		}

		draw(game)
	}
}
