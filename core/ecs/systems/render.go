package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/igarciacracco/climbing-go/core/ecs/components/physics"
	"github.com/igarciacracco/climbing-go/core/ecs/components/visuals"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var renderQuery = donburi.NewQuery(filter.Contains(visuals.Sprite, physics.Position))
var animationRenderQuery = donburi.NewQuery(filter.Contains(visuals.Animation, physics.Position))
var animManagerRenderQuery = donburi.NewQuery(filter.Contains(visuals.AnimManager, physics.Position))

// TODO: this is ugly
var handsQuery = donburi.NewQuery(filter.Contains(visuals.HandsComponent, physics.Position))

func DrawRender(ecs *ecs.ECS, screen *ebiten.Image) {
	renderQuery.Each(ecs.World, func(entry *donburi.Entry) {
		sprite := visuals.Sprite.Get(entry)
		position := physics.Position.Get(entry)

		op := &ebiten.DrawImageOptions{}
		if sprite.FlipX {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(float64(sprite.Width), 0)
		}
		op.GeoM.Translate(position.X, position.Y)
		screen.DrawImage(sprite.Image, op)
	})
	animationRenderQuery.Each(ecs.World, func(entry *donburi.Entry) {
		anim := visuals.Animation.Get(entry)
		position := physics.Position.Get(entry)
		if len(anim.Frames) == 0 {
			return
		}
		op := &ebiten.DrawImageOptions{}
		if anim.FlipX {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(float64(anim.Width), 0)
		}
		op.GeoM.Translate(position.X, position.Y)
		screen.DrawImage(anim.Frames[anim.CurrentFrame], op)
	})
	animManagerRenderQuery.Each(ecs.World, func(entry *donburi.Entry) {
		manager := visuals.AnimManager.Get(entry)
		position := physics.Position.Get(entry)
		anim := manager.Active()
		if anim == nil || len(anim.Frames) == 0 {
			return
		}
		op := &ebiten.DrawImageOptions{}
		if manager.FlipX {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(float64(anim.Width), 0)
		}
		op.GeoM.Translate(position.X, position.Y)
		screen.DrawImage(anim.Frames[anim.CurrentFrame], op)
	})
	handsQuery.Each(ecs.World, func(entry *donburi.Entry) {
		hands := visuals.HandsComponent.Get(entry)
		position := physics.Position.Get(entry)

		// Draw Left Hand
		opLeft := &ebiten.DrawImageOptions{}
		opLeft.GeoM.Translate(position.X+float64(hands.LeftHand.BaseOffset.X)+float64(hands.LeftHand.RockingOffset.X), position.Y+float64(hands.LeftHand.BaseOffset.Y)+float64(hands.LeftHand.RockingOffset.Y))
		screen.DrawImage(hands.LeftHand.Sprite, opLeft)

		// Draw Right Hand
		opRight := &ebiten.DrawImageOptions{}
		opRight.GeoM.Translate(position.X+float64(hands.RightHand.BaseOffset.X)+float64(hands.RightHand.RockingOffset.X), position.Y+float64(hands.RightHand.BaseOffset.Y)+float64(hands.RightHand.RockingOffset.Y))
		screen.DrawImage(hands.RightHand.Sprite, opRight)
	})
}
