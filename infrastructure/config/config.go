package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Settings holds user-facing configuration (options menu territory).
type Settings struct {
	Window  WindowSettings  `yaml:"window"`
	Display DisplaySettings `yaml:"display"`
}

type WindowSettings struct {
	Width      int    `yaml:"width"`
	Height     int    `yaml:"height"`
	Title      string `yaml:"title"`
	Fullscreen bool   `yaml:"fullscreen"`
}

type DisplaySettings struct {
	ScreenWidth  int `yaml:"screen_width"`
	ScreenHeight int `yaml:"screen_height"`
}

// Balance holds internal game tuning values (not exposed to players).
type Balance struct {
	Player    PlayerBalance    `yaml:"player"`
	Animation AnimationBalance `yaml:"animation"`
	Assets    AssetsBalance    `yaml:"assets"`
}

type PlayerBalance struct {
	StartX    float64 `yaml:"start_x"`
	StartY    float64 `yaml:"start_y"`
	MoveSpeed float64 `yaml:"move_speed"`
}

type AnimationBalance struct {
	IdleFrameDuration  float64 `yaml:"idle_frame_duration"`
	RunFrameDuration   float64 `yaml:"run_frame_duration"`
	JumpFrameDuration  float64 `yaml:"jump_frame_duration"`
	FallFrameDuration  float64 `yaml:"fall_frame_duration"`
	ClimbFrameDuration float64 `yaml:"climb_frame_duration"`
}

type AssetsBalance struct {
	Path string `yaml:"path"`
}

func loadYAML(path string, out any) error {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, out)
}

// LoadSettings reads user_data/settings.yaml from the given path.
// Returns defaults if the file is missing.
func LoadSettings() (Settings, Balance, error) {
	cfg := Settings{}
	balance := Balance{}
	if err := loadYAML("infrastructure/config/user_data/settings.yaml", &cfg); err != nil {
		return Settings{}, Balance{}, err
	}
	if err := loadYAML("infrastructure/config/data/balance.yaml", &balance); err != nil {
		return Settings{}, Balance{}, err
	}
	return cfg, balance, nil
}
