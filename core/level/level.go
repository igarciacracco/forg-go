package level

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

// Map represents the game map's structure and properties.
type Map struct {
	// TileData could be a 2D slice representing the map layout.
	TileData [][]int
}

// Level contains all the data for a single game level, including the ECS world.
type Level struct {
	ecs     *ecs.ECS
	gameMap *Map
}

// New creates a new Level, initializing an empty map and a new ECS world.
func New() *Level {
	return &Level{
		ecs:     ecs.NewECS(donburi.NewWorld()),
		gameMap: &Map{},
	}
}

// ECS returns the underlying ECS instance for the level.
func (l *Level) ECS() *ecs.ECS {
	return l.ecs
}
