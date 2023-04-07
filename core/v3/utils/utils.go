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
	if slice == nil {
		return
	}
	for i, value := range slice {
		currVal := value
		callback(i, &currVal)
	}
}

func FindInSlice[T interface{}](slice []T, compare func(value T) bool) (index int) {
	if slice == nil {
		return -1
	}

	for i, value := range slice {
		if compare(value) {
			return i
		}
	}
	return -1
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

	for _, fileType := range fileTypes {
		fileType := fileType
		typesMap[fileType.Key] = &fileType
	}

	return typesMap, nil
}
