package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/client"
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/geo"
	"github.com/niakr1s/nrg-go/src/img"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.TraceLevel)

	client := client.New()
	client.Init()

	// for test
	circle := geo.NewCircle(500, 500, 50, img.WhiteCircle)
	player := entity.NewEntity().WithComponent(component.DrawableID, circle)
	client.Reg.AddEntity(player)

	if err := ebiten.RunGame(client); err != nil {
		log.Fatal(err)
	}
}
