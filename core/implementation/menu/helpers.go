package menu

import (
	"github.com/labstack/gommon/log"
	"simple_text_editor/core/apperrors"
	"simple_text_editor/core/constants"
)

func sendErrorGenericMessage(receiver *AppMenu, errMsg string) {
	log.Info("sendErrorGenericMessage", errMsg)
	apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
		AppContext:  *receiver.GetContext(),
		Destination: constants.EventOnErrorHappened,
		Message:     errMsg,
	})
}

func sendErrorIO(receiver *AppMenu, err error) {
	log.Info("sendErrorIO", err)
	apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
		AppContext:    *receiver.GetContext(),
		Destination:   constants.EventOnErrorHappened,
		OriginalError: err,
		Message:       "Read/Write error happened.",
	})
}
