package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/igarciacracco/climbing-go/core/assets/images"
	"github.com/igarciacracco/climbing-go/core/level"
)

const SCREEN_WIDTH = 640
const SCREEN_HEIGHT = 480

type Game struct {
	levelManager *level.Manager
}

func NewGame(loader images.Loader) (*Game, error) {
	// Initialize the level manager, which handles all level-specific setup.
	levelManager, err := level.NewManager(loader)
	if err != nil {
		return nil, err
	}

	return &Game{levelManager: levelManager}, nil
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
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
