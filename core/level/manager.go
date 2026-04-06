package level

import (
	"github.com/igarciacracco/climbing-go/core/assets/images"
	"github.com/igarciacracco/climbing-go/core/ecs/archetypes"
	"github.com/igarciacracco/climbing-go/core/ecs/components/visuals"
	"github.com/igarciacracco/climbing-go/core/ecs/systems"
	"github.com/igarciacracco/climbing-go/core/ecs/systems/movement"
	visSystems "github.com/igarciacracco/climbing-go/core/ecs/systems/visuals"
	"github.com/yohamta/donburi/ecs"
)

// Manager is responsible for loading, unloading, and switching between levels.
type Manager struct {
	CurrentLevel *Level
}

// NewManager creates a new level manager and loads the initial level.
func NewManager(loader images.Loader) (*Manager, error) {
	level := New()

	level.ECS().AddSystem(systems.UpdatePhysics)
	level.ECS().AddSystem(movement.NewMovement().Update)
	level.ECS().AddSystem(visSystems.NewAnimationSystem().Update)
	level.ECS().AddSystem(visSystems.NewEntityHands().Update)
	level.ECS().AddSystem(visSystems.NewSpriteVariation().Update)
	level.ECS().AddRenderer(ecs.LayerDefault, systems.DrawRender)

	playerSheet, err := loader.Load("player-forg-g-idle.png")
	if err != nil {
		return nil, err
	}

	playerHandsImage, err := loader.Load("player-forg-hand-g.png")
	if err != nil {
		return nil, err
	}

	frames := visuals.FramesFromSheet(playerSheet, 32, 32)
	anim := func(frameDuration float64) *visuals.AnimationData {
		a := visuals.NewAnimationData(frames, frameDuration, false)
		return &a
	}

	animations := map[string]*visuals.AnimationData{
		visuals.AnimIdle:  anim(6),
		visuals.AnimRun:   anim(6),
		visuals.AnimJump:  anim(8),
		visuals.AnimFall:  anim(8),
		visuals.AnimClimb: anim(10),
	}

	// TODO: fix this, should not be hardcoded
	archetypes.NewPlayer(level.ECS(), 0, animations, playerHandsImage, 250, 200)

	return &Manager{CurrentLevel: level}, nil
}
