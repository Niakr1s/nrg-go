package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/client"
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/geo"
	"github.com/niakr1s/nrg-go/src/img"
	"github.com/niakr1s/nrg-go/src/shape"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.TraceLevel)

	client := client.New()
	client.Init()
	client.StartProduceBoard()

	// for test
	circle := shape.NewCircle(500, 500, 50, img.WhiteCircle)
	player := entity.NewEntity().
		WithComponent(component.ShapeID, circle).
		WithComponent(component.VectorID, geo.NewVector(0)).
		WithTags(component.PlayerTagID, component.UserTagID)
	client.Reg.AddEntity(player)

	if err := ebiten.RunGame(client); err != nil {
		log.Fatal(err)
	}
}
