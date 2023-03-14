package application

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"simple_text_editor/handlers"
)

func (application *EditorApplication) menuFileNewItemClicked(data *menu.CallbackData) {}
func (application *EditorApplication) menuFileOpenItemClicked(data *menu.CallbackData) {
	handlers.OpenFileDialog(application.AppContext)
}
func (application *EditorApplication) menuFileSaveItemClicked(data *menu.CallbackData)      {}
func (application *EditorApplication) menuFileCloseFileItemClicked(data *menu.CallbackData) {}
func (application *EditorApplication) menuFileCloseAppItemClicked(data *menu.CallbackData)  {}
func (application *EditorApplication) menuEditRevertItemClicked(data *menu.CallbackData)    {}
func (application *EditorApplication) menuEditRepeatItemClicked(data *menu.CallbackData)    {}
func (application *EditorApplication) menuEditCutItemClicked(data *menu.CallbackData)       {}
func (application *EditorApplication) menuEditCopyItemClicked(data *menu.CallbackData)      {}
func (application *EditorApplication) menuEditPasteItemClicked(data *menu.CallbackData)     {}
func (application *EditorApplication) menuEditSelectAllItemClicked(data *menu.CallbackData) {}
func (application *EditorApplication) menuEditSortItemClicked(data *menu.CallbackData)      {}
