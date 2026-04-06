package visuals

import (
	"github.com/igarciacracco/climbing-go/core/ecs/components/physics"
	"github.com/igarciacracco/climbing-go/core/ecs/components/visuals"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type SpriteVariation struct {
	spriteQuery   *donburi.Query
	animQuery     *donburi.Query
	managerQuery  *donburi.Query
}

func NewSpriteVariation() *SpriteVariation {
	return &SpriteVariation{
		spriteQuery:  donburi.NewQuery(filter.Contains(physics.Velocity, visuals.Sprite)),
		animQuery:    donburi.NewQuery(filter.Contains(physics.Velocity, visuals.Animation)),
		managerQuery: donburi.NewQuery(filter.Contains(physics.Velocity, visuals.AnimManager)),
	}
}

func (s *SpriteVariation) Update(ecs *ecs.ECS) {
	s.spriteQuery.Each(ecs.World, func(entry *donburi.Entry) {
		velocity := physics.Velocity.Get(entry)
		sprite := visuals.Sprite.Get(entry)
		if sprite.Directionless {
			return
		}
		if velocity.X < 0 {
			sprite.FlipX = true
		} else if velocity.X > 0 {
			sprite.FlipX = false
		}
	})

	s.animQuery.Each(ecs.World, func(entry *donburi.Entry) {
		velocity := physics.Velocity.Get(entry)
		anim := visuals.Animation.Get(entry)
		if anim.Directionless {
			return
		}
		if velocity.X < 0 {
			anim.FlipX = true
		} else if velocity.X > 0 {
			anim.FlipX = false
		}
	})

	s.managerQuery.Each(ecs.World, func(entry *donburi.Entry) {
		velocity := physics.Velocity.Get(entry)
		manager := visuals.AnimManager.Get(entry)
		if manager.Directionless {
			return
		}
		if velocity.X < 0 {
			manager.FlipX = true
		} else if velocity.X > 0 {
			manager.FlipX = false
		}
	})
}
