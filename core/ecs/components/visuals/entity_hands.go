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

func NewHands(sprite *ebiten.Image) HandsData {
	return HandsData{
		LeftHand: Hand{
			Sprite:           sprite,
			BaseOffset:       image.Point{-20, 0},
			RockingAmplitude: 5,
			RockingFrequency: 0.15,
		},
		RightHand: Hand{
			Sprite:           sprite,
			BaseOffset:       image.Point{20, 0},
			RockingAmplitude: 5,
			RockingFrequency: 0.15,
		},
	}
}

var HandsComponent = donburi.NewComponentType[HandsData]()
