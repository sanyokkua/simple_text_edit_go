package core

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/apperrors"
	"simple_text_editor/core/constants"
	"simple_text_editor/core/logic/files"
)

func (receiver *EditorApplication) menuFileNewItemClicked(data *menu.CallbackData) {
	file := receiver.createEmptyFile()
	receiver.ActivateFile(file.Descriptor.OpenTimeStamp)
	runtime.EventsEmit(receiver.AppContext, constants.EventOnNewFileCreate, file)
}
func (receiver *EditorApplication) menuFileOpenItemClicked(data *menu.CallbackData) {
	filePath, err := runtime.OpenFileDialog(receiver.AppContext, runtime.OpenDialogOptions{
		Title:   "Open Text File",
		Filters: constants.GetSupportedFileFilters(),
	})

	if err != nil {
		apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
			AppContext:    receiver.AppContext,
			Destination:   constants.EventOnErrorHappened,
			OriginalError: err,
			Message:       "Problem happened with opening OpenFileDialog",
		})
		return
	}

	if receiver.IsFileAlreadyOpened(filePath) {
		apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
			AppContext:  receiver.AppContext,
			Destination: constants.EventOnErrorHappened,
			Message:     "File is already opened",
		})
		return
	}

	fileContent, err := files.GetTextFromFile(filePath)

	if err != nil {
		apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
			AppContext:    receiver.AppContext,
			Destination:   constants.EventOnErrorHappened,
			OriginalError: err,
			Message:       "Problem happened with reading file content",
		})
		return
	}

	file := CreateExistingFileInMemory(filePath, fileContent)
	receiver.addFileInMemory(&file)
	receiver.ActivateFile(file.Descriptor.OpenTimeStamp)
	runtime.EventsEmit(receiver.AppContext, constants.EventOnFileOpened, file)
}
func (receiver *EditorApplication) menuFileSaveItemClicked(data *menu.CallbackData) {
	currentFile, getFileErr := receiver.GetActiveFile()

	if getFileErr != nil {
		apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
			AppContext:    receiver.AppContext,
			Destination:   constants.EventOnErrorHappened,
			OriginalError: getFileErr,
			Message:       "Problem happened with getting active file (opened now in editor)",
		})
		return
	}

	if currentFile.Descriptor.Exists {
		saveErr := files.SaveTextToFile(currentFile.Descriptor.Path, currentFile.ActualContent)

		if saveErr != nil {
			apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
				AppContext:    receiver.AppContext,
				Destination:   constants.EventOnErrorHappened,
				OriginalError: saveErr,
				Message:       "Problem happened with saving active file (opened now in editor)",
			})
			return
		}

		currentFile.SetOriginalContent(currentFile.ActualContent)
		runtime.EventsEmit(receiver.AppContext, constants.EventOnFileSaved, currentFile)
	} else {
		receiver.menuFileSaveAsItemClicked(data)
	}
}

