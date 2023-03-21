package editor

import (
	"github.com/labstack/gommon/log"
	"simple_text_editor/core/api"
	"simple_text_editor/core/implementation/utils"
)

type AppStruct struct {
	retrieveContext api.ContextRetriever
	dialogs         api.DialogsApi
	openedFiles     map[int64]*api.OpenedFile
}

func (receiver *AppStruct) GetFilesMap() *map[int64]*api.OpenedFile {
	log.Info("GetFilesMap", receiver.openedFiles)
	return &receiver.openedFiles
}
func (receiver *AppStruct) InactivateAllFiles() {
	filesMap := *receiver.GetFilesMap()
	log.Info("InactivateAllFiles", filesMap)
	for _, openedFile := range filesMap {
		derefFile := *openedFile
		info := *derefFile.GetInformation()
		info.SetIsOpened(false)
	}
}
func (receiver *AppStruct) ChangeFileStatusToOpened(uniqueIdentifier int64) {
	files := *receiver.GetFilesMap()
	log.Info("ChangeFileStatusToOpened", files)
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
func (receiver *AppStruct) FindOpenedFile() api.OpenedFile {
	files := *receiver.GetFilesMap()
	log.Info("FindOpenedFile", files)

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
func (receiver *AppStruct) CreateEmptyFileAndMakeItOpened() {
	log.Info("CreateEmptyFileAndMakeItOpened")
	emptyFile := utils.CreateEmptyFile()
	receiver.AddFileToMemory(&emptyFile)
	info := *emptyFile.GetInformation()
	receiver.ChangeFileStatusToOpened(info.GetOpenTimeStamp())
}
func (receiver *AppStruct) AddFileToMemory(openedFile *api.OpenedFile) *api.OpenedFile {
	log.Info("AddFileToMemory", openedFile)
	if openedFile == nil {
		log.Warn("Opened file is NIL, nil will be returned")
		return nil
	}
	filesMap := *receiver.GetFilesMap()
	information := *(*openedFile).GetInformation()
	timeStamp := information.GetOpenTimeStamp()
	filesMap[timeStamp] = openedFile
	log.Info("AddFileToMemory", openedFile)
	return openedFile
}
func (receiver *AppStruct) IsFileAlreadyOpened(filePath string) bool {
	log.Info("IsFileAlreadyOpened", filePath)
	if len(filePath) == 0 {
		log.Info("IsFileAlreadyOpened", false)
		return false
	}
	filesMap := *receiver.GetFilesMap()
	for _, openedFile := range filesMap {
		derefFile := *openedFile
		info := *derefFile.GetInformation()
		if filePath == info.GetPath() {
			log.Info("IsFileAlreadyOpened", true)
			return true
		}
	}
	log.Info("IsFileAlreadyOpened", false)
	return false
}
func (receiver *AppStruct) CloseFile(uniqueIdentifier int64) bool {
	log.Info("CloseFile", uniqueIdentifier)
	filesMap := *receiver.GetFilesMap()
	_, ok := filesMap[uniqueIdentifier]
	if !ok {
		log.Info("CloseFile", uniqueIdentifier, false)
		return false
	}
	delete(filesMap, uniqueIdentifier)
	log.Info("CloseFile", uniqueIdentifier, true)
	return true
}
func (receiver *AppStruct) FindAnyFileInMemory() *api.OpenedFile {
	log.Info("FindAnyFileInMemory")
	filesMap := *receiver.GetFilesMap()
	if len(filesMap) == 0 {
		log.Info("FindAnyFileInMemory, no files in memory, new will be created")
		newEmptyFileRef := utils.CreateEmptyFile()
		receiver.AddFileToMemory(&newEmptyFileRef)
	}
	var anyFile *api.OpenedFile
	for _, fileRef := range filesMap {
		anyFile = fileRef
	}
	log.Info("FindAnyFileInMemory", anyFile)
	return anyFile
}
func CreateEditorApplication(retrieveContext *api.ContextRetriever, dialogs *api.DialogsApi) api.EditorApplication {
	log.Info("CreateEditorApplication")
	return &AppStruct{
		retrieveContext: *retrieveContext,
		dialogs:         *dialogs,
		openedFiles:     make(map[int64]*api.OpenedFile),
	}
}
