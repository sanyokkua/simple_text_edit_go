package menuhelper

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/validators"
)

type MenuHelperStruct struct {
	MenuOpsHelper types.IMenuOpsHelper
}

func (r *MenuHelperStruct) CreateMenu() *menu.Menu {
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

func (r *MenuHelperStruct) menuFileNew(*menu.CallbackData) {
	r.MenuOpsHelper.NewFile()
}

func (r *MenuHelperStruct) menuFileOpen(*menu.CallbackData) {
	r.MenuOpsHelper.OpenFile()
}

func (r *MenuHelperStruct) menuFileSave(*menu.CallbackData) {
	r.MenuOpsHelper.SaveFile()
}

func (r *MenuHelperStruct) menuFileSaveAs(*menu.CallbackData) {
	r.MenuOpsHelper.SaveFileAs()
}

func (r *MenuHelperStruct) menuFileCloseFile(*menu.CallbackData) {
	r.MenuOpsHelper.CloseFile()
}

func (r *MenuHelperStruct) menuFileCloseApp(*menu.CallbackData) {
	r.MenuOpsHelper.CloseApplication()
}

func (r *MenuHelperStruct) menuEditFileInfo(*menu.CallbackData) {
	r.MenuOpsHelper.ShowFileInfoModal()
}

func (r *MenuHelperStruct) menuEditSort(*menu.CallbackData) {
	// TODO: add later
}

func CreateIMenuHelper(menuOpsHelper types.IMenuOpsHelper) types.IMenuHelper {
	validators.PanicOnNil(menuOpsHelper, "IMenuOpsHelper")

	return &MenuHelperStruct{
		MenuOpsHelper: menuOpsHelper,
	}
}
