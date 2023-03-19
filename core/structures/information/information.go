package information

import (
	"fmt"
	"path/filepath"
	"simple_text_editor/core/api"
	"simple_text_editor/core/constants"
	"strings"
	"time"
)

type InformationStruct struct {
	OpenTimeStamp int64  // OpenTimeStamp - Used to place tab in right order
	FilePath      string // FilePath - Full path to the file (empty for new file)
	FileName      string // FileName of the file (last item from the path without extension)
	FileExtension string // FileExtension - Extension of the file (last item after . in path if available) (empty for new file)
	FileType      string // FileType based on the extension. Can be the same or different
	FileExists    bool   // FileExists equal true if this file was opened and not just created
	FileOpened    bool   // FileOpened equal true if this file should be shown on the UI now
	HasChanges    bool   // HasChanges equal true if this the actual content is not equal to original content
}

func (receiver *InformationStruct) GetOpenTimeStamp() int64 {
	return receiver.OpenTimeStamp
}
func (receiver *InformationStruct) GetPath() string {
	return receiver.FilePath
}
func (receiver *InformationStruct) GetName() string {
	return receiver.FileName
}
func (receiver *InformationStruct) GetExt() string {
	return receiver.FileExtension
}
func (receiver *InformationStruct) GetType() string {
	return receiver.FileType
}
func (receiver *InformationStruct) GetExists() bool {
	return receiver.FileExists
}
func (receiver *InformationStruct) GetIsOpenedNow() bool {
	return receiver.FileOpened
}
func (receiver *InformationStruct) GetIsChanged() bool {
	return receiver.HasChanges
}
func (receiver *InformationStruct) SetOpenTimeStamp(timestamp int64) {
	receiver.OpenTimeStamp = timestamp
}
func (receiver *InformationStruct) SetPath(path string) {
	fileName := getFileNameFromPath(path)
	fileExtension := getFileExtensionFromPath(path)
	fileType := getFileType(fileExtension)
	exists := len(path) > 0

	receiver.FilePath = path
	receiver.FileName = fileName
	receiver.FileExtension = fileExtension
	receiver.FileType = fileType
	receiver.FileExists = exists
}
func (receiver *InformationStruct) SetName(name string) {
	if len(receiver.FileName) > 0 && receiver.FileName != name {
		receiver.FileName = name
		receiver.FileExists = false
	} else {
		receiver.FileName = name
	}

}
func (receiver *InformationStruct) SetExt(extension string) {
	fileType := getFileType(extension)

	receiver.FileExtension = extension
	receiver.FileType = fileType
}
func (receiver *InformationStruct) SetType(typeVal string) {
	receiver.FileExtension = fmt.Sprintf(".%s", typeVal)
	receiver.FileType = typeVal
}
func (receiver *InformationStruct) SetExists(exists bool) {
	receiver.FileExists = exists
}
func (receiver *InformationStruct) SetIsOpenedNow(opened bool) {
	receiver.FileOpened = opened
}
func (receiver *InformationStruct) SetIsChanged(changed bool) {
	receiver.HasChanges = changed
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

func CreateInformationFromPath(filePath string) *api.InformationApi {
	timestamp := time.Now().UnixNano()
	fileName := getFileNameFromPath(filePath)
	fileExtension := getFileExtensionFromPath(filePath)
	fileType := getFileType(fileExtension)
	exists := len(filePath) > 0

	var obj api.InformationApi
	obj = &InformationStruct{
		OpenTimeStamp: timestamp,
		FilePath:      filePath,
		FileName:      fileName,
		FileExtension: fileExtension,
		FileType:      fileType,
		FileExists:    exists,
		FileOpened:    false,
		HasChanges:    false,
	}

	return &obj
}
