package utils

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io"
	"os"
	"simple_text_editor/core/v2/api"
)

func LoadFileTypesJson(path string) []api.FileTypesJsonStruct {
	if len(path) == 0 {
		return []api.FileTypesJsonStruct{}
	}

	jsonFile, openFileErr := os.Open(path)

	if openFileErr != nil {
		log.Error("Error in opening JSON config file: ", openFileErr)
		return []api.FileTypesJsonStruct{}
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Error("Error in closing file: ", err)
		}
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)
	var typesJsons []api.FileTypesJsonStruct

	unmarshalErr := json.Unmarshal(byteValue, &typesJsons)
	if unmarshalErr != nil {
		log.Error("Error converting JSON to struct:", unmarshalErr)
		return []api.FileTypesJsonStruct{}
	}

	return typesJsons
}

func GetSupportedFileFilters(extensions map[string]api.FileTypesJsonStruct) []runtime.FileFilter {
	log.Info("GetSupportedFileFilters")

	fileFilters := make([]runtime.FileFilter, 0, len(extensions))
	for _, value := range extensions {
		fileFilters = append(fileFilters, createFileFilter(&value))
	}
	fileFilters = append(fileFilters, runtime.FileFilter{
		DisplayName: "Plain Text",
		Pattern:     "*.txt",
	})
	fileFilters = append(fileFilters, runtime.FileFilter{
		DisplayName: "Any File",
		Pattern:     "",
	})
	log.Info("GetSupportedFileFilters, return", fileFilters)
	return fileFilters
}

func createFileFilter(fileTypeInfo *api.FileTypesJsonStruct) runtime.FileFilter {
	log.Info("createFileFilter", *fileTypeInfo)
	typePattern := "*.%s;"
	pattern := ""
	for _, value := range fileTypeInfo.Extensions {
		pattern += fmt.Sprintf(typePattern, value)
	}
	log.Info("createFileFilter", pattern)
	return runtime.FileFilter{
		DisplayName: fileTypeInfo.Name,
		Pattern:     pattern,
	}
}
