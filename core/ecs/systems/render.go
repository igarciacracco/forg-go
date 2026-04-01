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

// TODO: this is ugly
var handsQuery = donburi.NewQuery(filter.Contains(visuals.HandsComponent, physics.Position))

func DrawRender(ecs *ecs.ECS, screen *ebiten.Image) {
	renderQuery.Each(ecs.World, func(entry *donburi.Entry) {
		sprite := visuals.Sprite.Get(entry)
		position := physics.Position.Get(entry)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(position.X, position.Y)
		screen.DrawImage(sprite.Image, op)
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
