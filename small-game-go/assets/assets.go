package assets

import (
	"embed"
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

var PlayerSprite = mustLoadImage("planes/plane.png")
var MeteorSprites = mustLoadImages("meteors")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadImages(dir string) []*ebiten.Image {
	images := []*ebiten.Image{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			images = append(images, mustLoadImage(path))
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error loading images: %v\n", err)
	}

	return images
}
