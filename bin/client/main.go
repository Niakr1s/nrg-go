package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/client"
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/tag"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.TraceLevel)

	client := client.New()
	client.Init()

	// other players
	for i := 0; i < 1; i++ {
		circle := component.NewCircle(50)
		player := entity.NewEntity().
			SetComponents(
				circle,
				component.NewPos(200, 200),
				component.NewVector(rand.Float64()*2*3.14), component.NewSpeed(1),
				component.NewGround(false),
				component.NewHP(100),
			).
			SetTags(tag.Player)
		client.Reg.AddEntity(player)
	}

	// obstacles
	for i := 0; i < 1; i++ {
		circle := component.NewCircle(50)
		player := entity.NewEntity().
			SetComponents(
				circle,
				component.NewPos(800, 800),
				component.NewGround(true),
			).
			SetTags(tag.Player)
		client.Reg.AddEntity(player)
	}

	// player
	player := entity.NewEntity().
		SetComponents(
			component.NewCircle(50),
			component.NewPos(500, 500),
			component.NewSpeed(10),
			component.NewGround(false),
			component.NewHP(100),
		).
		SetTags(tag.User, tag.Player)
	client.Reg.AddEntity(player)

	go func() {
		for {
			<-time.After(time.Second * 2)
			randPos := component.NewPos(float64(rand.Intn(500)+100), float64(rand.Intn(500)+100))
			bullet := entity.NewEntity().
				SetComponents(
					component.NewCircle(20),
					randPos,
					component.NewVectorFromPos(randPos, component.NewPos(500, 500)),
					component.NewSpeed(5),
					component.NewDamage(10),
				).
				SetTags(tag.Bullet)
			client.Reg.Lock()
			client.Reg.AddEntity(bullet)
			client.Reg.Unlock()
		}
	}()

	if err := ebiten.RunGame(client); err != nil {
		log.Fatal(err)
	}
}
