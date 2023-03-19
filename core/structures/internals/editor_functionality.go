package internals

import (
	"simple_text_editor/core/api"
	"simple_text_editor/core/structures/file"
)

type EditorInternalApi struct {
	appFilesMap *map[int64]*api.FileApi
}

func (receiver *EditorInternalApi) getFilesMap() map[int64]*api.FileApi {
	return *receiver.appFilesMap
}
func (receiver *EditorInternalApi) CreateEmptyFile() *api.FileApi {
	emptyFile := file.CreateEmptyFile()
	return &emptyFile
}
func (receiver *EditorInternalApi) CreateExistingFile(filePath string, fileContent string) *api.FileApi {
	existingFile := file.CreateExistingFile(filePath, fileContent)
	return &existingFile
}

func (receiver *EditorInternalApi) AddEmptyFile(file *api.FileApi) *api.FileApi {
	emptyFile := *file

	filesMap := receiver.getFilesMap()
	filesMap[(*emptyFile.GetInformationRef()).GetOpenTimeStamp()] = &emptyFile

	return file

}
func (receiver *EditorInternalApi) AddExistingFile(existingFile *api.FileApi) *api.FileApi {
	if existingFile == nil {
		// TODO: process error
		return nil
	}
	filesMap := receiver.getFilesMap()
	filesMap[(*(*existingFile).GetInformationRef()).GetOpenTimeStamp()] = existingFile
	return existingFile
}
func (receiver *EditorInternalApi) InactivateAllFiles() {
	filesMap := receiver.getFilesMap()
	for _, openedFile := range filesMap {
		derefFile := *openedFile
		info := *derefFile.GetInformationRef()
		info.SetIsOpenedNow(false)
	}
}
func (receiver *EditorInternalApi) IsFileAlreadyOpened(filePath string) bool {
	if len(filePath) == 0 {
		return false
	}
	filesMap := receiver.getFilesMap()
	for _, openedFile := range filesMap {
		derefFile := *openedFile
		info := *derefFile.GetInformationRef()
		if filePath == info.GetPath() {
			return true
		}
	}
	return false
}
func (receiver *EditorInternalApi) GetFilesInformation() []api.InformationApi {
	filesMap := receiver.getFilesMap()
	allOpenedFiles := make([]api.InformationApi, 0, len(filesMap))

	for _, fileRef := range filesMap {
		info := (*fileRef).GetInformationRef()
		allOpenedFiles = append(allOpenedFiles, *info)
	}

	return allOpenedFiles
}
func (receiver *EditorInternalApi) FindOpenedFile() api.FileApi {
	filesMap := receiver.getFilesMap()
	for _, fileRef := range filesMap {
		info := (*fileRef).GetInformationRef()
		if (*info).GetIsOpenedNow() {
			return *fileRef
		}
	}
	return nil
}
func (receiver *EditorInternalApi) ChangeFileStatusToOpened(uniqueIdentifier int64) {
	filesMap := receiver.getFilesMap()
	fileToActivate, ok := filesMap[uniqueIdentifier]
	if !ok {
		return
	}
	receiver.InactivateAllFiles()
	foundFile := *fileToActivate
	info := *foundFile.GetInformationRef()
	info.SetIsOpenedNow(true)
}
func (receiver *EditorInternalApi) ChangeFileContent(uniqueIdentifier int64, content string) bool {
	filesMap := receiver.getFilesMap()
	fileToUpdate, ok := filesMap[uniqueIdentifier]
	if !ok {
		return false
	}
	foundFile := *fileToUpdate
	foundFile.SetActualContent(content)
	return foundFile.HasChanges()
}

func (receiver *EditorInternalApi) CloseFile(uniqueIdentifier int64) bool {
	filesMap := receiver.getFilesMap()
	_, ok := filesMap[uniqueIdentifier]
	if !ok {
		return false
	}
	delete(filesMap, uniqueIdentifier)
	return true
}

func (receiver *EditorInternalApi) FindAnyFileInMemory() *api.FileApi {
	filesMap := receiver.getFilesMap()
	if len(filesMap) == 0 {
		newEmptyFileRef := receiver.CreateEmptyFile()
		receiver.AddEmptyFile(newEmptyFileRef)
	}
	var anyFile *api.FileApi
	for _, fileRef := range filesMap {
		anyFile = fileRef
	}
	return anyFile
}

func CreateEditorInternals(appFilesMap *map[int64]*api.FileApi) api.EditorApplicationApi {
	return &EditorInternalApi{
		appFilesMap: appFilesMap,
	}
}
