package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/client/game/level"
	"github.com/niakr1s/nrg-go/src/config"
	"github.com/niakr1s/nrg-go/src/ecs/system"
)

type Game struct {
	systems []system.System
	level   *level.Loader
}

func NewGame() *Game {
	return &Game{systems: make([]system.System, 0), level: level.NewLoader()}
}

// Init ...
func (g *Game) Init() {
	g.systems = append(g.systems,
		system.NewKeyBoard(),
		system.NewMove(config.BoardWidth, config.BoardHeight),
		system.NewContain(config.BoardWidth, config.BoardHeight),
		system.NewBounce(),
		system.NewWeapon(),
		system.NewDamage(),
		system.NewDestroy(config.BoardWidth, config.BoardHeight),
		system.NewClean(),
	)
	g.level.NextLevel()
}

func (g *Game) Update(screen *ebiten.Image) error {
	for _, s := range g.systems {
		s.Step(g.level.Reg)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawBoard(screen)
}
