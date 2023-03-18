package apperrors

import (
	"context"
	"errors"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ErrorNotificationStruct struct {
	AppContext    context.Context
	Destination   string
	Message       string
	OriginalError error
}

func CreateError(errorMessage string) error {
	return errors.New(errorMessage)
}

func WrapError(err error, errorMessage string) error {
	originalErrorMsg := err.Error()
	newError := CreateError(fmt.Sprintf("Error: %s. Original Error: %s", errorMessage, originalErrorMsg))
	return newError
}

func SendErrorEvent(notification *ErrorNotificationStruct) {
	if notification == nil || len(notification.Destination) == 0 || notification.AppContext == nil {
		return
	}
	var err error
	if len(notification.Message) > 0 && notification.OriginalError != nil {
		err = WrapError(notification.OriginalError, notification.Message)
	} else if len(notification.Message) > 0 {
		err = CreateError(notification.Message)
	} else if notification.OriginalError != nil {
		err = notification.OriginalError
	} else {
		err = CreateError("Unknown error")
	}
	msg := err.Error()
	println(msg)
	runtime.EventsEmit(notification.AppContext, notification.Destination, msg)
}
