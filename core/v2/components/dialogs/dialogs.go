package dialogs

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/v2/api"
	"simple_text_editor/core/v2/utils"
)

type dialogsStruct struct {
	getContext api.ContextProvider
	extensions map[string]api.FileTypesJsonStruct
}

func (r *dialogsStruct) OpenFileDialog() (filePath string, err error) {
	ctx := r.getContext()

	filePath, err = runtime.OpenFileDialog(ctx, runtime.OpenDialogOptions{
		Title:   "Open File",
		Filters: utils.GetSupportedFileFilters(r.extensions),
	})

	return filePath, err
}
func (r *dialogsStruct) SaveFileDialog(defaultFileNameWithExt string) (filePath string, err error) {
	ctx := r.getContext()

	filePath, err = runtime.SaveFileDialog(ctx, runtime.SaveDialogOptions{
		Title:           "Save File As...",
		DefaultFilename: defaultFileNameWithExt,
	})

	return filePath, err
}
func (r *dialogsStruct) OkCancelMessageDialog(title string, message string) (api.Button, error) {
	ctx := r.getContext()
	okBtn := string(api.BtnOk)
	cancelBtn := string(api.BtnCancel)

	clickedBtnName, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:          runtime.WarningDialog,
		Title:         title,
		Message:       message,
		Buttons:       []string{okBtn, cancelBtn},
		DefaultButton: okBtn,
		CancelButton:  cancelBtn,
	})

	return api.Button(clickedBtnName), err
}
func (r *dialogsStruct) InfoMessageDialog(title string, message string) error {
	ctx := r.getContext()
	okBtn := string(api.BtnOk)

	_, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:          runtime.WarningDialog,
		Title:         title,
		Message:       message,
		Buttons:       []string{okBtn},
		DefaultButton: okBtn,
		CancelButton:  okBtn,
	})

	return err
}

func CreateDialogs(ctx api.ContextProvider, extensions *map[string]api.FileTypesJsonStruct) api.IDialogs {
	if ctx == nil {
		panic("Create IDialogs failed because ctx is nil")
	}
	return &dialogsStruct{
		getContext: ctx,
		extensions: *extensions,
	}
}
