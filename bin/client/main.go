package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/client"
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	tag "github.com/niakr1s/nrg-go/src/ecs/tags"
	"github.com/niakr1s/nrg-go/src/geo"
	"github.com/niakr1s/nrg-go/src/img"
	"github.com/niakr1s/nrg-go/src/shape"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.TraceLevel)

	client := client.New()
	client.Init()
	// for test
	for i := 0; i < 100; i++ {
		circle := shape.NewCircle(500, 500, 50, img.WhiteCircle)
		player := entity.NewEntity().
			WithComponent(component.ShapeID, circle).
			WithComponent(component.VectorID, geo.NewVector(rand.Float64()*2*3.14)).
			WithTags(tag.PlayerID)
		client.Reg.AddEntity(player)
	}
	circle := shape.NewCircle(500, 500, 50, img.WhiteCircle)
	player := entity.NewEntity().
		WithComponent(component.ShapeID, circle).
		WithTags(tag.PlayerID, tag.UserID)
	client.Reg.AddEntity(player)

	client.StartProduceBoard()

	if err := ebiten.RunGame(client); err != nil {
		log.Fatal(err)
	}
}
