package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func importImageAsset(path string) *ebiten.Image {
	importedImage, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return importedImage
}

type Manifest struct {
	Assets struct {
		Images []struct {
			Identifier string `json:"identifier"`
			Path       string `json:"path"`
		} `json:"images"`
		Sounds []interface{} `json:"sounds"`
	} `json:"assets"`
}

func importManifest(imageAssets *map[string]*ebiten.Image) {
	manifestData, err := ioutil.ReadFile("assets/manifest.json")
	if err != nil {
		log.Fatal(err)
	}
	manifest := Manifest{}
	err = json.Unmarshal(manifestData, &manifest)
	if err != nil {
		log.Fatal(err)
	}
	//import image assets
	log.Println("[MANIFEST] Importing " + strconv.Itoa((len(manifest.Assets.Images))) + " image asset(s)...")
	for i := 0; i < len(manifest.Assets.Images); i++ {
		(*imageAssets)[manifest.Assets.Images[i].Identifier] = importImageAsset(manifest.Assets.Images[i].Path)
		log.Println("[MANIFEST] Imported image asset (" + strconv.Itoa(i) + "): " + manifest.Assets.Images[i].Path + " as " + manifest.Assets.Images[i].Identifier + ".")
	}
	log.Println("[MANIFEST] Finished importing assets.")
}
