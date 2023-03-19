package appmenu

import (
	"simple_text_editor/core/apperrors"
	"simple_text_editor/core/constants"
)

func sendErrorGenericMessage(receiver *AppMenu, errMsg string) {
	apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
		AppContext:  *receiver.GetContext(),
		Destination: constants.EventOnErrorHappened,
		Message:     errMsg,
	})
}

func sendErrorIO(receiver *AppMenu, err error) {
	apperrors.SendErrorEvent(&apperrors.ErrorNotificationStruct{
		AppContext:    *receiver.GetContext(),
		Destination:   constants.EventOnErrorHappened,
		OriginalError: err,
		Message:       "Read/Write error happened.",
	})
}
