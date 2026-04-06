package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/igarciacracco/forg-go/core/assets/images"
	ebt "github.com/igarciacracco/forg-go/core/ebiten"
	"github.com/igarciacracco/forg-go/infrastructure/config"
)

func main() {
	settings, balance, err := config.LoadSettings()
	if err != nil {
		log.Fatal(err)
	}

	loader := images.NewLoader(balance.Assets.Path)

	game, err := ebt.NewGame(loader, settings, balance)
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(settings.Window.Width, settings.Window.Height)
	ebiten.SetWindowTitle(settings.Window.Title)
	ebiten.SetFullscreen(settings.Window.Fullscreen)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
