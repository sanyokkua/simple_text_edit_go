package core

import (
	"context"
	"errors"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/apperrors"
	"simple_text_editor/core/constants"
)

type EditorApplication struct {
	AppContext    context.Context         // Application context
	FilesInMemory map[int64]*FileInMemory // Files that currently opened in the app memory
}

func (receiver *EditorApplication) createEmptyFile() *FileInMemory {
	file := CreateEmptyFileInMemory()
	receiver.FilesInMemory[file.Descriptor.OpenTimeStamp] = &file
	return &file
}

func (receiver *EditorApplication) addFileInMemory(file *FileInMemory) {
	if file == nil {
		// TODO: process error
		return
	}
	receiver.FilesInMemory[file.Descriptor.OpenTimeStamp] = file
}

func (receiver *EditorApplication) inactivateAllFiles() {
	for _, openedFile := range receiver.FilesInMemory {
		openedFile.Descriptor.IsOpenedNow = false
	}
}

func (receiver *EditorApplication) ActivateFile(fileId int64) {
	fileToActivate, ok := receiver.FilesInMemory[fileId]
	if !ok {
		return
	}
	receiver.inactivateAllFiles()
	fileToActivate.Descriptor.IsOpenedNow = true
}

func (receiver *EditorApplication) GetActiveFile() (*FileInMemory, error) {
	for _, file := range receiver.FilesInMemory {
		if file.Descriptor.IsOpenedNow {
			return file, nil
		}
	}
	return nil, errors.New("active file is not found")
}

func (receiver *EditorApplication) GetAllFilesInformation() []FileInformation {
	allOpenedFiles := make([]FileInformation, 0, len(receiver.FilesInMemory))

	for _, file := range receiver.FilesInMemory {
		descriptor := file.Descriptor
		allOpenedFiles = append(allOpenedFiles, descriptor)
	}
	return allOpenedFiles
}

func (receiver *EditorApplication) GetAllFilesInMemory() map[int64]*FileInMemory {
	return receiver.FilesInMemory
}

func (receiver *EditorApplication) IsFileAlreadyOpened(filePath string) bool {
	if len(filePath) == 0 {
		return false
	}
	allFiles := receiver.GetAllFilesInformation()
	for _, file := range allFiles {
		if file.Path == filePath {
			return true
		}
	}
	return false
}

func (receiver *EditorApplication) UpdateActiveFileContent(content string) bool {
	file, err := receiver.GetActiveFile()
	if err != nil {
		apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
			AppContext:    receiver.AppContext,
			Destination:   constants.EventOnErrorHappened,
			OriginalError: err,
			Message:       "There is a problem with updating file content on the backend",
		})
		return false
	}
	file.SetActualContent(content)
	return file.HasChanges()
}

func (receiver *EditorApplication) Startup(ctx context.Context) {
	receiver.AppContext = ctx
	receiver.registerEventHandlers()
}

func (receiver *EditorApplication) CreateMenu() *menu.Menu {
	appMenu := menu.NewMenu()

	file := appMenu.AddSubmenu("File")
	file.AddText("New", keys.CmdOrCtrl("N"), receiver.menuFileNewItemClicked)
	file.AddText("Open", keys.CmdOrCtrl("O"), receiver.menuFileOpenItemClicked)
	file.AddText("Save", keys.CmdOrCtrl("S"), receiver.menuFileSaveItemClicked)
	file.AddText("Save as", nil, receiver.menuFileSaveItemClicked)
	file.AddText("Close File", nil, receiver.menuFileCloseFileItemClicked)
	file.AddText("Close Application", keys.CmdOrCtrl("Q"), receiver.menuFileCloseAppItemClicked)

	edit := appMenu.AddSubmenu("Edit")
	edit.AddSeparator()
	edit.AddText("Sort", keys.CmdOrCtrl("L"), receiver.menuEditSortItemClicked)

	return appMenu
}

func (receiver *EditorApplication) registerEventHandlers() {
	runtime.EventsOn(receiver.AppContext, constants.EventOnActiveFileContentUpdated, func(optionalData ...interface{}) {

	})
}
func CreateNewApplication() *EditorApplication {
	editorApp := &EditorApplication{
		FilesInMemory: make(map[int64]*FileInMemory),
	}
	createdFile := editorApp.createEmptyFile()
	editorApp.ActivateFile(createdFile.Descriptor.OpenTimeStamp)
	return editorApp
}
