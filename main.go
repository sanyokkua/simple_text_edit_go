package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"simple_text_editor/core/v2"
	"simple_text_editor/core/v2/api"
	"simple_text_editor/core/v2/utils"
)

//go:embed all:frontend/dist
var assets embed.FS

const FileTypesFileName = "fileTypes.json"

func main() {
	newDefaultLogger := logger.NewDefaultLogger()

	extensions := createMappingForTypes()

	app := v2.CreateApplication(&extensions)
	frontendApi := app.GetFrontendApi()
	menuApi := app.GetMenuApi()

	appErr := wails.Run(&options.App{
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
		Menu:          menuApi.CreateMenu(),
		Bind: []interface{}{
			frontendApi,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
		Logger: newDefaultLogger,
	})

	if appErr != nil {
		newDefaultLogger.Fatal(appErr.Error())
	}
}

func createMappingForTypes() map[string]api.FileTypesJsonStruct {
	typesJson := utils.LoadFileTypesJson(FileTypesFileName)
	extensions := make(map[string]api.FileTypesJsonStruct, len(typesJson))
	for _, jsonStruct := range typesJson {
		for _, extension := range jsonStruct.Extensions {
			extensions[extension] = jsonStruct
		}
	}
	return extensions
}
