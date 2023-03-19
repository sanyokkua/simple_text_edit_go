package editor

import (
	"simple_text_editor/core/api"
	"simple_text_editor/core/implementation/utils"
)

type EditorApplicationStruct struct {
	retrieveContext api.ContextRetriever
	dialogs         api.DialogsApi
	openedFiles     map[int64]*api.OpenedFile
}

func (receiver *EditorApplicationStruct) GetFilesMap() *map[int64]*api.OpenedFile {
	return &receiver.openedFiles
}
func (receiver *EditorApplicationStruct) InactivateAllFiles() {
	filesMap := *receiver.GetFilesMap()
	for _, openedFile := range filesMap {
		derefFile := *openedFile
		info := *derefFile.GetInformation()
		info.SetIsOpened(false)
	}
}
func (receiver *EditorApplicationStruct) ChangeFileStatusToOpened(uniqueIdentifier int64) {
	files := *receiver.GetFilesMap()
	fileRef, ok := files[uniqueIdentifier]
	if !ok {
		return
	}
	receiver.InactivateAllFiles()
	file := *fileRef
	infoRef := file.GetInformation()
	info := *infoRef
	info.SetIsOpened(true)
}
func (receiver *EditorApplicationStruct) FindOpenedFile() api.OpenedFile {
	files := *receiver.GetFilesMap()

	for _, fileRef := range files {
		file := *fileRef
		infoRef := file.GetInformation()

		info := *infoRef
		if info.IsOpened() {
			return file
		}
	}
	return nil
}
func (receiver *EditorApplicationStruct) CreateEmptyFileAndMakeItOpened() {
	emptyFile := utils.CreateEmptyFile()
	receiver.AddFileToMemory(&emptyFile)
	info := *emptyFile.GetInformation()
	receiver.ChangeFileStatusToOpened(info.GetOpenTimeStamp())
}
func (receiver *EditorApplicationStruct) AddFileToMemory(openedFile *api.OpenedFile) *api.OpenedFile {
	if openedFile == nil {
		// TODO: process error
		return nil
	}
	filesMap := *receiver.GetFilesMap()
	information := *(*openedFile).GetInformation()
	timeStamp := information.GetOpenTimeStamp()
	filesMap[timeStamp] = openedFile
	return openedFile
}
func (receiver *EditorApplicationStruct) IsFileAlreadyOpened(filePath string) bool {
	if len(filePath) == 0 {
		return false
	}
	filesMap := *receiver.GetFilesMap()
	for _, openedFile := range filesMap {
		derefFile := *openedFile
		info := *derefFile.GetInformation()
		if filePath == info.GetPath() {
			return true
		}
	}
	return false
}
func (receiver *EditorApplicationStruct) CloseFile(uniqueIdentifier int64) bool {
	filesMap := *receiver.GetFilesMap()
	_, ok := filesMap[uniqueIdentifier]
	if !ok {
		return false
	}
	delete(filesMap, uniqueIdentifier)
	return true
}
func (receiver *EditorApplicationStruct) FindAnyFileInMemory() *api.OpenedFile {
	filesMap := *receiver.GetFilesMap()
	if len(filesMap) == 0 {
		newEmptyFileRef := utils.CreateEmptyFile()
		receiver.AddFileToMemory(&newEmptyFileRef)
	}
	var anyFile *api.OpenedFile
	for _, fileRef := range filesMap {
		anyFile = fileRef
	}
	return anyFile
}
func CreateEditorApplication(retrieveContext *api.ContextRetriever, dialogs *api.DialogsApi) api.EditorApplication {
	return &EditorApplicationStruct{
		retrieveContext: *retrieveContext,
		dialogs:         *dialogs,
		openedFiles:     make(map[int64]*api.OpenedFile),
	}
}
