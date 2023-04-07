package main

import (
	"embed"
	"encoding/json"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"simple_text_editor/core/v3/components/application"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/utils"
	"simple_text_editor/core/v3/validators"
)

//go:embed all:frontend/dist
var assets embed.FS

//const FileTypesFileName = "frontend/public/fileTypes.json" //TODO: remove

func main() {
	newDefaultLogger := logger.NewDefaultLogger()

	byteValue := []byte(ConfigJson)
	var typesJsons []types.FileTypesJsonStruct

	unmarshallErr := json.Unmarshal(byteValue, &typesJsons)
	if validators.HasError(unmarshallErr) {
		newDefaultLogger.Fatal(unmarshallErr.Error())
		return
	}

	typesMap, mapErr := utils.MapFileTypesJsonStructToTypesMap(typesJsons)
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
			OpenInspectorOnStartup: false,
		},
		Logger: newDefaultLogger,
	})

	if appErr != nil {
		newDefaultLogger.Fatal(appErr.Error())
	}
}
