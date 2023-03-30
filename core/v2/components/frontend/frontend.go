package frontend

import (
	"fmt"
	"simple_text_editor/core/v2/api"
	"sort"
)

type frontStruct struct {
	app         api.IEditor
	events      api.IEvents
	dialogs     api.IDialogs
	typeManager api.ITypeManager
}

func (r *frontStruct) GetFilesShortInfo() []api.FileInfoStruct {
	filesInfo := r.app.GetAllFilesInfo()
	sort.Slice(filesInfo, func(i, j int) bool {
		return filesInfo[i].Id < filesInfo[j].Id
	})
	return filesInfo
}
func (r *frontStruct) GetOpenedFile() api.FileStruct {
	file, err := r.app.GetOpenedFile()
	if err != nil {
		r.events.SendErrorEvent("Opened file is not found", err)
		return api.FileStruct{}
	}
	return *file
}
func (r *frontStruct) SwitchOpenedFileTo(fileId int64) {
	err := r.app.SwitchOpenedFileTo(fileId)
	if err != nil {
		r.events.SendErrorEvent("Can't switch to file", err)
		return
	}
	r.events.SendEvent(api.EventOnFileIsSwitched)
}
func (r *frontStruct) UpdateFileContent(fileId int64, content string) {
	err := r.app.UpdateFileContent(fileId, content)
	if err != nil {
		r.events.SendErrorEvent("File content was not updated", err)
		return
	}
	r.events.SendEvent(api.EventOnFileContentIsUpdated)
}
func (r *frontStruct) UpdateFileInformation(fileId int64, information api.FileInfoUpdateStruct) {
	err := r.app.UpdateFileInformation(fileId, information)
	if err != nil {
		r.events.SendWarnEvent("File information was not updated", err)
		return
	}

	r.events.SendEvent(api.EventOnFileInformationUpdated)
}
func (r *frontStruct) GetFileTypes() []api.KeyValuePairStruct {
	return r.typeManager.BuildFileTypeMappingKeyToName()
}
func (r *frontStruct) GetFileTypeExtension(fileTypeKey string) []api.KeyValuePairStruct {
	return r.typeManager.BuildFileTypeMappingExtToExt(fileTypeKey)
}

func (r *frontStruct) CloseCurrentFile(fileId int64) {
	file, getFileErr := r.app.GetFileById(fileId)
	if getFileErr != nil {
		r.events.SendErrorEvent("Failed to access file by id in memory", getFileErr)
		return
	}

	if !file.Changed {
		closeErr := r.app.CloseFile(file.Id)
		if closeErr != nil {
			r.events.SendErrorEvent("Failed to close file", closeErr)
			return
		}
	} else {
		var title = fmt.Sprintf("Close file? (%s)", file.Name)
		const message = "File has changes. Do you want proceed and close file? (Changes will not be saved)"

		btn, dialErr := r.dialogs.OkCancelMessageDialog(title, message)
		if dialErr != nil {
			r.events.SendErrorEvent("Failed to process Message dialog", dialErr)
			return
		}

		if btn.EqualTo(api.BtnCancel) {
			return
		}

		closeErr := r.app.CloseFile(file.Id)
		if closeErr != nil {
			r.events.SendErrorEvent("Failed to close file", closeErr)
			return
		}
	}

	r.events.SendEvent(api.EventOnFileClosed)
}

func CreateFrontendApi(editor api.IEditor, events api.IEvents, dialogs api.IDialogs, typeManager api.ITypeManager) api.IFrontendApi {
	if editor == nil || events == nil {
		panic("Required editor or events api are nil")
	}
	frontend := frontStruct{
		app:         editor,
		events:      events,
		dialogs:     dialogs,
		typeManager: typeManager,
	}
	return &frontend
}
