package visuals

import (
	"github.com/igarciacracco/forg-go/core/ecs/components/visuals"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type AnimationSystem struct {
	query        *donburi.Query
	managerQuery *donburi.Query
}

func NewAnimationSystem() *AnimationSystem {
	return &AnimationSystem{
		query:        donburi.NewQuery(filter.Contains(visuals.Animation)),
		managerQuery: donburi.NewQuery(filter.Contains(visuals.AnimManager)),
	}
}

func (a *AnimationSystem) Update(ecs *ecs.ECS) {
	a.query.Each(ecs.World, func(entry *donburi.Entry) {
		anim := visuals.Animation.Get(entry)
		if len(anim.Frames) <= 1 {
			return
		}
		anim.FrameTimer++
		if anim.FrameTimer >= anim.FrameDuration {
			anim.FrameTimer = 0
			anim.CurrentFrame = (anim.CurrentFrame + 1) % len(anim.Frames)
		}
	})

	a.managerQuery.Each(ecs.World, func(entry *donburi.Entry) {
		manager := visuals.AnimManager.Get(entry)
		anim := manager.Active()
		if anim == nil || len(anim.Frames) <= 1 {
			return
		}
		anim.FrameTimer++
		if anim.FrameTimer >= anim.FrameDuration {
			anim.FrameTimer = 0
			anim.CurrentFrame = (anim.CurrentFrame + 1) % len(anim.Frames)
		}
	})
}
