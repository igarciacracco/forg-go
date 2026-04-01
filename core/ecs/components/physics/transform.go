package physics

import "github.com/yohamta/donburi"

type TransformData struct {
	Position PositionData
	Rotation RotationData
	Scale    ScaleData
}

var Transform = donburi.NewComponentType[TransformData]()
