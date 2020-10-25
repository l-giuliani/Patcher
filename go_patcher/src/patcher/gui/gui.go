package gui

import (
	"fmt"
	"log"
	//"os"
	"encoding/json"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	dto "patcher/gui/dto"
	guiPatchService "patcher/gui/services"
)

func StartGui(){

	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	/*currPath, err2 := os.Getwd()
	if err2 != nil {
		l.Fatal(err2)
	}*/

	// Create astilectron
	a, err := astilectron.New(nil, astilectron.Options{
		AppName:           "Patcher",
		BaseDirectoryPath: "resources",
		//AppIconDefaultPath: currPath + "/resources/images/logo2.png",
	})
	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer a.Close()

	// Handle signals
	a.HandleSignals()

	// Start
	if err = a.Start(); err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
		return
	}

	// New window
	var w *astilectron.Window
	if w, err = a.NewWindow("resources/index.html", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(720),
		Width:  astikit.IntPtr(600),
	}); err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
		return
	}

	// Create windows
	if err = w.Create(); err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
		return
	}

	w.OnMessage(func(m *astilectron.EventMessage) interface{} {
			// Unmarshal
			request := dto.Request{}
			m.Unmarshal(&request)
			//fmt.Println(request)

			payload := dto.GeneratePayload{}
			json.Unmarshal([]byte(request.Data), &payload)
			//fmt.Println(payload)

			res := guiPatchService.GuiGeneratePatch(request.RequestType, payload)

			return res
	})

	//w.OpenDevTools()

	// Blocking pattern
	a.Wait()

	fmt.Println("END")
}