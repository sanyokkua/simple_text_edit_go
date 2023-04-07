package dialoghelper

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/validators"
)

type DialogHelperStruct struct {
	GetContext  types.ContextProvider
	TypeManager types.ITypeManager
}

func (r *DialogHelperStruct) OpenFileDialog() (filePath string, err error) {
	ctx := r.GetContext()

	filePath, err = runtime.OpenFileDialog(ctx, runtime.OpenDialogOptions{
		Title:   "Open File",
		Filters: r.TypeManager.GetSupportedFileFilters(), //TODO:
	})

	return filePath, err
}

func (r *DialogHelperStruct) SaveFileDialog(defaultFileNameWithExt string) (filePath string, err error) {
	ctx := r.GetContext()

	filePath, err = runtime.SaveFileDialog(ctx, runtime.SaveDialogOptions{
		Title:           "Save File As...",
		DefaultFilename: defaultFileNameWithExt,
	})

	return filePath, err
}

func (r *DialogHelperStruct) OkCancelMessageDialog(title string, message string) (types.Button, error) {
	ctx := r.GetContext()
	okBtn := string(types.BtnOk)
	cancelBtn := string(types.BtnCancel)

	clickedBtnName, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:          runtime.WarningDialog,
		Title:         title,
		Message:       message,
		Buttons:       []string{okBtn, cancelBtn},
		DefaultButton: okBtn,
		CancelButton:  cancelBtn,
	})

	return types.Button(clickedBtnName), err
}

func CreateIDialogHelper(provider types.ContextProvider, typeManager types.ITypeManager) types.IDialogHelper {
	validators.PanicOnNil(provider, "ContextProvider")
	validators.PanicOnNil(typeManager, "TypeManager")

	return &DialogHelperStruct{
		GetContext:  provider,
		TypeManager: typeManager,
	}
}
