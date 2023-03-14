package handlers

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/constants"
	"simple_text_editor/logic"
)

func OpenFileDialog(ctx context.Context) {
	filePath, err := runtime.OpenFileDialog(ctx, runtime.OpenDialogOptions{
		Title: "Open Text File",
	})
	if err != nil {
		log.Warn(err)
		runtime.EventsEmit(ctx, constants.GENERIC_ERROR_HAPPENED, "Problem with opening dialog.")
		return
	}

	log.Info("File: ", filePath)

	text, err := logic.GetTextFromFile(filePath)
	if err != nil {
		log.Warn(err)
		runtime.EventsEmit(ctx, constants.GENERIC_ERROR_HAPPENED, "Problem with reading file.")
		return
	}

	runtime.EventsEmit(ctx, constants.EVENT_IS_FILE_OPENED, text)
}

func SaveFileDialog() {

}
