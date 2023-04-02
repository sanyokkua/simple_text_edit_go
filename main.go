package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	//"simple_text_editor/core/v2"
	//"simple_text_editor/core/v2/utils"
	"simple_text_editor/core/v3/factories/application"
	utils2 "simple_text_editor/core/v3/utils"
	"simple_text_editor/core/v3/validators"
)

//go:embed all:frontend/dist
var assets embed.FS

const FileTypesFileName = "fileTypes.json"

func main() {
	newDefaultLogger := logger.NewDefaultLogger()

	fileTypesJson, readErr := utils2.ReadFileTypesJson(FileTypesFileName)
	if validators.HasError(readErr) {
		panic("Failed to read config JSON")
	}
	typesMap, mapErr := utils2.MapFileTypesJsonStructToTypesMap(fileTypesJson)
	if validators.HasError(mapErr) {
		panic("Failed to map fileTypes to FileTypeMap object")
	}

	iApplication := application.CreateIApplication(typesMap)
	iFrontApi := iApplication.GetFrontendApi()
	iMenuHelper := iApplication.GetMenuApi()

	appErr := wails.Run(&options.App{
		Title:  "Simple Text Editor",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:     iApplication.OnStartup,
		OnDomReady:    iApplication.OnDomReady,
		OnShutdown:    iApplication.OnShutdown,
		OnBeforeClose: iApplication.OnBeforeClose,
		Menu:          iMenuHelper.CreateMenu(),
		Bind: []interface{}{
			iFrontApi,
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
