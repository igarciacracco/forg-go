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
func NewPlayer(ecs *ecs.ECS, l ecs.LayerID, image *ebiten.Image, handsImage *ebiten.Image, x, y float64) *donburi.Entry {
	// Create the entity with its components
	entry := ecs.World.Entry(
		ecs.Create(
			l,
			visuals.Sprite,
			physics.Position,
			physics.Velocity,
			input.InputComponent,
			stats.EntityStatsComponent,
			visuals.HandsComponent,
		),
	)

	// Set the data for the Sprite component
	donburi.SetValue(entry, visuals.Sprite, visuals.SpriteData{
		Image: image, // TODO: Get from other source, not hardcoded
	})

	// Set the data for the Position component
	donburi.SetValue(entry, physics.Position, physics.PositionData{
		X: x,
		Y: y,
	})

	// Set the data for the Velocity component.
	// We'll give it a slight downward velocity to show the physics system is working.
	donburi.SetValue(entry, physics.Velocity, physics.VelocityData{
		X: 0,
		Y: 0,
	})

	// Set the data for the Input component
	donburi.SetValue(entry, input.InputComponent, input.InputData{})

	// Set the data for the EntityStats component
	donburi.SetValue(entry, stats.EntityStatsComponent, stats.EntityStatsData{
		MoveSpeed: 3.0,
	})

	// Set the data for hands
	donburi.SetValue(entry, visuals.HandsComponent, visuals.NewHands(
		handsImage,
	))

	return entry
}
