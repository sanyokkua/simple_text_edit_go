package main

import (
	"embed"
	"log"
	"simple_text_editor/application"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

func main() {
	editorApp := application.CreateNewApplication()

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
			OpenInspectorOnStartup: false,
		},
	})

	println("From main GO APP:")
	println(editorApp.AppContext)

	if err != nil {
		log.Fatal(err.Error())
	}
}
