package appmenu

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/api"
	"simple_text_editor/core/constants"
	"simple_text_editor/core/logic/files"
)

type AppMenu struct {
	contextRetriever api.GetContext
	appFilesMap      *map[int64]*api.FileApi
	applicationApi   api.EditorApplicationApi
	dialogs          api.DialogsApi
}

func (r *AppMenu) getFilesMap() map[int64]*api.FileApi {
	return *r.appFilesMap
}

func (r *AppMenu) CreateMenu() *menu.Menu {
	appMenu := menu.NewMenu()

	file := appMenu.AddSubmenu("FileApi")
	file.AddText("New", keys.CmdOrCtrl("N"), r.menuFileNewItemClicked)
	file.AddText("Open", keys.CmdOrCtrl("O"), r.menuFileOpenItemClicked)
	file.AddText("Save", keys.CmdOrCtrl("S"), r.menuFileSaveItemClicked)
	file.AddText("Save as", nil, r.menuFileSaveAsItemClicked)
	file.AddText("Close FileApi", nil, r.menuFileCloseFileItemClicked)
	file.AddText("Close Application", keys.CmdOrCtrl("Q"), r.menuFileCloseAppItemClicked)

	edit := appMenu.AddSubmenu("Edit")
	edit.AddText("Sort", keys.CmdOrCtrl("L"), r.menuEditSortItemClicked)

	return appMenu
}
func (r *AppMenu) GetContext() *context.Context {
	ctx := r.contextRetriever()
	return &ctx
}
func (r *AppMenu) SendEvent(destination string, optionalData ...interface{}) {
	ctx := r.GetContext()
	runtime.EventsEmit(*ctx, destination, optionalData...)
}
func (r *AppMenu) menuFileNewItemClicked(*menu.CallbackData) {
	emptyFileRef := r.applicationApi.CreateEmptyFile()
	addedFileRef := r.applicationApi.AddEmptyFile(emptyFileRef)
	addedFileInf := *(*addedFileRef).GetInformationRef()

	r.applicationApi.ChangeFileStatusToOpened(addedFileInf.GetOpenTimeStamp())

	r.SendEvent(constants.EventOnNewFileCreate, addedFileRef)
}
func (r *AppMenu) menuFileOpenItemClicked(*menu.CallbackData) {
	filePath, dialogErr := r.dialogs.OpenFileDialog()
	if dialogErr != nil {
		return
	}

	fileAlreadyOpened := r.applicationApi.IsFileAlreadyOpened(filePath)
	if fileAlreadyOpened {
		sendErrorGenericMessage(r, "FileApi is already opened in application")
		return
	}

	fileContent, ioError := files.GetTextFromFile(filePath)
	if ioError != nil {
		sendErrorIO(r, dialogErr)
		return
	}

	existingFileRef := r.applicationApi.CreateExistingFile(filePath, fileContent)
	addedFileRef := r.applicationApi.AddExistingFile(existingFileRef)
	informationObj := *(*addedFileRef).GetInformationRef()

	r.applicationApi.ChangeFileStatusToOpened(informationObj.GetOpenTimeStamp())

	r.SendEvent(constants.EventOnFileOpened, addedFileRef)
}
func (r *AppMenu) menuFileSaveItemClicked(data *menu.CallbackData) {
	openedFile := r.applicationApi.FindOpenedFile()
	if openedFile == nil {
		sendErrorGenericMessage(r, "Active file is not found. Internal error of app.")
		return
	}

	info := *openedFile.GetInformationRef()
	if !info.GetExists() {
		r.menuFileSaveAsItemClicked(data)
		return
	}

	ioError := files.SaveTextToFile(info.GetPath(), openedFile.GetActualContent())
	if ioError != nil {
		sendErrorIO(r, ioError)
		return
	}

	openedFile.SetOriginalContent(openedFile.GetActualContent())

	r.SendEvent(constants.EventOnFileSaved, openedFile)
}
func (r *AppMenu) menuFileSaveAsItemClicked(*menu.CallbackData) {
	filePath, dialogErr := r.dialogs.SaveFileDialog()
	if dialogErr != nil {
		return
	}

	openedFile := r.applicationApi.FindOpenedFile()
	if openedFile == nil {
		sendErrorGenericMessage(r, "Active file is not found. Internal error of app.")
		return
	}

	ioError := files.SaveTextToFile(filePath, openedFile.GetActualContent())
	if ioError != nil {
		sendErrorIO(r, ioError)
		return
	}

	infoObj := *openedFile.GetInformationRef()
	infoObj.SetPath(filePath)

	openedFile.SetOriginalContent(openedFile.GetActualContent())

	r.SendEvent(constants.EventOnFileSaved, openedFile)
}
func (r *AppMenu) menuFileCloseFileItemClicked(*menu.CallbackData) {
	openedFile := r.applicationApi.FindOpenedFile()
	if openedFile != nil {
		sendErrorGenericMessage(r, "Problem happened with getting active file (opened now in application)")
		return
	}

	infoObj := *openedFile.GetInformationRef()
	uniqueId := infoObj.GetOpenTimeStamp()

	if !openedFile.HasChanges() {
		r.closeFileAndChoseNextOrNew(uniqueId)
		return
	}

	dialogResult, err := r.dialogs.OkCancelMessageDialog(
		"FileApi has changes",
		"You are trying to close file that have changes, continue?",
	)
	if err != nil {
		return
	}

	if dialogResult == "Ok" {
		r.closeFileAndChoseNextOrNew(uniqueId)
		return
	}
}
func (r *AppMenu) closeFileAndChoseNextOrNew(uniqueId int64) {
	r.applicationApi.CloseFile(uniqueId)
	anyFile := r.applicationApi.FindAnyFileInMemory()
	anyFileInfo := *(*anyFile).GetInformationRef()
	r.applicationApi.ChangeFileStatusToOpened(anyFileInfo.GetOpenTimeStamp())

	r.SendEvent(constants.EventOnFileClosed)
}
func (r *AppMenu) menuFileCloseAppItemClicked(*menu.CallbackData) {
	var hasChanges bool
	for _, file := range r.getFilesMap() {
		if (*file).HasChanges() {
			hasChanges = true
			break
		}
	}

	if !hasChanges {
		runtime.Quit(*r.GetContext())
	}

	dialogResult, err := r.dialogs.OkCancelMessageDialog(
		"flaskApplicationApi has files with changes",
		"You are trying to close application that have files with changes, continue?",
	)
	if err != nil {
		return
	}

	if dialogResult == "Ok" {
		runtime.Quit(*r.GetContext())
	}
}
func (r *AppMenu) menuEditSortItemClicked(*menu.CallbackData) {

}

func CreateApplicationMenu(contextRetriever api.GetContext,
	appFilesMap *map[int64]*api.FileApi,
	applicationApi api.EditorApplicationApi,
	dialogs api.DialogsApi) api.ApplicationMenu {
	return &AppMenu{
		contextRetriever: contextRetriever,
		appFilesMap:      appFilesMap,
		applicationApi:   applicationApi,
		dialogs:          dialogs,
	}
}
