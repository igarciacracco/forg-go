package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/igarciacracco/forg-go/core/assets/images"
	"github.com/igarciacracco/forg-go/core/level"
	"github.com/igarciacracco/forg-go/infrastructure/config"
)

type Game struct {
	levelManager *level.Manager
	screenWidth  int
	screenHeight int
}

func NewGame(loader images.Loader, settings config.Settings, balance config.Balance) (*Game, error) {
	levelManager, err := level.NewManager(loader, balance)
	if err != nil {
		return nil, err
	}

	return &Game{
		levelManager: levelManager,
		screenWidth:  settings.Display.ScreenWidth,
		screenHeight: settings.Display.ScreenHeight,
	}, nil
}

func (g *Game) Update() error {
	g.levelManager.CurrentLevel.ECS().Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	g.levelManager.CurrentLevel.ECS().Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}
