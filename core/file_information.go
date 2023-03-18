package core

import (
	"path/filepath"
	"simple_text_editor/core/constants"
	"strings"
	"time"
)

type FileInformation struct {
	OpenTimeStamp int64  // OpenTimeStamp - Used to place tab in right order
	Path          string // Path - Full path to the file (empty for new file)
	Name          string // Name of the file (last item from the path without extension)
	Ext           string // Ext - Extension of the file (last item after . in path if available) (empty for new file)
	Type          string // Type based on the extension. Can be the same or different
	Exists        bool   // Exists equal true if this file was opened and not just created
	IsOpenedNow   bool   // IsOpenedNow equal true if this file should be shown on the UI now
	IsChanged     bool   //IsChanged equal true if this the actual content is not equal to original content
}

func CreateFileInformationStruct(path string) FileInformation {
	timestamp := time.Now().UnixNano()
	fileName := getFileNameFromPath(path)
	fileExtension := getFileExtensionFromPath(path)
	fileType := getFileType(fileExtension)
	exists := len(path) > 0

	return FileInformation{
		OpenTimeStamp: timestamp,
		Path:          path,
		Name:          fileName,
		Ext:           fileExtension,
		Type:          fileType,
		Exists:        exists,
		IsOpenedNow:   false,
		IsChanged:     false,
	}
}

func (receiver *FileInformation) setPath(path string) {
	fileName := getFileNameFromPath(path)
	fileExtension := getFileExtensionFromPath(path)
	fileType := getFileType(fileExtension)
	exists := len(path) > 0

	receiver.Path = path
	receiver.Name = fileName
	receiver.Ext = fileExtension
	receiver.Type = fileType
	receiver.Exists = exists
}

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
