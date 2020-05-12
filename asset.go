package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func importImageAsset(path string) *ebiten.Image {
	importedImage, _, _  := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	return importedImage
}
