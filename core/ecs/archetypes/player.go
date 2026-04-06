package archetypes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/igarciacracco/climbing-go/core/ecs/components/input"
	"github.com/igarciacracco/climbing-go/core/ecs/components/physics"
	"github.com/igarciacracco/climbing-go/core/ecs/components/stats"
	"github.com/igarciacracco/climbing-go/core/ecs/components/visuals"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

// NewPlayer creates a player entity with all its necessary components.
// animations is a map of named AnimationData (use the visuals.Anim* constants as keys).
// The manager starts on visuals.AnimIdle.
func NewPlayer(ecs *ecs.ECS, l ecs.LayerID, animations map[string]*visuals.AnimationData, handsImage *ebiten.Image, x, y float64) *donburi.Entry {
	entry := ecs.World.Entry(
		ecs.Create(
			l,
			visuals.AnimManager,
			physics.Position,
			physics.Velocity,
			physics.Scale,
			input.InputComponent,
			stats.EntityStatsComponent,
			visuals.HandsComponent,
		),
	)

	donburi.SetValue(entry, visuals.AnimManager, visuals.NewAnimationManager(animations, visuals.AnimIdle, false))

	donburi.SetValue(entry, physics.Position, physics.PositionData{X: x, Y: y})
	donburi.SetValue(entry, physics.Velocity, physics.VelocityData{X: 0, Y: 0})
	donburi.SetValue(entry, input.InputComponent, input.InputData{})
	donburi.SetValue(entry, stats.EntityStatsComponent, stats.EntityStatsData{MoveSpeed: 3.0})

	// Derive hand offsets from the idle animation's first frame
	idleFrames := animations[visuals.AnimIdle].Frames
	donburi.SetValue(entry, visuals.HandsComponent, visuals.NewHands(handsImage, idleFrames[0].Bounds()))

	return entry
}
