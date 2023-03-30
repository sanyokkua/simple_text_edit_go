package utils

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
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
