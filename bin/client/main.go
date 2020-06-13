package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/client"
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
	"github.com/niakr1s/nrg-go/src/img"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.TraceLevel)

	client := client.New()
	client.Init()
	// for test
	for i := 0; i < 100; i++ {
		circle := component.NewCircle(50, img.WhiteCircle)
		player := entity.NewEntity().
			WithComponent(component.ShapeID, circle).
			WithComponent(component.PosID, component.NewPos(500, 500)).
			WithComponent(component.VectorID, component.NewVector(rand.Float64()*2*3.14)).
			WithComponent(component.SpeedID, component.NewSpeed(1)).
			WithTags(tag.PlayerID)
		client.Reg.AddEntity(player)
	}
	circle := component.NewCircle(50, img.WhiteCircle)
	player := entity.NewEntity().
		WithComponent(component.ShapeID, circle).
		WithComponent(component.PosID, component.NewPos(500, 500)).
		WithComponent(component.SpeedID, component.NewSpeed(1)).
		WithTags(tag.PlayerID, tag.UserID)
	client.Reg.AddEntity(player)

	if err := ebiten.RunGame(client); err != nil {
		log.Fatal(err)
	}
}
