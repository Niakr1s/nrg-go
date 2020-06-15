package main

import (
	"math/rand"
	"time"

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

	// other players
	for i := 0; i < 10; i++ {
		circle := component.NewCircle(50, img.WhiteCircle)
		player := entity.NewEntity().
			SetComponents(
				circle,
				component.NewPos(500, 500),
				component.NewVector(rand.Float64()*2*3.14), component.NewSpeed(1),
				component.NewGround(false),
				component.NewHP(100),
			).
			SetTags(tag.PlayerID)
		client.Reg.AddEntity(player)
	}

	// obstacles
	for i := 0; i < 10; i++ {
		circle := component.NewCircle(50, img.RedCircle)
		player := entity.NewEntity().
			SetComponents(
				circle,
				component.NewPos(float64(rand.Intn(500)+100), float64(rand.Intn(500)+100)),
				component.NewGround(true),
			).
			SetTags(tag.PlayerID)
		client.Reg.AddEntity(player)
	}

	// player
	player := entity.NewEntity().
		SetComponents(
			component.NewCircle(50, img.BlueCircle),
			component.NewPos(500, 500),
			component.NewSpeed(10),
			component.NewGround(false),
			component.NewHP(100),
		).
		SetTags(tag.UserID, tag.PlayerID)
	client.Reg.AddEntity(player)

	go func() {
		for {
			player.RLock()
			hp := player.GetComponents(component.HpID)
			if hp == nil {
				return
			}
			log.Tracef("Player's hp: %d", hp[0].(component.HP).Current)
			player.RUnlock()
			<-time.After(time.Second)
		}
	}()

	go func() {
		for {
			randPos := component.NewPos(float64(rand.Intn(500)+100), float64(rand.Intn(500)+100))
			bullet := entity.NewEntity().
				SetComponents(
					component.NewCircle(20, img.RedCircle),
					randPos,
					component.NewVectorFromPos(randPos, component.NewPos(500, 500)),
					component.NewSpeed(5),
					component.NewDamage(10),
				)
			client.Reg.Lock()
			client.Reg.AddEntity(bullet)
			client.Reg.Unlock()
			<-time.After(time.Millisecond * 200)
		}
	}()

	if err := ebiten.RunGame(client); err != nil {
		log.Fatal(err)
	}
}
