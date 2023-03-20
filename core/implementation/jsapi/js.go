package jsapi

import (
	"github.com/labstack/gommon/log"
	"simple_text_editor/core/api"
)

type JsStruct struct {
	app api.EditorApplication
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

func CreateJsApi(app *api.EditorApplication) api.JsApi {
	log.Info("CreateJsApi", *app)
	return &JsStruct{
		app: *app,
	}
}
