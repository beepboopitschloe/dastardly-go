package player

import tl "github.com/JoelOtter/termloop"

type Player struct {
	entity *tl.Entity
	level  *tl.BaseLevel
	prevX  int
	prevY  int
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.entity.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.entity.Draw(screen)
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		player.prevX, player.prevY = player.entity.Position()

		switch event.Key {
		case tl.KeyArrowRight:
			player.entity.SetPosition(player.prevX+1, player.prevY)
		case tl.KeyArrowLeft:
			player.entity.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowDown:
			player.entity.SetPosition(player.prevX, player.prevY+1)
		case tl.KeyArrowUp:
			player.entity.SetPosition(player.prevX, player.prevY-1)
		}
	}
}
func (player *Player) Position() (int, int) {
	return player.entity.Position()
}

func (player *Player) Size() (int, int) {
	return player.entity.Size()
}

func (player *Player) Collide(collision tl.Physical) {
	if _, ok := collision.(*tl.Rectangle); ok {
		player.entity.SetPosition(player.prevX, player.prevY)
	}
}

func NewPlayer(level *tl.BaseLevel) (player *Player) {
	player = &Player{
		entity: tl.NewEntity(1, 1, 1, 1),
		level:  level,
	}
	player.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorWhite, Ch: '@'})

	return player
}
