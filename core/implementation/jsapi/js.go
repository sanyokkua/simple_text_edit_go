package jsapi

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/api"
	"simple_text_editor/core/apperrors"
	"simple_text_editor/core/constants"
)

type JsStruct struct {
	app       api.EditorApplication
	retriever *api.ContextRetriever
}

func (r *JsStruct) getContext() context.Context {
	return (*r.retriever)()
}

func (r *JsStruct) GetFilesInformation() []api.FileInformation {
	log.Info("GetFilesInformation")
	filesMap := *r.app.GetFilesMap()

	allOpenedFiles := make([]api.FileInformation, 0, len(filesMap))

	for _, fileRef := range filesMap {
		infoRef := (*fileRef).GetInformation()
		allOpenedFiles = append(allOpenedFiles, *infoRef)
	}
	log.Info("GetFilesInformation, return", allOpenedFiles)
	return allOpenedFiles
}
func (r *JsStruct) FindOpenedFile() api.OpenedFile {
	log.Info("FindOpenedFile")
	return r.app.FindOpenedFile()
}
func (r *JsStruct) ChangeFileStatusToOpened(uniqueIdentifier int64) {
	log.Info("ChangeFileStatusToOpened", uniqueIdentifier)
	r.app.ChangeFileStatusToOpened(uniqueIdentifier)
}
func (r *JsStruct) ChangeFileContent(uniqueIdentifier int64, content string) bool {
	log.Info("ChangeFileContent", uniqueIdentifier, content)
	files := *r.app.GetFilesMap()

	for _, fileRef := range files {
		file := *fileRef
		infoRef := file.GetInformation()

		info := *infoRef
		if info.GetOpenTimeStamp() == uniqueIdentifier && info.IsOpened() { // Just to be sure that all in sync
			file.SetActualContent(content)
			return file.HasChanges()
		}
	}

	return false
}

func (r *JsStruct) GetFileTypeInformation() []constants.FileTypeInformation {
	return constants.GetFileTypeInformation()
}

func (r *JsStruct) ChangeFileInformation(dialRes api.DialogResult) {
	openedFile := r.app.FindOpenedFile()
	if openedFile == nil {
		apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
			AppContext:  r.getContext(),
			Destination: constants.EventOnErrorHappened,
			Message:     "Opened file is not found",
		})
		return
	}

	if len(dialRes.FileType) == 0 && len(dialRes.FileExt) > 0 || len(dialRes.FileType) > 0 && len(dialRes.FileExt) == 0 {
		apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
			AppContext:  r.getContext(),
			Destination: constants.EventOnErrorHappened,
			Message:     "File type set and extension is not or vice-versa",
		})
		return
	}

	information := *openedFile.GetInformation()
	if len(dialRes.FileName) > 0 {
		information.SetName(dialRes.FileName)
		information.SetPath("")
		information.SetExists(false)
	}
	if len(dialRes.FileType) > 0 && len(dialRes.FileExt) > 0 {
		if dialRes.FileExt != information.GetExt() {
			information.SetExists(false)
		}
		information.SetType(dialRes.FileType)
		information.SetExt(dialRes.FileExt)
	}

	runtime.EventsEmit(r.getContext(), constants.EventOnFileInformationUpdated, "Information updated")
}

func CreateJsApi(app *api.EditorApplication, retriever *api.ContextRetriever) api.JsApi {
	log.Info("CreateJsApi", *app)
	return &JsStruct{
		app:       *app,
		retriever: retriever,
	}
}
