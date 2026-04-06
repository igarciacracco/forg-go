package visuals

import "github.com/yohamta/donburi"

const (
	AnimIdle  = "idle"
	AnimRun   = "run"
	AnimJump  = "jump"
	AnimFall  = "fall"
	AnimClimb = "climb"
)

type AnimationManager struct {
	Animations    map[string]*AnimationData
	Current       string
	Directionless bool
	FlipX         bool
}

func NewAnimationManager(animations map[string]*AnimationData, initial string, directionless bool) AnimationManager {
	return AnimationManager{
		Animations:    animations,
		Current:       initial,
		Directionless: directionless,
	}
}

// Play switches to the named animation, resetting its frame state.
// Does nothing if the animation is already active or the name is unknown.
func (m *AnimationManager) Play(name string) {
	if m.Current == name {
		return
	}
	if anim, ok := m.Animations[name]; ok {
		anim.CurrentFrame = 0
		anim.FrameTimer = 0
		m.Current = name
	}
}

// Active returns the currently playing AnimationData, or nil if unset.
func (m *AnimationManager) Active() *AnimationData {
	return m.Animations[m.Current]
}

var AnimManager = donburi.NewComponentType[AnimationManager]()
