package physics

import "github.com/yohamta/donburi"

type PositionData struct {
	X, Y, Z float64
}

var Position = donburi.NewComponentType[PositionData]()
