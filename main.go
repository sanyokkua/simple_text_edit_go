package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"simple_text_editor/core/implementation"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := *implementation.CreateApplicationContextHolderApi()
	jsApi := app.GetJsApi()
	menu := app.GetApplicationMenu()
	newDefaultLogger := logger.NewDefaultLogger()

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
		Menu:          menu.CreateMenu(),
		Bind: []interface{}{
			jsApi,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
		Logger: newDefaultLogger,
	})

	if err != nil {
		newDefaultLogger.Fatal(err.Error())
	}
}
