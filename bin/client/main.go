package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/client"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.TraceLevel)

	client := client.New()
	client.Init()

	if err := ebiten.RunGame(client); err != nil {
		log.Fatal(err)
	}
}
