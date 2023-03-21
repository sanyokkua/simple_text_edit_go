package menu

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/api"
	"simple_text_editor/core/constants"
	"simple_text_editor/core/implementation/utils"
	"simple_text_editor/core/logic/files"
)

type AppMenu struct {
	contextRetriever api.ContextRetriever
	applicationApi   api.EditorApplication
	dialogs          api.DialogsApi
}

func (r *AppMenu) getFilesMap() map[int64]*api.OpenedFile {
	log.Info("getFilesMap", *r.applicationApi.GetFilesMap())
	return *r.applicationApi.GetFilesMap()
}

func (r *AppMenu) CreateMenu() *menu.Menu {
	log.Info("CreateMenu")
	appMenu := menu.NewMenu()

	file := appMenu.AddSubmenu("OpenedFile")
	file.AddText("New", keys.CmdOrCtrl("N"), r.menuFileNewItemClicked)
	file.AddText("Open", keys.CmdOrCtrl("O"), r.menuFileOpenItemClicked)
	file.AddText("Save", keys.CmdOrCtrl("S"), r.menuFileSaveItemClicked)
	file.AddText("Save as", nil, r.menuFileSaveAsItemClicked)
	file.AddText("Get File Info", nil, r.menuFileGetFileInfoClicked)
	file.AddText("Close OpenedFile", nil, r.menuFileCloseFileItemClicked)
	file.AddText("Close Application", keys.CmdOrCtrl("Q"), r.menuFileCloseAppItemClicked)

	edit := appMenu.AddSubmenu("Edit")
	edit.AddText("Edit File Information", nil, r.menuEditFileInformationClicked)
	edit.AddText("Sort", keys.CmdOrCtrl("L"), r.menuEditSortItemClicked)

	return appMenu
}
func (r *AppMenu) GetContext() *context.Context {
	log.Info("GetContext")
	ctx := r.contextRetriever()
	return &ctx
}
func (r *AppMenu) SendEvent(destination string, optionalData ...interface{}) {
	log.Info("SendEvent", destination, optionalData)
	ctx := r.GetContext()
	runtime.EventsEmit(*ctx, destination, optionalData...)
}
func (r *AppMenu) menuFileNewItemClicked(*menu.CallbackData) {
	log.Info("menuFileNewItemClicked")
	emptyFileRef := utils.CreateEmptyFile()
	addedFileRef := r.applicationApi.AddFileToMemory(&emptyFileRef)
	addedFileInf := *(*addedFileRef).GetInformation()

	r.applicationApi.ChangeFileStatusToOpened(addedFileInf.GetOpenTimeStamp())

	r.SendEvent(constants.EventOnNewFileCreate, addedFileRef)
}
func (r *AppMenu) menuFileOpenItemClicked(*menu.CallbackData) {
	log.Info("menuFileOpenItemClicked")
	filePath, dialogErr := r.dialogs.OpenFileDialog()
	if dialogErr != nil {
		return
	}

	fileAlreadyOpened := r.applicationApi.IsFileAlreadyOpened(filePath)
	if fileAlreadyOpened {
		sendErrorGenericMessage(r, "OpenedFile is already opened in application")
		return
	}

	fileContent, ioError := files.GetTextFromFile(filePath)
	if ioError != nil {
		sendErrorIO(r, dialogErr)
		return
	}

	existingFileRef := utils.CreateExistingFile(filePath, fileContent)
	addedFileRef := r.applicationApi.AddFileToMemory(&existingFileRef)
	informationObj := *(*addedFileRef).GetInformation()

	r.applicationApi.ChangeFileStatusToOpened(informationObj.GetOpenTimeStamp())

	r.SendEvent(constants.EventOnFileOpened, addedFileRef)
}
func (r *AppMenu) menuFileSaveItemClicked(data *menu.CallbackData) {
	log.Info("menuFileSaveItemClicked")
	openedFile := r.applicationApi.FindOpenedFile()
	if openedFile == nil {
		sendErrorGenericMessage(r, "Active file is not found. Internal error of app.")
		return
	}

	info := *openedFile.GetInformation()
	if !info.Exists() {
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
	log.Info("menuFileSaveAsItemClicked")
	openedFile := r.applicationApi.FindOpenedFile()
	if openedFile == nil {
		sendErrorGenericMessage(r, "Active file is not found. Internal error of app.")
		return
	}
	info := *openedFile.GetInformation()

	defaultName := info.GetName() + "." + info.GetExt()
	filePath, dialogErr := r.dialogs.SaveFileDialog(defaultName)
	if dialogErr != nil {
		return
	}

	extInPath := utils.GetFileExtensionFromPath(filePath)
	var extToAdd string
	if len(extInPath) > 0 {
		extToAdd = ""
	} else {
		if len(info.GetExt()) > 0 {
			extToAdd = info.GetExt()
		} else {
			extToAdd = ".txt"
		}
	}
	pathWithExt := filePath + extToAdd

	ioError := files.SaveTextToFile(pathWithExt, openedFile.GetActualContent())
	if ioError != nil {
		sendErrorIO(r, ioError)
		return
	}

	infoObj := *openedFile.GetInformation()

	upd := utils.CreateInformationFromPath(pathWithExt)
	infoObj.SetPath((*upd).GetPath())
	infoObj.SetName((*upd).GetName())
	infoObj.SetExt((*upd).GetExt())
	infoObj.SetType((*upd).GetType())
	infoObj.SetExists((*upd).Exists())

	openedFile.SetOriginalContent(openedFile.GetActualContent())

	r.SendEvent(constants.EventOnFileSaved, openedFile)
}
func (r *AppMenu) menuFileCloseFileItemClicked(*menu.CallbackData) {
	log.Info("menuFileCloseFileItemClicked")
	openedFile := r.applicationApi.FindOpenedFile()
	if openedFile != nil {
		sendErrorGenericMessage(r, "Problem happened with getting active file (opened now in application)")
		return
	}

	infoObj := *openedFile.GetInformation()
	uniqueId := infoObj.GetOpenTimeStamp()

	if !openedFile.HasChanges() {
		log.Info("menuFileCloseFileItemClicked, Doesn't have changes")
		r.closeFileAndChoseNextOrNew(uniqueId)
		return
	}
	log.Info("menuFileCloseFileItemClicked, Have changes")

	dialogResult, err := r.dialogs.OkCancelMessageDialog(
		"OpenedFile has changes",
		"You are trying to close file that have changes, continue?",
	)
	if err != nil {
		return
	}

	if dialogResult == "Ok" {
		log.Info("menuFileCloseFileItemClicked, Dialog result - OK")
		r.closeFileAndChoseNextOrNew(uniqueId)
		return
	}
}
func (r *AppMenu) closeFileAndChoseNextOrNew(uniqueId int64) {
	log.Info("closeFileAndChoseNextOrNew", uniqueId)
	r.applicationApi.CloseFile(uniqueId)
	anyFile := r.applicationApi.FindAnyFileInMemory()
	anyFileInfo := *(*anyFile).GetInformation()
	r.applicationApi.ChangeFileStatusToOpened(anyFileInfo.GetOpenTimeStamp())

	r.SendEvent(constants.EventOnFileClosed)
}
func (r *AppMenu) menuFileCloseAppItemClicked(*menu.CallbackData) {
	log.Info("menuFileCloseAppItemClicked")
	var hasChanges bool
	for _, file := range r.getFilesMap() {
		if (*file).HasChanges() {
			hasChanges = true
			break
		}
	}

	if !hasChanges {
		log.Info("menuFileCloseAppItemClicked, Doesn't have changes")
		runtime.Quit(*r.GetContext())
	}
	log.Info("menuFileCloseAppItemClicked, Have changes")

	dialogResult, err := r.dialogs.OkCancelMessageDialog(
		"flaskApplicationApi has files with changes",
		"You are trying to close application that have files with changes, continue?",
	)
	if err != nil {
		return
	}

	if dialogResult == "Ok" {
		log.Info("menuFileCloseAppItemClicked, Dialog result - OK")
		runtime.Quit(*r.GetContext())
	}
}
func (r *AppMenu) menuEditSortItemClicked(*menu.CallbackData) {
	log.Info("menuEditSortItemClicked")
}

func (r *AppMenu) menuFileGetFileInfoClicked(*menu.CallbackData) {
	openedFile := r.applicationApi.FindOpenedFile()
	var msg string
	if openedFile == nil {
		msg = "No opened files found"
	} else {
		fileInfo := *openedFile.GetInformation()
		filePath := fileInfo.GetPath()
		fileName := fileInfo.GetName()
		fileExt := fileInfo.GetExt()
		fileType := fileInfo.GetType()
		msg = fmt.Sprintf("File path: %s\nFile Name: %s\n File Extension: %s\n File Type: %s",
			filePath, fileName, fileExt, fileType)
	}

	err := r.dialogs.InfoMessageDialog("Information about current file", msg)
	if err != nil {
		log.Error(err)
		return
	}
}

func (r *AppMenu) menuEditFileInformationClicked(*menu.CallbackData) {
	r.SendEvent(constants.EventOnFileInformationChange)
}

func CreateApplicationMenu(contextRetriever *api.ContextRetriever, applicationApi *api.EditorApplication, dialogs *api.DialogsApi) api.ApplicationMenu {
	log.Info("CreateApplicationMenu", *contextRetriever, *applicationApi, *dialogs)
	return &AppMenu{
		contextRetriever: *contextRetriever,
		applicationApi:   *applicationApi,
		dialogs:          *dialogs,
	}
}
