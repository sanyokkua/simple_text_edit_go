package dialogs

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/api"
	"simple_text_editor/core/apperrors"
	"simple_text_editor/core/constants"
)

type DialogStruct struct {
	contextRetriever api.GetContext
}

func (r *DialogStruct) OpenFileDialog() (filePath string, err error) {
	ctx := r.contextRetriever()

	filePath, err = runtime.OpenFileDialog(ctx, runtime.OpenDialogOptions{
		Title:   "Open Text FileApi",
		Filters: constants.GetSupportedFileFilters(),
	})
	return r.processDialogResults("OpenFileDialog", filePath, err)
}
func (r *DialogStruct) SaveFileDialog() (filePath string, err error) {
	ctx := r.contextRetriever()

	filePath, err = runtime.SaveFileDialog(ctx, runtime.SaveDialogOptions{
		Title:   "Save FileApi As...",
		Filters: constants.GetSupportedFileFilters(),
	})
	return r.processDialogResults("SaveFileDialog", filePath, err)
}
func (r *DialogStruct) OkCancelMessageDialog(title string, message string) (clickedBtnName string, err error) {
	ctx := r.contextRetriever()

	clickedBtnName, err = runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:          runtime.WarningDialog,
		Title:         title,
		Message:       message,
		Buttons:       []string{"Ok", "Cancel"},
		DefaultButton: "Ok",
		CancelButton:  "Cancel",
	})
	return r.processDialogResults("MessageDialog", clickedBtnName, err)
}
func (r *DialogStruct) processDialogResults(dialogName string, result string, err error) (string, error) {
	if err != nil {
		r.sendErrorDialogReturnedError(dialogName, err)
		return "", err
	}
	return result, nil
}
func (r *DialogStruct) sendErrorDialogReturnedError(dialogName string, err error) {
	ctx := r.contextRetriever()
	msg := fmt.Sprintf("Problem happened with openening %s", dialogName)
	apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
		AppContext:    ctx,
		Destination:   constants.EventOnErrorHappened,
		OriginalError: err,
		Message:       msg,
	})
}

func CreateDialogApi(contextRetriever api.GetContext) api.DialogsApi {
	return &DialogStruct{
		contextRetriever: contextRetriever,
	}
}
