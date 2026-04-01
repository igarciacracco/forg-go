package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/igarciacracco/climbing-go/core/assets/images"
	ebt "github.com/igarciacracco/climbing-go/core/ebiten"
)

func main() {

	loader := images.NewLoader("./assets/images/")

	game, err := ebt.NewGame(loader)
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
