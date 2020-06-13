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
			SetComponents(circle, component.NewPos(500, 500), component.NewVector(rand.Float64()*2*3.14), component.NewSpeed(1)).
			SetTags(tag.PlayerID)
		client.Reg.AddEntity(player)
	}

	for i := 0; i < 10; i++ {
		circle := component.NewCircle(50, img.RedCircle)
		player := entity.NewEntity().
			SetComponents(circle,
				component.NewPos(float64(rand.Intn(500)+100), float64(rand.Intn(500)+100))).
			SetTags(tag.PlayerID, tag.GroundID)
		client.Reg.AddEntity(player)
	}

	player := entity.NewEntity().
		SetComponents(component.NewCircle(50, img.BlueCircle), component.NewPos(500, 500), component.NewSpeed(10)).
		SetTags(tag.UserID, tag.PlayerID)

	client.Reg.AddEntity(player)

	if err := ebiten.RunGame(client); err != nil {
		log.Fatal(err)
	}
}
