package application

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

type EditorApplication struct {
	AppContext  context.Context
	OpenedFiles map[string]OpenedFile
}

type OpenedFile struct {
	fileName     string
	fileContent  string
	selectedText string
}

func CreateNewApplication() *EditorApplication {
	return &EditorApplication{
		OpenedFiles: make(map[string]OpenedFile),
	}
}

func (application *EditorApplication) Startup(ctx context.Context) {
	application.AppContext = ctx
	application.registerEventHandlers()
}

func (application *EditorApplication) CreateMenu() *menu.Menu {
	appMenu := menu.NewMenu()
	createFileSubmenu(appMenu, application)
	createEditSubmenu(appMenu, application)
	return appMenu
}
func createFileSubmenu(mainMenu *menu.Menu, application *EditorApplication) {
	file := mainMenu.AddSubmenu("File")
	file.AddText("New", keys.CmdOrCtrl("N"), application.menuFileNewItemClicked)
	file.AddText("Open", keys.CmdOrCtrl("O"), application.menuFileOpenItemClicked)
	file.AddText("Save", keys.CmdOrCtrl("S"), application.menuFileSaveItemClicked)
	file.AddText("Close File", nil, application.menuFileCloseFileItemClicked)
	file.AddText("Close Application", keys.CmdOrCtrl("Q"), application.menuFileCloseAppItemClicked)
}
func createEditSubmenu(mainMenu *menu.Menu, application *EditorApplication) {
	file := mainMenu.AddSubmenu("Edit")
	file.AddText("Revert", keys.CmdOrCtrl("Z"), application.menuEditRevertItemClicked)
	file.AddText("Repeat", nil, application.menuEditRepeatItemClicked)
	file.AddSeparator()
	file.AddText("Cut", keys.CmdOrCtrl("X"), application.menuEditCutItemClicked)
	file.AddText("Copy", keys.CmdOrCtrl("C"), application.menuEditCopyItemClicked)
	file.AddText("Paste", keys.CmdOrCtrl("V"), application.menuEditPasteItemClicked)
	file.AddText("Select ALL", keys.CmdOrCtrl("A"), application.menuEditSelectAllItemClicked)
	file.AddSeparator()
	file.AddText("Sort", keys.CmdOrCtrl("L"), application.menuEditSortItemClicked)
}
