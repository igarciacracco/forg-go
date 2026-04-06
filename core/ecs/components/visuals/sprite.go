package visuals

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SpriteData struct {
	Image         *ebiten.Image
	Width, Height int64
	Directionless bool
	FlipX         bool
}

func NewSpriteData(image *ebiten.Image, directionless bool) SpriteData {
	return SpriteData{
		Image:         image,
		Width:         int64(image.Bounds().Dx()),
		Height:        int64(image.Bounds().Dy()),
		Directionless: false,
		FlipX:         false,
	}
}

var Sprite = donburi.NewComponentType[SpriteData]()
