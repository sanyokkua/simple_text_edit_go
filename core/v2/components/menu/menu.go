package menu

import (
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"simple_text_editor/core/v2/api"
)

type menuStruct struct {
	fileOps api.IFilesOperations
}

func (r *menuStruct) CreateMenu() *menu.Menu {
	log.Info("CreateMenu")
	appMenu := menu.NewMenu()

	file := appMenu.AddSubmenu("OpenedFile")
	file.AddText("New", keys.CmdOrCtrl("N"), r.menuFileNew)
	file.AddText("Open", keys.CmdOrCtrl("O"), r.menuFileOpen)
	file.AddText("Save", keys.CmdOrCtrl("S"), r.menuFileSave)
	file.AddText("Save as", nil, r.menuFileSaveAs)
	file.AddText("Close File", nil, r.menuFileCloseFile)
	file.AddText("Close Application", keys.CmdOrCtrl("Q"), r.menuFileCloseApp)

	edit := appMenu.AddSubmenu("Edit")
	edit.AddText("Edit File Information", nil, r.menuEditFileInfo)
	edit.AddText("Sort", keys.CmdOrCtrl("L"), r.menuEditSort)

	return appMenu
}

func (r *menuStruct) menuFileNew(*menu.CallbackData) {
	r.fileOps.CreateNewFile()
}

func (r *menuStruct) menuFileOpen(*menu.CallbackData) {
	r.fileOps.OpenFile()
}

func (r *menuStruct) menuFileSave(*menu.CallbackData) {
	r.fileOps.SaveCurrentFile()
}

func (r *menuStruct) menuFileSaveAs(*menu.CallbackData) {
	r.fileOps.SaveCurrentFileAs()
}

func (r *menuStruct) menuFileCloseFile(*menu.CallbackData) {
	r.fileOps.CloseCurrentFile()
}

func (r *menuStruct) menuFileCloseApp(*menu.CallbackData) {
	r.fileOps.CloseApplication()
}

func (r *menuStruct) menuEditFileInfo(*menu.CallbackData) {
	r.fileOps.EditCurrentFileInfo()
}

func (r *menuStruct) menuEditSort(*menu.CallbackData) {
	// TODO: add later
}

func CreateMenuApi(filesOps api.IFilesOperations) api.IMenuApi {
	if filesOps == nil {
		panic("CreateMenuApi failed due to nil filesOps api.IFilesOperations")
	}
	menuApi := menuStruct{
		fileOps: filesOps,
	}
	return &menuApi
}
