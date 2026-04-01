package systems

import (
	"github.com/igarciacracco/climbing-go/core/ecs/components/physics"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var physicsQuery = donburi.NewQuery(filter.Contains(physics.Position, physics.Velocity))

func UpdatePhysics(ecs *ecs.ECS) {
	physicsQuery.Each(ecs.World, func(entry *donburi.Entry) {
		position := physics.Position.Get(entry)
		velocity := physics.Velocity.Get(entry)

		// add delta time
		position.X += velocity.X
		position.Y += velocity.Y
	})
}
