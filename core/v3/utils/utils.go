package utils

import (
	"encoding/json"
	"errors"
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/validators"
)

func ForEach[T interface{}](slice []T, callback func(index int, data *T)) {
	if validators.IsNil(slice) {
		return
	}
	for i, value := range slice {
		callback(i, &value)
	}
}

func ReadFileTypesJson(path string) ([]types.FileTypesJsonStruct, error) {
	if validators.IsEmptyString(path) {
		return nil, errors.New("path to file types config json is empty")
	}

	jsonFile, openFileErr := os.Open(path)

	if validators.HasError(openFileErr) {
		return nil, openFileErr
	}
	defer func(jsonFile *os.File) {
		closeErr := jsonFile.Close()
		if validators.HasError(closeErr) {
			log.Error("Error in closing file: ", closeErr)
		}
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)
	var typesJsons []types.FileTypesJsonStruct

	unmarshalErr := json.Unmarshal(byteValue, &typesJsons)
	if validators.HasError(unmarshalErr) {
		return nil, unmarshalErr
	}

	return typesJsons, nil
}

func MapFileTypesJsonStructToTypesMap(fileTypes []types.FileTypesJsonStruct) (types.TypesMap, error) {
	if validators.IsNilOrEmptySlice(fileTypes) {
		return nil, errors.New("fileTypes is empty slice, can't build map")
	}

	typesMap := make(types.TypesMap, len(fileTypes))
	ForEach(fileTypes, func(_ int, data *types.FileTypesJsonStruct) {
		typesMap[data.Key] = data
	})

	return typesMap, nil
}
