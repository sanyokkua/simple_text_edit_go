package ops

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/v2/api"
)

type filesOperations struct {
	getContext api.ContextProvider
	events     api.IEvents
	dialogs    api.IDialogs
	app        api.IEditor
}

func (r *filesOperations) CreateNewFile() {
	err := r.app.CreateNewFileInEditor()
	if err != nil {
		r.events.SendErrorEvent("IApplication failed to create new file", err)
		return
	}

	r.events.SendEvent(api.EventOnNewFileCreated)
}
func (r *filesOperations) OpenFile() {
	filePath, dialErr := r.dialogs.OpenFileDialog()
	if dialErr != nil {
		r.events.SendErrorEvent("Failed to process Open File Dialog", dialErr)
		return
	}

	openFileErr := r.app.OpenFile(filePath)
	if openFileErr != nil {
		r.events.SendErrorEvent("Failed to Open File", openFileErr)
		return
	}

	r.events.SendEvent(api.EventOnFileOpened)
}
func (r *filesOperations) SaveCurrentFile() {
	file, getFileErr := r.app.GetOpenedFile()
	if getFileErr != nil {
		r.events.SendErrorEvent("Failed to access current file in memory", getFileErr)
		return
	}

	if file.New {
		r.SaveCurrentFileAs()
		return
	}

	r.saveFile(file.Id)
}
func (r *filesOperations) SaveCurrentFileAs() {
	file, getFileErr := r.app.GetOpenedFile()
	if getFileErr != nil {
		r.events.SendErrorEvent("Failed to access current file in memory", getFileErr)
		return
	}

	tmpName := fmt.Sprintf("%s.%s", file.Name, file.Extension)
	filePath, dialErr := r.dialogs.SaveFileDialog(tmpName)
	if dialErr != nil {
		r.events.SendErrorEvent("Failed to process Save File Dialog", dialErr)
		return
	}

	file.Path = filePath
	r.saveFile(file.Id)
}
func (r *filesOperations) saveFile(fileId int64) {
	saveErr := r.app.SaveFile(fileId)
	if saveErr != nil {
		r.events.SendErrorEvent("Failed to save current file", saveErr)
		return
	}

	r.events.SendEvent(api.EventOnFileSaved)
}
func (r *filesOperations) OpenCurrentFileInfo() {
	r.events.SendEvent(api.EventOnFileInformationDisplay)
}
func (r *filesOperations) EditCurrentFileInfo() {
	r.events.SendEvent(api.EventOnFileInformationEdit)
}
func (r *filesOperations) CloseCurrentFile() {
	file, getFileErr := r.app.GetOpenedFile()
	if getFileErr != nil {
		r.events.SendErrorEvent("Failed to access current file in memory", getFileErr)
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
func (r *filesOperations) CloseApplication() {
	var hasChanges bool
	allFilesInfo := r.app.GetAllFilesInfo()
	for _, file := range allFilesInfo {
		if file.Changed {
			hasChanges = true
			return
		}
	}
	if !hasChanges {
		runtime.Quit(r.getContext())
		return
	}

	const title = "Close application?"
	const message = "Files in editor have changes. Do you want proceed and close all files? (Changes will not be saved)"

	btn, dialErr := r.dialogs.OkCancelMessageDialog(title, message)
	if dialErr != nil {
		r.events.SendErrorEvent("Failed to process Message dialog", dialErr)
		return
	}

	if btn.EqualTo(api.BtnOk) {
		runtime.Quit(r.getContext())
	}
}

func CreateFilesOperations(ctx api.ContextProvider, app api.IEditor, events api.IEvents, dialogs api.IDialogs) api.IFilesOperations {
	if events == nil || dialogs == nil || app == nil || ctx == nil {
		panic("Cant create IFilesOperations because one of the required params is nil")
	}
	return &filesOperations{
		events:     events,
		dialogs:    dialogs,
		app:        app,
		getContext: ctx,
	}
}
