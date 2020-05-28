package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"sync"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/radovskyb/watcher"
	"github.com/robertkrimen/otto"
	"golang.org/x/image/font"
)

type GameState struct {
	gfxObjects      *[]gfxObject
	gfxObjectsMutex *sync.Mutex
	vm              *otto.Otto
	imageAssets     map[string]*ebiten.Image
}

var (
	tt              *truetype.Font
	mplusMediumFont font.Face
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

func init() {
	fontCache = map[int]*font.Face{}
	tt2, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	tt = tt2
	const dpi = 200
	mplusMediumFont = truetype.NewFace(tt, &truetype.Options{
		Size:    2,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    9,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	mplusBigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func main() {
	go func() { log.Println(http.ListenAndServe("localhost:6060", nil)) }()
	gfxObjects := []gfxObject{}
	gfxObjectsMutex := &sync.Mutex{}
	vm := vmService(&gfxObjects, gfxObjectsMutex)
	w := watcher.New()
	importScripts(vm, w)
	go func() {
		log.Println("[VM] Starting hotreload service.")
		log.Println("[VM] Watching " + strconv.Itoa(len(w.WatchedFiles())) + " files.")
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event)
				code, err := ioutil.ReadFile(event.Path)
				if err != nil {
					log.Fatal(err)
				}
				_, err = vm.Run(string(code))
				if err != nil {
					log.Fatal(err)
				}
				log.Println("[VM] Hotreloaded: " + event.Path)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()
	go w.Start(time.Millisecond * 100)
	imageAssets := map[string]*ebiten.Image{}
	importManifest(&imageAssets)
	gameState := GameState{
		gfxObjects:      &gfxObjects,
		gfxObjectsMutex: gfxObjectsMutex,
		vm:              vm,
		imageAssets:     imageAssets,
	}
	if err := ebiten.Run(gameState.update, 320, 240, 2, "h"); err != nil {
		log.Fatal(err)
	}
}
