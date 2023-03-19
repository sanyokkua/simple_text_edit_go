package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"log"
	"simple_text_editor/core/api"
	"simple_text_editor/core/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := application.CreateApplicationContextHolderApi()
	//editorApp := core.CreateNewApplication() //TODO: remove

	var jsApi api.JsApi = app.GetJsApi()

	err := wails.Run(&options.App{
		Title:  "Simple Text Editor",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:     app.OnStartup,
		OnDomReady:    app.OnDomReady,
		OnShutdown:    app.OnShutdown,
		OnBeforeClose: app.OnBeforeClose,
		Menu:          app.GetApplicationMenu().CreateMenu(),
		Bind: []interface{}{
			//editorApp,
			jsApi,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
