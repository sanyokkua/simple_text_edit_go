package frontendapi

import (
	"fmt"
	"simple_text_editor/core/v3/components/eventsender"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/validators"
	"sort"
)

type FrontendApiStruct struct {
	Editor       types.IEditor
	EventSender  types.IEventSender
	DialogHelper types.IDialogHelper
	TypeManager  types.ITypeManager
}

func (r *FrontendApiStruct) NewFile() {
	fileCreationErr := r.Editor.CreateFileAndShow()

	if validators.HasError(fileCreationErr) {
		r.EventSender.SendErrorEvent("IApplication failed to create new file", fileCreationErr)
		return
	}

	r.EventSender.SendNotificationEvent(eventsender.EventOnNewFileCreated)
}

func (r *FrontendApiStruct) GetFilesShortInfo() types.FrontendFileInfoArrayContainerStruct {
	filesInfo, getFilesErr := r.Editor.GetFilesShortInfo()

	if validators.HasError(getFilesErr) {
		return types.FrontendFileInfoArrayContainerStruct{
			HasError: true,
			Error:    getFilesErr.Error(),
			Files:    []types.FileInfoStruct{},
		}
	}

	sort.Slice(filesInfo, func(i, j int) bool {
		return filesInfo[i].Id < filesInfo[j].Id
	})

	return types.FrontendFileInfoArrayContainerStruct{
		Files: filesInfo,
	}
}

func (r *FrontendApiStruct) GetOpenedFile() types.FrontendFileContainerStruct {
	file, getOpenedFileErr := r.Editor.GetOpenedFile()
	if validators.HasError(getOpenedFileErr) {
		return types.FrontendFileContainerStruct{
			HasError: true,
			Error:    getOpenedFileErr.Error(),
			File:     types.FileStruct{},
		}
	}

	return types.FrontendFileContainerStruct{
		File: *file,
	}
}

// GetFileTypes Deprecated. TODO: Remove
func (r *FrontendApiStruct) GetFileTypes() types.FrontendKeyValueArrayContainerStruct {
	fileTypes, buildErr := r.TypeManager.BuildFileTypeMappingKeyToName()
	if validators.HasError(buildErr) {
		return types.FrontendKeyValueArrayContainerStruct{
			HasError: true,
			Error:    buildErr.Error(),
			Pairs:    []types.KeyValuePairStruct{},
		}
	}

	return types.FrontendKeyValueArrayContainerStruct{
		Pairs: fileTypes,
	}
}

// GetFileTypeExtension Deprecated. TODO: Remove
func (r *FrontendApiStruct) GetFileTypeExtension(fileTypeKey string) types.FrontendKeyValueArrayContainerStruct {
	fileExtensions, buildErr := r.TypeManager.BuildFileTypeMappingExtToExt(types.FileTypeKey(fileTypeKey))
	if validators.HasError(buildErr) {
		return types.FrontendKeyValueArrayContainerStruct{
			HasError: true,
			Error:    buildErr.Error(),
			Pairs:    []types.KeyValuePairStruct{},
		}
	}

	return types.FrontendKeyValueArrayContainerStruct{
		Pairs: fileExtensions,
	}
}

func (r *FrontendApiStruct) SwitchOpenedFileTo(fileId int64) {
	showErr := r.Editor.ShowFile(fileId)
	if validators.HasError(showErr) {
		r.EventSender.SendErrorEvent("Can't show file with provided file ID", showErr)
		return
	}

	r.EventSender.SendNotificationEvent(eventsender.EventOnFileIsSwitched)
}

func (r *FrontendApiStruct) UpdateFileContent(fileId int64, content string) {
	updateErr := r.Editor.UpdateFileContent(fileId, content)
	if validators.HasError(updateErr) {
		r.EventSender.SendErrorEvent("File content was not updated", updateErr)
		return
	}

	r.EventSender.SendNotificationEvent(eventsender.EventOnFileContentIsUpdated)
}

func (r *FrontendApiStruct) CloseFile(fileId int64) {
	r.EventSender.SendNotificationEvent(eventsender.EventOnBlockUiTrue)

	file, getFileErr := r.Editor.GetFileById(fileId)
	if validators.HasError(getFileErr) {
		r.EventSender.SendErrorEvent("Failed to access file by id in memory", getFileErr)
		return
	}

	if !file.Changed {
		closeErr := r.Editor.CloseFile(file.Id)
		if closeErr != nil {
			r.EventSender.SendErrorEvent("Failed to close file", closeErr)
			return
		}
	} else {
		var title = fmt.Sprintf("Close file? (%s)", file.Name)
		const message = "File has changes. Do you want proceed and close file? (Changes will not be saved)"

		btn, dialErr := r.DialogHelper.OkCancelMessageDialog(title, message)
		if dialErr != nil {
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

	r.EventSender.SendNotificationEvent(eventsender.EventOnBlockUiFalse)
	r.EventSender.SendNotificationEvent(eventsender.EventOnFileClosed)
}

func CreateIFrontendApi(editor types.IEditor, eventSender types.IEventSender,
	dialogHelper types.IDialogHelper, typeManager types.ITypeManager) types.IFrontendApi {
	validators.PanicOnNil(editor, "IEditor")
	validators.PanicOnNil(eventSender, "IEventSender")
	validators.PanicOnNil(dialogHelper, "IDialogHelper")
	validators.PanicOnNil(typeManager, "ITypeManager")

	return &FrontendApiStruct{
		Editor:       editor,
		EventSender:  eventSender,
		DialogHelper: dialogHelper,
		TypeManager:  typeManager,
	}
}
