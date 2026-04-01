package images

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

type Loader struct {
	basePath string
}

func NewLoader(
	basePath string,
) Loader {
	return Loader{
		basePath: basePath,
	}
}

func (l *Loader) Load(path string) (*ebiten.Image, error) {
	f, err := os.Open(filepath.Join(l.basePath, path))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}
