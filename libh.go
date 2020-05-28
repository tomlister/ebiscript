package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/hajimehoshi/ebiten"
	"github.com/robertkrimen/otto"
)

func importLibh(vm *otto.Otto, gfxObjects *[]gfxObject, gfxObjectsMutex *sync.Mutex, entitlements EntitlementsData) {
	vm.Set("drawDebugText", func(call otto.FunctionCall) otto.Value {
		gfxObjectsMutex.Lock()
		gfxObject := gfxObject{
			Type: "debugtext",
			Data: gfxDebugText{
				Content: call.Argument(0).String(),
			},
		}
		*gfxObjects = append(*gfxObjects, gfxObject)
		gfxObjectsMutex.Unlock()
		return otto.Value{}
	})
	vm.Set("getMouseX", func(call otto.FunctionCall) otto.Value {
		x, _ := ebiten.CursorPosition()
		result, _ := vm.ToValue(x)
		return result
	})
	vm.Set("getMouseY", func(call otto.FunctionCall) otto.Value {
		_, y := ebiten.CursorPosition()
		result, _ := vm.ToValue(y)
		return result
	})
	vm.Set("isLeftMouseDown", func(call otto.FunctionCall) otto.Value {
		result, _ := vm.ToValue(ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft))
		return result
	})
	vm.Set("drawText", func(call otto.FunctionCall) otto.Value {
		gfxObjectsMutex.Lock()
		v3, err := call.Argument(2).ToInteger()
		if err != nil {
			log.Fatal(err)
		}
		v4, err := call.Argument(3).ToInteger()
		if err != nil {
			log.Fatal(err)
		}
		v5, err := call.Argument(4).ToInteger()
		if err != nil {
			log.Fatal(err)
		}
		gfxObject := gfxObject{
			Type: "text",
			Data: gfxText{
				Content: call.Argument(0).String(),
				Color:   call.Argument(1).String(),
				Size:    int(v3),
				X:       int(v4),
				Y:       int(v5),
			},
		}
		*gfxObjects = append(*gfxObjects, gfxObject)
		gfxObjectsMutex.Unlock()
		return otto.Value{}
	})
	vm.Set("drawImage", func(call otto.FunctionCall) otto.Value {
		gfxObjectsMutex.Lock()
		v2, err := call.Argument(1).ToInteger()
		if err != nil {
			log.Fatal(err)
		}
		v3, err := call.Argument(2).ToInteger()
		if err != nil {
			log.Fatal(err)
		}
		gfxObject := gfxObject{
			Type: "image",
			Data: gfxImage{
				Name: call.Argument(0).String(),
				X:    int(v2),
				Y:    int(v3),
			},
		}
		*gfxObjects = append(*gfxObjects, gfxObject)
		gfxObjectsMutex.Unlock()
		return otto.Value{}
	})
	vm.Set("drawSolidImage", func(call otto.FunctionCall) otto.Value {
		gfxObjectsMutex.Lock()
		v2, err := call.Argument(1).ToInteger()
		if err != nil {
			log.Fatal(err)
		}
		v3, err := call.Argument(2).ToInteger()
		if err != nil {
			log.Fatal(err)
		}
		v4, err := call.Argument(3).ToInteger()
		if err != nil {
			log.Fatal(err)
		}
		v5, err := call.Argument(4).ToInteger()
		if err != nil {
			log.Fatal(err)
		}
		v6, err := call.Argument(5).ToFloat()
		if err != nil {
			log.Fatal(err)
		}
		gfxObject := gfxObject{
			Type: "solidimage",
			Data: gfxSolidImage{
				Color: call.Argument(0).String(),
				X:     int(v2),
				Y:     int(v3),
				W:     int(v4),
				H:     int(v5),
				A:     v6,
			},
		}
		*gfxObjects = append(*gfxObjects, gfxObject)
		gfxObjectsMutex.Unlock()
		return otto.Value{}
	})
	vm.Set("clearObjects", func(call otto.FunctionCall) otto.Value {
		gfxObjectsMutex.Lock()
		*gfxObjects = nil
		gfxObjectsMutex.Unlock()
		return otto.Value{}
	})

	vm.Set("httpGET", func(call otto.FunctionCall) otto.Value {
		if entitlements.Entitlements.VMInternet == true {
			resp, err := http.Get(call.Argument(0).String())
			if err != nil {
				rtr := httpResponse{
					Status: resp.StatusCode,
				}
				returnval, _ := vm.ToValue(rtr)
				return returnval
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			rtr := httpResponse{
				Status: resp.StatusCode,
				Body:   string(body),
			}
			returnval, _ := vm.ToValue(rtr)
			return returnval
		} else {
			rtr := httpResponse{
				Status: 403,
				Body:   "VM lacks internet entitlement.",
			}
			returnval, _ := vm.ToValue(rtr)
			return returnval
		}
	})
}

type httpResponse struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}
