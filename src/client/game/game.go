package game

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/config"
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	"github.com/niakr1s/nrg-go/src/ecs/system"
)

type Game struct {
	Reg *registry.Registry

	systems []system.System
}

func NewGame() *Game {
	return &Game{Reg: registry.NewRegistry(), systems: make([]system.System, 0)}
}

// Init ...
func (g *Game) Init() {
	g.systems = append(g.systems,
		system.NewKeyBoard(g.Reg),
		system.NewMove(g.Reg, config.BoardWidth, config.BoardHeight),
		system.NewContain(g.Reg, config.BoardWidth, config.BoardHeight),
		system.NewBounce(g.Reg),
		system.NewWeapon(g.Reg),
		system.NewDamage(g.Reg),
		system.NewDestroy(g.Reg, config.BoardWidth, config.BoardHeight),
		system.NewClean(g.Reg),
	)

	userWeap := component.NewUserControlledWeapon(component.NewVector(0))
	userWeap.SetDirection(component.NewUserControlledWeaponDirection(component.NewVector(1.5*math.Pi), component.NewVector(0.3*math.Pi)))

	enemyWeap := component.NewAutoWeapon(component.NewVector(0),
		component.NewVector(0),
		component.NewVector(0.5*math.Pi), component.NewVector(math.Pi), component.NewVector(1.5*math.Pi))
	enemyWeap.SetDirection(component.NewAutoWeaponDirection(1.5*math.Pi, component.NewVector(0.3*math.Pi)))

	g.Reg.AddEntity(entity.NewUser(component.NewPos(500, 500)).SetComponents(userWeap))
	g.Reg.AddEntity(entity.NewEnemy(component.NewPos(200, 200)).SetComponents(enemyWeap))
	g.Reg.AddEntity(entity.NewObstacle(component.NewPos(700, 200)))
}

func (g *Game) Update(screen *ebiten.Image) error {
	for _, s := range g.systems {
		s.Step()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawBoard(screen)
}
