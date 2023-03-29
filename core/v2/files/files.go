package files

import (
	"simple_text_editor/core/v2/api"
	"simple_text_editor/core/v2/components/id"
	"simple_text_editor/core/v2/utils"
)

var idProvider = id.CreateIdProvider()

func CreateNewFileEmpty() api.FileStruct {
	return api.FileStruct{
		Id:   idProvider.GetId(),
		Name: "New",
		New:  true,
	}
}

func CreateNewFileWithData(path string, originalContent string, extensions map[string]api.FileTypesJsonStruct) api.FileStruct {
	fileName := utils.GetFileNameFromPath(path)
	fileExtension := utils.GetFileExtensionFromPath(path)
	fileType := utils.GetFileType(fileExtension, extensions)
	isNew := len(path) == 0

	return api.FileStruct{
		Id:             idProvider.GetId(),
		Path:           path,
		Name:           fileName,
		Extension:      fileExtension,
		Type:           fileType,
		InitialContent: originalContent,
		ActualContent:  originalContent,
		New:            isNew,
		Opened:         false,
		Changed:        false,
	}
}
