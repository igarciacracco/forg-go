package physics

import "github.com/yohamta/donburi"

type RotationData struct {
	X, Y, Z float64
}

var Rotation = donburi.NewComponentType[RotationData]()
