package input

import "github.com/yohamta/donburi"

type InputData struct {
	// Define input fields here, e.g., key presses
}

var InputComponent = donburi.NewComponentType[InputData]()
