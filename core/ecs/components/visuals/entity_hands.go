package visuals

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type HandSide int

const (
	LeftHand HandSide = iota
	RightHand
)

type Hand struct {
	Sprite            *ebiten.Image
	BaseOffset        image.Point
	RockingOffset     image.Point
	RockingAmplitude  float64
	RockingFrequency  float64
	RockingTimer      float64
	IsEquipped        bool
	CurrentEquipped   *ebiten.Image
	EquippedOffset    image.Point
	EquippedAngle     float64
	EquippedScale     float64
	EquippedIsFlipped bool
}

type HandsData struct {
	LeftHand  Hand
	RightHand Hand
}

func NewHands(sprite *ebiten.Image, parentSize image.Rectangle) HandsData {
	offset := image.Point{
		X: int(float64(parentSize.Dx()) * float64(2) / 3),
		Y: int(float64(parentSize.Dx()) * float64(2) / 3),
	}
	rockingAmplitude := int(float64(parentSize.Dy()) * float64(1) / 10)

	return HandsData{
		LeftHand: Hand{
			Sprite:           sprite,
			BaseOffset:       image.Point{0, offset.Y},
			RockingAmplitude: float64(rockingAmplitude),
			RockingFrequency: 0.15,
		},
		RightHand: Hand{
			Sprite:           sprite,
			BaseOffset:       image.Point{offset.X, offset.Y},
			RockingAmplitude: float64(rockingAmplitude),
			RockingFrequency: 0.15,
		},
	}
}

var HandsComponent = donburi.NewComponentType[HandsData]()
