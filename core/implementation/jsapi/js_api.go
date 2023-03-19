package jsapi

import "simple_text_editor/core/api"

type JsApiStruct struct {
	app api.EditorApplication
}

func (r *JsApiStruct) GetFilesInformation() []api.FileInformation {
	filesMap := *r.app.GetFilesMap()

	allOpenedFiles := make([]api.FileInformation, 0, len(filesMap))

	for _, fileRef := range filesMap {
		infoRef := (*fileRef).GetInformation()
		allOpenedFiles = append(allOpenedFiles, *infoRef)
	}

	return allOpenedFiles
}
func (r *JsApiStruct) FindOpenedFile() api.OpenedFile {
	return r.app.FindOpenedFile()
}
func (r *JsApiStruct) ChangeFileStatusToOpened(uniqueIdentifier int64) {
	r.app.ChangeFileStatusToOpened(uniqueIdentifier)
}
func (r *JsApiStruct) ChangeFileContent(uniqueIdentifier int64, content string) bool {
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
	return &JsApiStruct{
		app: *app,
	}
}
