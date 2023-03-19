package file

import (
	"simple_text_editor/core/api"
	"simple_text_editor/core/structures/information"
)

type FileStruct struct {
	FileInfo        api.InformationApi // FileInfo - contains base file information
	OriginalContent string             // OriginalContent - Content that was read after open
	ActualContent   string             // ActualContent - Content that can be changed during the time
}

func (receiver *FileStruct) GetInformationRef() *api.InformationApi {
	return &receiver.FileInfo
}

func (receiver *FileStruct) GetOriginalContent() string {
	return receiver.OriginalContent
}

func (receiver *FileStruct) GetActualContent() string {
	return receiver.ActualContent
}

func (receiver *FileStruct) HasChanges() bool {
	return receiver.OriginalContent == receiver.ActualContent
}

func (receiver *FileStruct) SetOriginalContent(content string) {
	receiver.OriginalContent = content
}

func (receiver *FileStruct) SetActualContent(content string) {
	receiver.ActualContent = content
}

func CreateFile(descriptor *api.InformationApi, content string) api.FileApi {
	var fileToReturn api.FileApi
	fileToReturn = &FileStruct{
		FileInfo:        *descriptor,
		OriginalContent: content,
		ActualContent:   content,
	}
	return fileToReturn
}

func CreateEmptyFile() api.FileApi {
	var fileToReturn api.FileApi
	fileInfo := information.CreateInformationFromPath("")
	fileToReturn = CreateFile(fileInfo, "")
	return fileToReturn
}

func CreateExistingFile(filePath string, fileContent string) api.FileApi {
	var fileToReturn api.FileApi
	fileInfo := information.CreateInformationFromPath(filePath)
	fileToReturn = CreateFile(fileInfo, fileContent)
	return fileToReturn
}
