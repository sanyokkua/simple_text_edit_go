package dialogs

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/api"
	"simple_text_editor/core/apperrors"
	"simple_text_editor/core/constants"
)

type DialogStruct struct {
	contextRetriever api.ContextRetriever
}

func (r *DialogStruct) GetContext() *context.Context {
	log.Info("(r *DialogStruct) GetContext()")
	ctx := r.contextRetriever()
	log.Info("(r *DialogStruct) GetContext(), returns", ctx)
	return &ctx
}

func (r *DialogStruct) OpenFileDialog() (filePath string, err error) {
	log.Info("OpenFileDialog")
	ctx := r.contextRetriever()

	filePath, err = runtime.OpenFileDialog(ctx, runtime.OpenDialogOptions{
		Title:   "Open Text OpenedFile",
		Filters: constants.GetSupportedFileFilters(),
	})
	log.Info("OpenFileDialog", filePath, err)
	return r.processDialogResults("OpenFileDialog", filePath, err)
}
func (r *DialogStruct) SaveFileDialog() (filePath string, err error) {
	log.Info("SaveFileDialog")
	ctx := r.contextRetriever()

	filePath, err = runtime.SaveFileDialog(ctx, runtime.SaveDialogOptions{
		Title: "Save OpenedFile As...",
		//Filters:         constants.GetSupportedFileFilters(),
		DefaultFilename: "NewFile.txt",
	})
	log.Info("SaveFileDialog", filePath, err)
	return r.processDialogResults("SaveFileDialog", filePath, err)
}
func (r *DialogStruct) OkCancelMessageDialog(title string, message string) (clickedBtnName string, err error) {
	log.Info("OkCancelMessageDialog", title, message)
	ctx := r.contextRetriever()

	clickedBtnName, err = runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:          runtime.WarningDialog,
		Title:         title,
		Message:       message,
		Buttons:       []string{"Ok", "Cancel"},
		DefaultButton: "Ok",
		CancelButton:  "Cancel",
	})
	log.Info("OkCancelMessageDialog", clickedBtnName, err)
	return r.processDialogResults("MessageDialog", clickedBtnName, err)
}
func (r *DialogStruct) processDialogResults(dialogName string, result string, err error) (string, error) {
	log.Info("processDialogResults", dialogName, result, err)
	if err != nil {
		r.sendErrorDialogReturnedError(dialogName, err)
		return "", err
	}
	return result, nil
}
func (r *DialogStruct) sendErrorDialogReturnedError(dialogName string, err error) {
	log.Info("sendErrorDialogReturnedError", dialogName, err)
	ctx := r.contextRetriever()
	msg := fmt.Sprintf("Problem happened with openening %s", dialogName)
	apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
		AppContext:    ctx,
		Destination:   constants.EventOnErrorHappened,
		OriginalError: err,
		Message:       msg,
	})
}

func (r *DialogStruct) InfoMessageDialog(title string, message string) (err error) {
	log.Info("InfoMessageDialog", title, message)
	ctx := r.contextRetriever()

	_, err = runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:          runtime.WarningDialog,
		Title:         title,
		Message:       message,
		Buttons:       []string{"Ok"},
		DefaultButton: "Ok",
		CancelButton:  "Ok",
	})
	log.Info("InfoMessageDialog", err)
	_, err = r.processDialogResults("MessageDialog", "", err)
	return err
}

func CreateDialogApi(contextRetriever *api.ContextRetriever) api.DialogsApi {
	log.Info("CreateDialogApi", *contextRetriever)
	return &DialogStruct{
		contextRetriever: *contextRetriever,
	}
}
