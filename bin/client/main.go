package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/client"
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.TraceLevel)

	client := client.New()
	client.Init()

	client.Reg.AddEntity(entity.NewUser(component.NewPos(500, 500)).
		SetComponents(component.NewWeapon(component.NewVector(0),
			component.NewVector(0), component.NewVector(math.Pi))),
	)

	client.Reg.AddEntity(entity.NewEnemy(component.NewPos(200, 200)).
		SetComponents(component.NewWeapon(component.NewVector(0),
			component.NewVector(0), component.NewVector(0.5*math.Pi), component.NewVector(math.Pi), component.NewVector(1.5*math.Pi))),
	)

	client.Reg.AddEntity(entity.NewObstacle(component.NewPos(700, 200)))

	if err := ebiten.RunGame(client); err != nil {
		log.Fatal(err)
	}
}
