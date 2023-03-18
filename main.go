package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"log"
	"simple_text_editor/core"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	editorApp := core.CreateNewApplication()

	err := wails.Run(&options.App{
		Title:  "Simple Text Editor",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: editorApp.Startup,
		Menu:      editorApp.CreateMenu(),
		Bind: []interface{}{
			editorApp,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
