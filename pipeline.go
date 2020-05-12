package main

import (
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/icza/gox/imagex/colorx"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
)

type gfxObject struct {
	Hidden bool
	Type   string
	Data   interface{}
}

type gfxDebugText struct {
	Content string
}

type gfxText struct {
	Color   string
	Content string
	Size    int
	X       int
	Y       int
}

type gfxHotReloaded struct {
	Life int
}

type gfxImage struct {
	Name string
	X    int
	Y    int
}

type gfxSolidImage struct {
	Color string
	X     int
	Y     int
	W     int
	H     int
	A     float64
}

var (
	fontCache map[int]*font.Face
)

func (gameState GameState) update(screen *ebiten.Image) error {
	gameState.vm.Run("main()")
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	gameState.gfxObjectsMutex.Lock()
	for i := 0; i < len((*gameState.gfxObjects)); i++ {
		if (*gameState.gfxObjects)[i].Hidden == false {
			if (*gameState.gfxObjects)[i].Type == "debugtext" {
				textObj := (*gameState.gfxObjects)[i].Data.(gfxDebugText)
				ebitenutil.DebugPrint(screen, textObj.Content)
			} else if (*gameState.gfxObjects)[i].Type == "hotreloaded" {
				textObj := (*gameState.gfxObjects)[i].Data.(gfxHotReloaded)
				if textObj.Life > 0 {
					textObj.Life = textObj.Life - 10
				}
				if textObj.Life <= 0 {
					(*gameState.gfxObjects)[i].Hidden = true
				}
				(*gameState.gfxObjects)[i].Data = textObj
				ebitenutil.DebugPrint(screen, "\n\nHot reloaded!")
			} else if (*gameState.gfxObjects)[i].Type == "text" {
				textObj := (*gameState.gfxObjects)[i].Data.(gfxText)
				color, err := colorx.ParseHexColor(textObj.Color)
				if err != nil {
					log.Fatal(err)
				}
				if fontCache[textObj.Size] == nil {
					font := truetype.NewFace(tt, &truetype.Options{
						Size:    float64(textObj.Size),
						DPI:     200,
						Hinting: font.HintingFull,
					})
					fontCache[textObj.Size] = &font
				}
				text.Draw(screen, textObj.Content, (*fontCache[textObj.Size]), textObj.X, textObj.Y, color)
			} else if (*gameState.gfxObjects)[i].Type == "image" {
				imageObj := (*gameState.gfxObjects)[i].Data.(gfxImage)
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(imageObj.X), float64(imageObj.Y))
				screen.DrawImage(gameState.imageAssets[imageObj.Name], opts)
			} else if (*gameState.gfxObjects)[i].Type == "solidimage" {
				imageObj := (*gameState.gfxObjects)[i].Data.(gfxSolidImage)
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(imageObj.X), float64(imageObj.Y))
				opts.ColorM.Translate(0, 0, 0, -imageObj.A)
				image, err := ebiten.NewImage(imageObj.W, imageObj.H, ebiten.FilterDefault)
				if err != nil {
					log.Fatal(err)
				}
				color, err := colorx.ParseHexColor(imageObj.Color)
				if err != nil {
					log.Fatal(err)
				}
				image.Fill(color)
				screen.DrawImage(image, opts)
				image.Dispose()
			}
		}
	}
	gameState.gfxObjectsMutex.Unlock()
	return nil
}
