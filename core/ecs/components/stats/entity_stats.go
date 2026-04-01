package stats

import "github.com/yohamta/donburi"

type EntityStatsData struct {
	MoveSpeed float64
}

var EntityStatsComponent = donburi.NewComponentType[EntityStatsData]()
