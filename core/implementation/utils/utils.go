package utils

import (
	"path/filepath"
	"simple_text_editor/core/api"
	"simple_text_editor/core/constants"
	"simple_text_editor/core/implementation/file"
	"simple_text_editor/core/implementation/info"
	"strings"
	"time"
)

func getFileNameFromPath(filePath string) string {
	if len(filePath) == 0 {
		return filePath
	}
	fileName := filepath.Base(filePath)
	return fileName
}
func getFileExtensionFromPath(filePath string) string {
	if len(filePath) == 0 {
		return filePath
	}
	fileName := filepath.Ext(filePath)
	return fileName
}
func getFileType(fileExtension string) string {
	if len(fileExtension) > 0 && strings.HasPrefix(fileExtension, ".") {
		typeToSearch := fileExtension[1:] // remove . in begin. .yml => yml
		supportedFileTypes := constants.GetExtToLangMapping()
		typeInformation, ok := (*supportedFileTypes)[typeToSearch]
		if !ok {
			return ""
		}
		return strings.ToLower(typeInformation)
	}
	return fileExtension
}
func createInformationFromPath(filePath string) *api.FileInformation {
	timestamp := time.Now().UnixNano()
	fileName := getFileNameFromPath(filePath)
	fileExtension := getFileExtensionFromPath(filePath)
	fileType := getFileType(fileExtension)
	exists := len(filePath) > 0

	var obj api.FileInformation
	obj = &info.FileInformationStruct{
		OpenTimeStamp:  timestamp,
		FilePath:       filePath,
		FileName:       fileName,
		FileExtension:  fileExtension,
		FileType:       fileType,
		FileExists:     exists,
		FileIsOpened:   false,
		FileHasChanges: false,
	}

	return &obj
}
func createOpenedFile(descriptor *api.FileInformation, content string) api.OpenedFile {
	var fileToReturn api.OpenedFile
	fileToReturn = &file.OpenedFileStruct{
		FileInfo:        *descriptor,
		OriginalContent: content,
		ActualContent:   content,
	}
	return fileToReturn
}

func CreateEmptyFile() api.OpenedFile {
	var fileToReturn api.OpenedFile
	fileInfo := createInformationFromPath("")
	fileToReturn = createOpenedFile(fileInfo, "")
	return fileToReturn
}

func CreateExistingFile(filePath string, fileContent string) api.OpenedFile {
	var fileToReturn api.OpenedFile
	fileInfo := createInformationFromPath(filePath)
	fileToReturn = createOpenedFile(fileInfo, fileContent)
	return fileToReturn
}
