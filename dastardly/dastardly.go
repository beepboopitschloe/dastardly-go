package dastardly

import (
	"github.com/nmuth/dastardly-go/screen"
	tb "github.com/nsf/termbox-go"
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
	tb.Clear(tb.ColorDefault, tb.ColorDefault)
	game.viewPanel.Clear(tb.ColorDefault, tb.ColorBlack)
	game.logPanel.Clear(tb.ColorDefault, tb.ColorBlack)
	game.infoPanel.Clear(tb.ColorDefault, tb.ColorBlack)
	game.otherPanel.Clear(tb.ColorDefault, tb.ColorBlack)

	game.viewPanel.DrawBorder(tb.ColorMagenta, tb.ColorBlack)
	game.logPanel.DrawBorder(tb.ColorBlue, tb.ColorBlack)
	game.infoPanel.DrawBorder(tb.ColorCyan, tb.ColorBlack)
	game.otherPanel.DrawBorder(tb.ColorGreen, tb.ColorBlack)

	x, y := game.player.Position()
	game.viewPanel.SetCell(x, y, '@', tb.ColorWhite, tb.ColorBlack)

	game.viewPanel.Blit(40, 0)
	game.logPanel.Blit(40, 30)
	game.infoPanel.Blit(0, 0)
	game.otherPanel.Blit(0, 20)
	tb.Flush()
}

func NewGame() *Game {
	return &Game{
		viewPanel:       screen.NewScreen(80, 30),
		logPanel:        screen.NewScreen(80, 10),
		otherPanel:      screen.NewScreen(40, 20),
		infoPanel:       screen.NewScreen(40, 20),
		smallPopupPanel: screen.NewScreen(3, 80),
		largePopupPanel: screen.NewScreen(80, 30),

		player: &Entity{Ch: '@', Fg: tb.ColorWhite, Bg: tb.ColorDefault},
	}
}

func Start() {
	game := NewGame()

	game.player.SetPosition(10, 10)

	err := tb.Init()
	if err != nil {
		panic(err)
	}
	defer tb.Close()

	tb.HideCursor()

	draw(game)

loop:
	for {
		switch ev := tb.PollEvent(); ev.Type {
		case tb.EventKey:
			switch ev.Key {
			case tb.KeyEsc:
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