func (receiver *EditorApplication) menuFileSaveAsItemClicked(data *menu.CallbackData) {
	filePath, err := runtime.SaveFileDialog(receiver.AppContext, runtime.SaveDialogOptions{
		Title:   "Save File As...",
		Filters: constants.GetSupportedFileFilters(),
	})

	if err != nil {
		apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
			AppContext:    receiver.AppContext,
			Destination:   constants.EventOnErrorHappened,
			OriginalError: err,
			Message:       "Problem happened with opening SaveFileDialog",
		})
		return
	}

	currentFile, getFileErr := receiver.GetActiveFile()

	if getFileErr != nil {
		apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
			AppContext:    receiver.AppContext,
			Destination:   constants.EventOnErrorHappened,
			OriginalError: getFileErr,
			Message:       "Problem happened with getting active file (opened now in editor)",
		})
		return
	}

	saveErr := files.SaveTextToFile(filePath, currentFile.ActualContent)

	if saveErr != nil {
		apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
			AppContext:    receiver.AppContext,
			Destination:   constants.EventOnErrorHappened,
			OriginalError: saveErr,
			Message:       "Problem happened with saving active file (opened now in editor)",
		})
		return
	}

	currentFile.Descriptor.setPath(filePath)
	currentFile.SetOriginalContent(currentFile.ActualContent)
	runtime.EventsEmit(receiver.AppContext, constants.EventOnFileSaved, currentFile)
}
func (receiver *EditorApplication) menuFileCloseFileItemClicked(data *menu.CallbackData) {
	currentFile, getFileErr := receiver.GetActiveFile()

	if getFileErr != nil {
		apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
			AppContext:    receiver.AppContext,
			Destination:   constants.EventOnErrorHappened,
			OriginalError: getFileErr,
			Message:       "Problem happened with getting active file (opened now in editor)",
		})
		return
	}

	if currentFile.HasChanges() {
		dialogResult, err := runtime.MessageDialog(receiver.AppContext, runtime.MessageDialogOptions{
			Type:          runtime.WarningDialog,
			Title:         "File has changes",
			Message:       "You are trying to close file that have changes, continue?",
			Buttons:       []string{"Ok", "Cancel"},
			DefaultButton: "Ok",
			CancelButton:  "Cancel",
		})
		if err != nil {
			apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
				AppContext:    receiver.AppContext,
				Destination:   constants.EventOnErrorHappened,
				OriginalError: err,
				Message:       "Problem with opening close file dialog",
			})
			return
		}

		if dialogResult == "Ok" {
			delete(receiver.FilesInMemory, currentFile.Descriptor.OpenTimeStamp)
			if len(receiver.FilesInMemory) > 0 {
				var firstFile *FileInMemory
				for _, file := range receiver.FilesInMemory {
					firstFile = file
					break
				}
				receiver.ActivateFile(firstFile.Descriptor.OpenTimeStamp)
			} else {
				newFile := receiver.createEmptyFile()
				receiver.ActivateFile(newFile.Descriptor.OpenTimeStamp)
			}
			runtime.EventsEmit(receiver.AppContext, constants.EventOnFileClosed)
		}
	} else {
		delete(receiver.FilesInMemory, currentFile.Descriptor.OpenTimeStamp)
		if len(receiver.FilesInMemory) > 0 {
			var firstFile *FileInMemory
			for _, file := range receiver.FilesInMemory {
				firstFile = file
				break
			}
			receiver.ActivateFile(firstFile.Descriptor.OpenTimeStamp)
		} else {
			newFile := receiver.createEmptyFile()
			receiver.ActivateFile(newFile.Descriptor.OpenTimeStamp)
		}
		runtime.EventsEmit(receiver.AppContext, constants.EventOnFileClosed)
	}
}
func (receiver *EditorApplication) menuFileCloseAppItemClicked(data *menu.CallbackData) {
	var hasChanges bool
	for _, file := range receiver.FilesInMemory {
		if file.HasChanges() {
			hasChanges = true
			break
		}
	}
	if hasChanges {
		dialogResult, err := runtime.MessageDialog(receiver.AppContext, runtime.MessageDialogOptions{
			Type:          runtime.WarningDialog,
			Title:         "Editor has files with changes",
			Message:       "You are trying to close application that have files with changes, continue?",
			Buttons:       []string{"Ok", "Cancel"},
			DefaultButton: "Ok",
			CancelButton:  "Cancel",
		})
		if err != nil {
			apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
				AppContext:    receiver.AppContext,
				Destination:   constants.EventOnErrorHappened,
				OriginalError: err,
				Message:       "Problem with opening close file dialog",
			})
			return
		}

		if dialogResult == "Ok" {
			runtime.Quit(receiver.AppContext)
		}
	} else {
		runtime.Quit(receiver.AppContext)
	}
}
func (receiver *EditorApplication) menuEditSortItemClicked(data *menu.CallbackData) {}
