package movement

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/igarciacracco/climbing-go/core/ecs/components/input"
	"github.com/igarciacracco/climbing-go/core/ecs/components/physics"
	"github.com/igarciacracco/climbing-go/core/ecs/components/stats"
)

type Movement struct {
	query *donburi.Query
}

func NewMovement() *Movement {
	return &Movement{
		query: donburi.NewQuery(filter.Contains(
			input.InputComponent,
			physics.Velocity,
			stats.EntityStatsComponent,
		)),
	}
}

// TODO: Normalize direction and speed
func (s *Movement) Update(ecs *ecs.ECS) {
	s.query.Each(ecs.World, func(entry *donburi.Entry) {
		velocity := physics.Velocity.Get(entry)
		stats := stats.EntityStatsComponent.Get(entry)

		switch {
		case ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
			velocity.X = -stats.MoveSpeed
		case ebiten.IsKeyPressed(ebiten.KeyArrowRight):
			velocity.X = stats.MoveSpeed
		default:
			velocity.X = 0
		}

		switch {
		case ebiten.IsKeyPressed(ebiten.KeyArrowUp):
			velocity.Y = -stats.MoveSpeed
		case ebiten.IsKeyPressed(ebiten.KeyArrowDown):
			velocity.Y = stats.MoveSpeed
		default:
			velocity.Y = 0
		}

	})
}
