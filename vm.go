package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/radovskyb/watcher"

	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
)

func importScripts(vm *otto.Otto, w *watcher.Watcher) {
	files := []string{}
	rootdir := "bin/scripts"
	err := filepath.Walk(rootdir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".js" {
			files = append(files, path)
			return nil
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		code, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		_, err = vm.Run(string(code))
		if err != nil {
			log.Fatal(err)
		}
		if err := w.Add(file); err != nil {
			log.Fatalln(err)
		}
		log.Println("[VM] Imported script: " + file)
	}
}

func vmService(gfxObjects *[]gfxObject, gfxObjectsMutex *sync.Mutex, entitlements EntitlementsData) *otto.Otto {
	vm := otto.New()
	log.Println("[VM] Providing engine bindings: libh")
	importLibh(vm, gfxObjects, gfxObjectsMutex, entitlements)
	return vm
}

type EntitlementsData struct {
	Entitlements struct {
		VMInternet bool     `json:"vm_internet"`
		VMFiles    []string `json:"vm_files"`
	} `json:"entitlements"`
}

func loadEntitlements() EntitlementsData {
	entitlementsData, err := ioutil.ReadFile("entitlements.json")
	if err != nil {
		log.Fatal(err)
	}
	entitlements := EntitlementsData{}
	err = json.Unmarshal(entitlementsData, &entitlements)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[VM] Loaded entitlements.")
	return entitlements
}
