package game

import (
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/client/game/level"
	"github.com/niakr1s/nrg-go/src/config"
	"github.com/niakr1s/nrg-go/src/ecs/system"
)

type Game struct {
	systems []system.System
	level   *level.Loader
	status  *system.Status
}

func NewGame() *Game {
	return &Game{systems: make([]system.System, 0), level: level.NewLoader()}
}

// Init ...
func (g *Game) Init() {
	g.status = system.NewStatus()
	g.systems = append(g.systems,
		system.NewKeyBoard(),
		system.NewMove(config.BoardWidth, config.BoardHeight),
		system.NewContain(config.BoardWidth, config.BoardHeight),
		system.NewBounce(),
		system.NewWeapon(),
		system.NewDamage(),
		system.NewDestroy(config.BoardWidth, config.BoardHeight),
		system.NewClean(),
		g.status,
	)
	g.level.LoadLevel()
}

func (g *Game) Update(screen *ebiten.Image) error {
	for _, s := range g.systems {
		s.Step(g.level.Reg)
	}
	if g.status.LevelCompleted {
		if !g.level.NextLevel() {
			os.Exit(0)
		}
	}
	if g.status.LevelFailed {
		g.level.LoadLevel()
	}
	g.status.Reset()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawBoard(screen)

}
