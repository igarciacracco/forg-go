package physics

import "github.com/yohamta/donburi"

type VelocityData struct {
	X, Y, Z float64
}

var Velocity = donburi.NewComponentType[VelocityData]()
