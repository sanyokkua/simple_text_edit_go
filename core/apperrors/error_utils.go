package apperrors

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ErrorNotificationStruct struct {
	AppContext    context.Context
	Destination   string
	Message       string
	OriginalError error
}

func CreateError(errorMessage string) error {
	log.Info("CreateError", errorMessage)
	return errors.New(errorMessage)
}

func WrapError(err error, errorMessage string) error {
	log.Info("WrapError", errorMessage, err)
	originalErrorMsg := err.Error()
	newError := CreateError(fmt.Sprintf("Error: %s. Original Error: %s", errorMessage, originalErrorMsg))
	return newError
}

func SendErrorEvent(notification *ErrorNotificationStruct) {
	log.Info("SendErrorEvent", *notification)
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
	log.Info("SendErrorEvent", msg)
	runtime.EventsEmit(notification.AppContext, notification.Destination, msg)
}
