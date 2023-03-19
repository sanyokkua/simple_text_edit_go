package editor

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/api"
)

type EditorStruct struct {
	contextRetriever api.GetContext
	appContext       context.Context        // Application context
	dialogsApi       api.DialogsApi         // Dialogs Api
	openedFiles      map[int64]*api.FileApi // Files that currently opened in the app memory
}

func (r *EditorStruct) OnStartup(ctx context.Context) {

}

func (r *EditorStruct) OnDomReady(ctx context.Context) {

}

func (r *EditorStruct) OnShutdown(ctx context.Context) {

}

func (r *EditorStruct) OnBeforeClose(ctx context.Context) (prevent bool) {
	var hasChanges bool
	for _, file := range r.GetFilesMap() {
		if (*file).HasChanges() {
			hasChanges = true
			break
		}
	}

	if !hasChanges {
		runtime.Quit(*r.GetContext())
	}

	dialogResult, err := r.dialogsApi.OkCancelMessageDialog(
		"flaskApplicationApi has files with changes",
		"You are trying to close application that have files with changes, continue?",
	)
	if err != nil {
		return
	}

	if dialogResult == "Ok" {
		runtime.Quit(*r.GetContext())
	}

	return false
}

func (r *EditorStruct) GetContext() *context.Context {
	ctx := r.contextRetriever()
	return &ctx
}

func (r *EditorStruct) GetFilesMap() map[int64]*api.FileApi {
	return r.openedFiles
}

func (r *EditorStruct) inactivateAllFiles() {
	files := r.GetFilesMap()
	for _, openedFileRef := range files {
		file := *openedFileRef
		informationRef := file.GetInformationRef()
		info := *informationRef
		info.SetIsChanged(false)
	}
}
func (r *EditorStruct) ActivateFile(uniqueIdentifier int64) {
	files := r.GetFilesMap()
	fileRef, ok := files[uniqueIdentifier]
	if !ok {
		return
	}
	r.inactivateAllFiles()
	file := *fileRef
	infoRef := file.GetInformationRef()
	info := *infoRef
	info.SetIsOpenedNow(true)
}
func (r *EditorStruct) GetActiveFile() api.FileApi {
	files := r.GetFilesMap()

	for _, fileRef := range files {
		file := *fileRef
		infoRef := file.GetInformationRef()

		info := *infoRef
		if info.GetIsOpenedNow() {
			return file
		}
	}
	return nil
}
func (r *EditorStruct) GetAllFilesInformation() []api.InformationApi {
	filesMap := r.GetFilesMap()
	allOpenedFiles := make([]api.InformationApi, 0, len(filesMap))

	for _, fileRef := range filesMap {
		infoRef := (*fileRef).GetInformationRef()
		allOpenedFiles = append(allOpenedFiles, *infoRef)
	}
	return allOpenedFiles
}
func (r *EditorStruct) UpdateActiveFileContent(content string) bool {
	files := r.GetFilesMap()

	for _, fileRef := range files {
		file := *fileRef
		infoRef := file.GetInformationRef()

		info := *infoRef
		if info.GetIsOpenedNow() {
			file.SetActualContent(content)
			return file.HasChanges()
		}
	}

	return false
}

func CreateFlaskApplicationApi(dialogsApi *api.DialogsApi) *api.FlaskApplicationApi {
	var editorToReturn api.FlaskApplicationApi
	editorToReturn = &EditorStruct{
		openedFiles: make(map[int64]*api.FileApi),
		dialogsApi:  *dialogsApi,
	}
	return &editorToReturn
}
