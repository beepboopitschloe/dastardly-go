package game

import (
	"github.com/nmuth/dastardly-go/entity"
	"github.com/nmuth/dastardly-go/screen"
	tb "github.com/nsf/termbox-go"
)

type Panel struct {
	X, Y, Width, Height int
	Screen              *screen.Screen
	BorderColor         tb.Attribute
}

func NewPanel(x, y, w, h int, borderColor tb.Attribute) *Panel {
	return &Panel{
		X:           x,
		Y:           y,
		Width:       w,
		Height:      h,
		Screen:      screen.NewScreen(w, h),
		BorderColor: borderColor,
	}
}

func (this *Panel) Blit() {
	this.Screen.Blit(this.X, this.Y)
}

func (this *Panel) DrawBorder() {
	this.Screen.DrawBorder(this.BorderColor, tb.ColorBlack)
}

func (this *Panel) Clear() {
	this.Screen.Clear(tb.ColorDefault, tb.ColorBlack)
}

type PanelProcessor func(panel Panel)

type Game struct {
	ViewPanel,
	LogPanel,
	InfoPanel,
	OtherPanel,
	SmallPopupPanel,
	LargePopupPanel *Panel

	Player *entity.Entity
}

func NewGame() *Game {
	game := &Game{
		OtherPanel: NewPanel(0, 0, 40, 20, tb.ColorGreen),
		InfoPanel:  NewPanel(0, 20, 40, 20, tb.ColorCyan),
		ViewPanel:  NewPanel(40, 0, 80, 30, tb.ColorMagenta),
		LogPanel:   NewPanel(40, 30, 80, 10, tb.ColorBlue),

		SmallPopupPanel: NewPanel(0, 0, 3, 80, tb.ColorWhite),
		LargePopupPanel: NewPanel(0, 0, 80, 30, tb.ColorWhite),

		Player: &entity.Entity{
			Ch: '@',
			Fg: tb.ColorWhite,
			Bg: tb.ColorDefault,
		},
	}

	game.Player.SetPosition(10, 10)

	return game
}

func (this *Game) Draw() {
	tb.Clear(tb.ColorDefault, tb.ColorDefault)

	this.ViewPanel.Clear()
	this.LogPanel.Clear()
	this.InfoPanel.Clear()
	this.OtherPanel.Clear()

	this.ViewPanel.DrawBorder()
	this.LogPanel.DrawBorder()
	this.InfoPanel.DrawBorder()
	this.OtherPanel.DrawBorder()

	x, y := this.Player.Position()
	this.ViewPanel.Screen.SetCell(x, y, '@', tb.ColorWhite, tb.ColorBlack)

	this.ViewPanel.Blit()
	this.LogPanel.Blit()
	this.InfoPanel.Blit()
	this.OtherPanel.Blit()

	tb.Flush()
}

func (this *Game) Run() {
	err := tb.Init()
	if err != nil {
		panic(err)
	}
	defer tb.Close()

	tb.HideCursor()

	this.Draw()

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
				this.Player.Move(-1, 0)
			case 'j':
				this.Player.Move(0, 1)
			case 'k':
				this.Player.Move(0, -1)
			case 'l':
				this.Player.Move(1, 0)
			}
		}

		this.Draw()
	}
}
