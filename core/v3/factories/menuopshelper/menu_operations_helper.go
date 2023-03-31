package menuopshelper

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/v3/factories/eventsender"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/validators"
)

type MenuHelperOperationsStruct struct {
	GetContext   types.ContextProvider
	EventSender  types.IEventSender
	DialogHelper types.IDialogHelper
	Editor       types.IEditor
}

func (r *MenuHelperOperationsStruct) NewFile() {
	fileCreationErr := r.Editor.CreateFileAndShow()

	if validators.HasError(fileCreationErr) {
		r.EventSender.SendErrorEvent("IApplication failed to create new file", fileCreationErr)
		return
	}

	r.EventSender.SendNotificationEvent(eventsender.EventOnNewFileCreated)
}

func (r *MenuHelperOperationsStruct) OpenFile() {
	filePath, dialErr := r.DialogHelper.OpenFileDialog()
	if validators.HasError(dialErr) {
		r.EventSender.SendErrorEvent("Failed to process Open File Dialog", dialErr)
		return
	}

	openFileErr := r.Editor.OpenFileAndShow(filePath)
	if validators.HasError(openFileErr) {
		r.EventSender.SendErrorEvent("Failed to Open File", openFileErr)
		return
	}

	r.EventSender.SendNotificationEvent(eventsender.EventOnFileOpened)
}

func (r *MenuHelperOperationsStruct) SaveFile() {
	file, getFileErr := r.Editor.GetOpenedFile()
	if validators.HasError(getFileErr) {
		r.EventSender.SendErrorEvent("Failed to access current file in memory", getFileErr)
		return
	}

	if file.New {
		r.SaveFileAs()
		return
	}

	r.saveFile(file.Id)
}

func (r *MenuHelperOperationsStruct) saveFile(fileId int64) {
	saveErr := r.Editor.SaveFile(fileId)
	if validators.HasError(saveErr) {
		r.EventSender.SendErrorEvent("Failed to save current file", saveErr)
		return
	}

	r.EventSender.SendNotificationEvent(eventsender.EventOnFileSaved)
}

func (r *MenuHelperOperationsStruct) SaveFileAs() {
	file, getFileErr := r.Editor.GetOpenedFile()
	if validators.HasError(getFileErr) {
		r.EventSender.SendErrorEvent("Failed to access current file in memory", getFileErr)
		return
	}

	filePath, dialErr := r.DialogHelper.SaveFileDialog(file.Name)
	if validators.HasError(dialErr) {
		r.EventSender.SendErrorEvent("Failed to process Save File Dialog", dialErr)
		return
	}

	file.Path = filePath
	r.saveFile(file.Id)
}

func (r *MenuHelperOperationsStruct) ShowFileInfoModal() {
	r.EventSender.SendNotificationEvent(eventsender.EventOnFileInformationDisplay)
}

func (r *MenuHelperOperationsStruct) CloseFile() {
	file, getFileErr := r.Editor.GetOpenedFile()
	if validators.HasError(getFileErr) {
		r.EventSender.SendErrorEvent("Failed to access current file in memory", getFileErr)
		return
	}

	if !file.Changed {
		closeErr := r.Editor.CloseFile(file.Id)
		if validators.HasError(closeErr) {
			r.EventSender.SendErrorEvent("Failed to close file", closeErr)
			return
		}
	} else {
		var title = fmt.Sprintf("Close file? (%s)", file.Name)
		const message = "File has changes. Do you want proceed and close file? (Changes will not be saved)"

		btn, dialErr := r.DialogHelper.OkCancelMessageDialog(title, message)
		if validators.HasError(dialErr) {
			r.EventSender.SendErrorEvent("Failed to process Message dialog", dialErr)
			return
		}

		if btn.EqualTo(types.BtnCancel) {
			return
		}

		closeErr := r.Editor.CloseFile(file.Id)
		if closeErr != nil {
			r.EventSender.SendErrorEvent("Failed to close file", closeErr)
			return
		}
	}

	r.EventSender.SendNotificationEvent(eventsender.EventOnFileClosed)
}

func (r *MenuHelperOperationsStruct) CloseApplication() {
	allFilesInfo, getFilesErr := r.Editor.GetFilesShortInfo()
	if validators.HasError(getFilesErr) {
		r.EventSender.SendErrorEvent("Failed to get all files")
		return
	}

	var hasChanges bool
	for _, file := range allFilesInfo {
		if file.Changed {
			hasChanges = true
			return
		}
	}
	if !hasChanges {
		runtime.Quit(r.GetContext())
		return
	}

	const title = "Close application?"
	const message = "Files in editor have changes. Do you want proceed and close all files? (Changes will not be saved)"

	btn, dialErr := r.DialogHelper.OkCancelMessageDialog(title, message)
	if validators.HasError(dialErr) {
		r.EventSender.SendErrorEvent("Failed to process Message dialog", dialErr)
		return
	}

	if btn.EqualTo(types.BtnOk) {
		runtime.Quit(r.GetContext())
	}
}

func CreateIMenuOpsHelper(getContext types.ContextProvider,
	eventSender types.IEventSender,
	dialogHelper types.IDialogHelper,
	editor types.IEditor) types.IMenuOpsHelper {

	validators.PanicOnNil(getContext, "ContextProvider")
	validators.PanicOnNil(eventSender, "IEventSender")
	validators.PanicOnNil(dialogHelper, "IDialogHelper")
	validators.PanicOnNil(editor, "IEditor")

	return &MenuHelperOperationsStruct{
		GetContext:   getContext,
		EventSender:  eventSender,
		DialogHelper: dialogHelper,
		Editor:       editor,
	}

}
