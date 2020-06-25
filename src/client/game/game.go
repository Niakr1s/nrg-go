package game

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/niakr1s/nrg-go/src/client/game/level"
	"github.com/niakr1s/nrg-go/src/client/state"
	"github.com/niakr1s/nrg-go/src/config"
	"github.com/niakr1s/nrg-go/src/ecs/system"
)

type Game struct {
	systems []system.System
	level   *level.Loader
	status  *system.Status
	next    state.State
}

func NewGame() *Game {
	res := &Game{systems: make([]system.System, 0), level: level.NewLoader()}
	res.init()
	return res
}

// Init ...
func (g *Game) init() {
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
		g.onLevelCompleted()
	}
	if g.status.LevelFailed {
		g.onLevelFail()
	}
	g.status.Reset()
	if isPauseRequested() {
		log.Print("pause is requested")
		g.next = NewPauseMenu(g)
	}
	return nil
}

func isPauseRequested() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEscape)
}

func (g *Game) onLevelCompleted() {
	if !g.level.NextLevel() {
		os.Exit(0)
	}
}

func (g *Game) onLevelFail() {
	g.level.LoadLevel()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawBoard(screen)

}

func (g *Game) Next() state.State {
	return g.next
}
