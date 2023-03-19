package dialogs

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/api"
	"simple_text_editor/core/apperrors"
	"simple_text_editor/core/constants"
)

type DialogStruct struct {
	contextRetriever api.ContextRetriever
}

func (r *DialogStruct) GetContext() *context.Context {
	ctx := r.contextRetriever()
	return &ctx
}

func (r *DialogStruct) OpenFileDialog() (filePath string, err error) {
	ctx := r.contextRetriever()

	filePath, err = runtime.OpenFileDialog(ctx, runtime.OpenDialogOptions{
		Title:   "Open Text OpenedFile",
		Filters: constants.GetSupportedFileFilters(),
	})
	return r.processDialogResults("OpenFileDialog", filePath, err)
}
func (r *DialogStruct) SaveFileDialog() (filePath string, err error) {
	ctx := r.contextRetriever()

	filePath, err = runtime.SaveFileDialog(ctx, runtime.SaveDialogOptions{
		Title:   "Save OpenedFile As...",
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

func CreateDialogApi(contextRetriever *api.ContextRetriever) api.DialogsApi {
	return &DialogStruct{
		contextRetriever: *contextRetriever,
	}
}
