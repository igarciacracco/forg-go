package physics

import "github.com/yohamta/donburi"

type ScaleData struct {
	X, Y float64
}

var Scale = donburi.NewComponentType[ScaleData]()
