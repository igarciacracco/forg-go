package visuals

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type AnimationData struct {
	Frames        []*ebiten.Image
	Width         int64
	Height        int64
	CurrentFrame  int
	FrameTimer    float64
	FrameDuration float64 // ticks per frame
	Directionless bool
	FlipX         bool
}

func NewAnimationData(frames []*ebiten.Image, frameDuration float64, directionless bool) AnimationData {
	var w, h int64
	if len(frames) > 0 {
		w = int64(frames[0].Bounds().Dx())
		h = int64(frames[0].Bounds().Dy())
	}
	return AnimationData{
		Frames:        frames,
		Width:         w,
		Height:        h,
		FrameDuration: frameDuration,
		Directionless: directionless,
	}
}

// FramesFromSheet slices a horizontal spritesheet into individual frames.
// frameW and frameH are the pixel dimensions of a single frame.
func FramesFromSheet(sheet *ebiten.Image, frameW, frameH int) []*ebiten.Image {
	bounds := sheet.Bounds()
	cols := bounds.Dx() / frameW
	rows := bounds.Dy() / frameH
	frames := make([]*ebiten.Image, 0, cols*rows)
	for row := range rows {
		for col := range cols {
			x, y := col*frameW, row*frameH
			frames = append(frames, sheet.SubImage(image.Rect(x, y, x+frameW, y+frameH)).(*ebiten.Image))
		}
	}
	return frames
}

var Animation = donburi.NewComponentType[AnimationData]()
