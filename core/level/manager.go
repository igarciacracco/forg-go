package level

import (
	"github.com/igarciacracco/climbing-go/core/assets/images"
	"github.com/igarciacracco/climbing-go/core/ecs/archetypes"
	"github.com/igarciacracco/climbing-go/core/ecs/systems"
	"github.com/igarciacracco/climbing-go/core/ecs/systems/animation"
	"github.com/igarciacracco/climbing-go/core/ecs/systems/movement"
	"github.com/yohamta/donburi/ecs"
)

// Manager is responsible for loading, unloading, and switching between levels.
type Manager struct {
	CurrentLevel *Level
}

// NewManager creates a new level manager and loads the initial level.
func NewManager(loader images.Loader) (*Manager, error) {
	// Create the first level.
	level := New()

	// Add systems to the level's ECS.
	level.ECS().AddSystem(systems.UpdatePhysics)
	level.ECS().AddSystem(movement.NewMovement().Update)
	level.ECS().AddRenderer(ecs.LayerDefault, systems.DrawRender)
	level.ECS().AddSystem(animation.NewEntityHands().Update)

	// Load assets and create entities for this level.
	playerImage, err := loader.Load("player-forg-g.png")

	playerHandsImage, err := loader.Load("player-forg-hand-g.png")
	if err != nil {
		return nil, err
	}
	// TODO: fix this, should not be hardcoded
	archetypes.NewPlayer(level.ECS(), 0, playerImage, playerHandsImage, 250, 200)

	return &Manager{CurrentLevel: level}, nil
}
