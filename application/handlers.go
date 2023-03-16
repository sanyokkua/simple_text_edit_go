package application

import (
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/constants"
	"simple_text_editor/logic"
)

func CreateNewFile(app *EditorApplication) {
	inactivateAllFiles(app)

	newFile := createAppFileForNewFile()
	app.AppFiles[newFile.Descriptor.FileId] = newFile

	runtime.EventsEmit(app.AppContext, constants.EVENT_IS_FILE_OPENED, newFile)
}

func OpenFileDialog(app *EditorApplication) {
	filePath, err := runtime.OpenFileDialog(app.AppContext, runtime.OpenDialogOptions{
		Title:   "Open Text File",
		Filters: constants.GetSupportedFileFilters(),
	})
	if err != nil {
		processGenericError(app, err, "Problem with opening dialog.")
		return
	}
	log.Info("Open File path: ", filePath)

	fileContent, err := logic.GetTextFromFile(filePath)
	if err != nil {
		processGenericError(app, err, "Problem with reading file.")
		return
	}

	inactivateAllFiles(app)

	openedFileStruct := createAppFileForExistingFile(filePath, fileContent)
	app.AppFiles[openedFileStruct.Descriptor.FileId] = openedFileStruct

	runtime.EventsEmit(app.AppContext, constants.EVENT_IS_FILE_OPENED, openedFileStruct)
}

func SaveFileDialog(app *EditorApplication) {
	filePath, err := runtime.SaveFileDialog(app.AppContext, runtime.SaveDialogOptions{
		Title:   "Save File As...",
		Filters: constants.GetSupportedFileFilters(),
	})

	if err != nil {
		processGenericError(app, err, "Problem with opening dialog.")
		return
	}
	log.Info("Save File Path: ", filePath)

	var currentFile *AppFile
	for _, openedFile := range app.AppFiles {
		if openedFile.Descriptor.IsActive {
			currentFile = openedFile
			break
		}
	}

	if !currentFile.Descriptor.IsActive {
		saveFileError := errors.New("current file is not found")
		processGenericError(app, saveFileError, "Problem with saving file")
		return
	}

	saveErr := logic.SaveTextToFile(filePath, currentFile.FileContent)
	if saveErr != nil {
		processGenericError(app, saveErr, "Problem with saving file content")
		return
	}
}
