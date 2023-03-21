package utils

import (
	"github.com/labstack/gommon/log"
	"path/filepath"
	"simple_text_editor/core/api"
	"simple_text_editor/core/constants"
	"simple_text_editor/core/implementation/file"
	"simple_text_editor/core/implementation/info"
	"strings"
	"time"
)

func getFileNameFromPath(filePath string) string {
	log.Info("getFileNameFromPath", filePath)
	if len(filePath) == 0 {
		return filePath
	}
	fileName := filepath.Base(filePath)
	log.Info("getFileNameFromPath: return", fileName)
	return fileName
}
func GetFileExtensionFromPath(filePath string) string {
	log.Info("GetFileExtensionFromPath", filePath)
	if len(filePath) == 0 {
		return filePath
	}
	fileName := filepath.Ext(filePath)
	log.Info("GetFileExtensionFromPath, return", fileName)
	return fileName
}
func getFileType(fileExtension string) string {
	log.Info("getFileType", fileExtension)
	if len(fileExtension) > 0 && strings.HasPrefix(fileExtension, ".") {
		typeToSearch := fileExtension[1:] // remove . in begin. .yml => yml
		supportedFileTypes := constants.GetExtToLangMapping()
		typeInformation, ok := (*supportedFileTypes)[typeToSearch]
		if !ok {
			return ""
		}
		log.Info("getFileType", strings.ToLower(typeInformation))
		return strings.ToLower(typeInformation)
	}
	log.Info("getFileType", fileExtension)
	return fileExtension
}
func CreateInformationFromPath(filePath string) *api.FileInformation {
	log.Info("CreateInformationFromPath", filePath)
	timestamp := time.Now().UnixNano()
	fileName := getFileNameFromPath(filePath)
	fileExtension := GetFileExtensionFromPath(filePath)
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
	log.Info("CreateInformationFromPath, return", obj)
	return &obj
}
func createOpenedFile(descriptor *api.FileInformation, content string) api.OpenedFile {
	log.Info("createOpenedFile", *descriptor, content)
	var fileToReturn api.OpenedFile
	fileToReturn = &file.OpenedFileStruct{
		FileInfo:        descriptor,
		OriginalContent: content,
		ActualContent:   content,
	}
	log.Info("createOpenedFile, return", fileToReturn)
	return fileToReturn
}

func CreateEmptyFile() api.OpenedFile {
	log.Info("CreateEmptyFile")
	var fileToReturn api.OpenedFile
	fileInfo := CreateInformationFromPath("")
	fileToReturn = createOpenedFile(fileInfo, "")
	log.Info("CreateEmptyFile, return", fileToReturn)
	return fileToReturn
}

func CreateExistingFile(filePath string, fileContent string) api.OpenedFile {
	log.Info("CreateExistingFile", filePath, fileContent)
	var fileToReturn api.OpenedFile
	fileInfo := CreateInformationFromPath(filePath)
	fileToReturn = createOpenedFile(fileInfo, fileContent)
	log.Info("CreateExistingFile, return", fileToReturn)
	return fileToReturn
}
