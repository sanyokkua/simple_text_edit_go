package eventsender

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/validators"
)

const (
	EventOnInternalWarning        types.Destination = "EventOnInternalWarning"
	EventOnInternalError          types.Destination = "EventOnInternalError"
	EventOnNewFileCreated         types.Destination = "EventOnNewFileCreated"
	EventOnFileOpened             types.Destination = "EventOnFileOpened"
	EventOnFileSaved              types.Destination = "EventOnFileSaved"
	EventOnFileClosed             types.Destination = "EventOnFileClosed"
	EventOnFileInformationUpdated types.Destination = "EventOnFileInformationUpdated"
	EventOnFileIsSwitched         types.Destination = "EventOnFileIsSwitched"
	EventOnFileContentIsUpdated   types.Destination = "EventOnFileContentIsUpdated"
	EventOnBlockUiTrue            types.Destination = "EventOnBlockUiTrue"
	EventOnBlockUiFalse           types.Destination = "EventOnBlockUiFalse"
)

type EventSenderStruct struct {
	GetContext types.ContextProvider
}

func (r *EventSenderStruct) SendNotificationEvent(destination types.Destination, optionalData ...interface{}) {
	ctx := r.GetContext()
	runtime.EventsEmit(ctx, string(destination), optionalData)
}

func (r *EventSenderStruct) SendErrorEvent(message string, optionalErrs ...error) {
	ctx := r.GetContext()
	runtime.EventsEmit(ctx, string(EventOnInternalError), message, optionalErrs)
}

func (r *EventSenderStruct) SendWarnEvent(message string, optionalErrs ...error) {
	ctx := r.GetContext()
	runtime.EventsEmit(ctx, string(EventOnInternalWarning), message, optionalErrs)
}

func CreateIEventSender(provider types.ContextProvider) types.IEventSender {
	validators.PanicOnNil(provider, "ContextProvider")

	return &EventSenderStruct{
		GetContext: provider,
	}
}
