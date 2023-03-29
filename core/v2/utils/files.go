package utils

import (
	"github.com/labstack/gommon/log"
	"path/filepath"
	"simple_text_editor/core/v2/api"
	"strings"
)

func GetFileNameFromPath(filePath string) string {
	log.Info("GetFileNameFromPath", filePath)
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
func GetFileType(fileExtension string, extensions map[string]api.FileTypesJsonStruct) string {
	log.Info("getFileType", fileExtension)
	if len(fileExtension) > 0 && strings.HasPrefix(fileExtension, ".") {
		typeToSearch := fileExtension[1:] // remove . in begin. .yml => yml
		supportedFileTypes := extensions
		typeInformation, ok := supportedFileTypes[typeToSearch]
		if !ok {
			return ""
		}
		log.Info("getFileType", strings.ToLower(typeInformation.Key))
		return strings.ToLower(typeInformation.Key)
	}
	log.Info("getFileType", fileExtension)
	return fileExtension
}
