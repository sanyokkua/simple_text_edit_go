package menuopshelper

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/v3/components/eventsender"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/validators"
)

var RuntimeQuit = runtime.Quit

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

	var fileName string
	var fileExtension string

	if file.New && file.Name == types.NewFileName {
		fileName = ""
	} else {
		fileName = file.Name
	}
	if file.New && len(file.Extension) == 0 {
		fileExtension = ".txt"
	} else {
		fileExtension = file.Extension
	}

	filePath, dialErr := r.DialogHelper.SaveFileDialog(fileName + fileExtension)
	if validators.HasError(dialErr) {
		r.EventSender.SendErrorEvent("Failed to process Save File Dialog", dialErr)
		return
	}

	file.Path = filePath
	r.saveFile(file.Id)
}

func (r *MenuHelperOperationsStruct) CloseFile() {
	r.EventSender.SendNotificationEvent(eventsender.EventOnBlockUiTrue)

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
			r.EventSender.SendNotificationEvent(eventsender.EventOnBlockUiFalse)
			return
		}

		closeErr := r.Editor.CloseFile(file.Id)
		if closeErr != nil {
			r.EventSender.SendErrorEvent("Failed to close file", closeErr)
			return
		}
	}

	r.EventSender.SendNotificationEvent(eventsender.EventOnFileClosed)
	r.EventSender.SendNotificationEvent(eventsender.EventOnBlockUiFalse)
}

func (r *MenuHelperOperationsStruct) CloseApplication() {
	r.EventSender.SendNotificationEvent(eventsender.EventOnBlockUiTrue)

	allFilesInfo, getFilesErr := r.Editor.GetFilesShortInfo()
	if validators.HasError(getFilesErr) {
		r.EventSender.SendErrorEvent("Failed to get all files")
		return
	}
	var hasChanges bool
	for _, file := range allFilesInfo {
		if file.Changed {
			hasChanges = true
			break
		}
	}
	if !hasChanges {
		RuntimeQuit(r.GetContext())
		return
	}

	const title = "Close application?"
	const message = "Files in editor have changes. Do you want proceed and close all files? (Changes will not be saved)"

	btn, dialErr := r.DialogHelper.OkCancelMessageDialog(title, message)
	if validators.HasError(dialErr) {
		r.EventSender.SendErrorEvent("Failed to process Message dialog", dialErr)
		return
	}

	r.EventSender.SendNotificationEvent(eventsender.EventOnBlockUiFalse)

	if btn.EqualTo(types.BtnOk) {
		RuntimeQuit(r.GetContext())
		return
	}
}

func (r *MenuHelperOperationsStruct) ChangeExtension(fileType types.FileTypeKey, fileExtension types.FileTypeExtension) {
	file, getFileErr := r.Editor.GetOpenedFile()
	if validators.HasError(getFileErr) {
		r.EventSender.SendErrorEvent("Failed to access current file in memory", getFileErr)
		return
	}

	updateStruct := types.FileTypeUpdateStruct{
		Id:        file.Id,
		Type:      string(fileType),
		Extension: string(fileExtension),
	}

	updateErr := r.Editor.UpdateFileInformation(file.Id, updateStruct)
	if validators.HasError(updateErr) {
		r.EventSender.SendWarnEvent("File information was not updated", updateErr)
		return
	}

	r.EventSender.SendNotificationEvent(eventsender.EventOnFileInformationUpdated)
}

func (r *MenuHelperOperationsStruct) BlockUI(state bool) {
	var dst types.Destination
	if state {
		dst = eventsender.EventOnBlockUiTrue
	} else {
		dst = eventsender.EventOnBlockUiFalse
	}
	r.EventSender.SendNotificationEvent(dst)
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
