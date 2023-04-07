package menuhelper

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/validators"
)

type MenuHelperStruct struct {
	MenuOpsHelper types.IMenuOpsHelper
	TypeManager   types.ITypeManager
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
	edit.AddText("Sort", keys.CmdOrCtrl("L"), r.menuEditSort)
	changeType := edit.AddSubmenu("Change File Type")

	err := r.CreateFileTypeSubmenu(changeType)
	if err != nil {
		return nil
	}

	debug := appMenu.AddSubmenu("debug")
	debug.AddText("Block UI", nil, func(data *menu.CallbackData) {
		r.MenuOpsHelper.BlockUI(true)
	})
	debug.AddText("Unblock UI", nil, func(data *menu.CallbackData) {
		r.MenuOpsHelper.BlockUI(false)
	})

	return appMenu
}

func (r *MenuHelperStruct) CreateFileTypeSubmenu(topMenu *menu.Menu) error {
	typeMappings, getMappingErr := r.TypeManager.BuildFileTypeMappingKeyToName()
	if getMappingErr != nil {
		return getMappingErr
	}

	for _, typeMapping := range typeMappings {
		extensions, extErr := r.TypeManager.GetExtensionsForType(types.FileTypeKey(typeMapping.Key))
		if extErr != nil {
			return extErr
		}

		header := typeMapping.Value
		submenu := topMenu.AddSubmenu(header)

		for _, extension := range extensions {
			submenu.AddText(string(extension), nil, func(data *menu.CallbackData) {
				fType := types.FileTypeKey(typeMapping.Key)
				fExt := extension
				r.MenuOpsHelper.ChangeExtension(fType, fExt)
			})
		}
	}

	return nil
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

func (r *MenuHelperStruct) menuEditSort(*menu.CallbackData) {
	// TODO: add later
}

func CreateIMenuHelper(menuOpsHelper types.IMenuOpsHelper, typeManager types.ITypeManager) types.IMenuHelper {
	validators.PanicOnNil(menuOpsHelper, "IMenuOpsHelper")
	validators.PanicOnNil(typeManager, "ITypeManager")

	return &MenuHelperStruct{
		MenuOpsHelper: menuOpsHelper,
		TypeManager:   typeManager,
	}
}
