package animation

import (
	"math"

	"github.com/igarciacracco/climbing-go/core/ecs/components/physics"
	"github.com/igarciacracco/climbing-go/core/ecs/components/visuals"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type EntityHands struct {
	query *donburi.Query
}

func NewEntityHands() *EntityHands {
	return &EntityHands{
		query: donburi.NewQuery(filter.Contains(
			visuals.HandsComponent,
			physics.Velocity,
		)),
	}
}

func (s *EntityHands) Update(ecs *ecs.ECS) {
	s.query.Each(ecs.World, func(entry *donburi.Entry) {
		hands := visuals.HandsComponent.Get(entry)
		velocity := physics.Velocity.Get(entry)

		s.updateHand(&hands.LeftHand, velocity)
		s.updateHand(&hands.RightHand, velocity)
	})
}

func (s *EntityHands) updateHand(hand *visuals.Hand, velocity *physics.VelocityData) {
	speed := math.Sqrt(velocity.X*velocity.X + velocity.Y*velocity.Y)
	if speed > 0.1 {
		hand.RockingTimer += hand.RockingFrequency * speed
		offsetY := math.Sin(hand.RockingTimer) * hand.RockingAmplitude
		hand.RockingOffset.Y = int(offsetY)
	} else if hand.RockingOffset.Y != 0 {
		// Smoothly return to base position
		if math.Abs(float64(hand.RockingOffset.Y)) > 0.1 {
			hand.RockingOffset.Y = int(float64(hand.RockingOffset.Y) * 0.9)
		} else {
			hand.RockingOffset.Y = 0
		}
	}
}
