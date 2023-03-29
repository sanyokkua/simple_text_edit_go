package events

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/v2/api"
)

type eventsStruct struct {
	getContext api.ContextProvider
}

func (r *eventsStruct) SendEvent(destination api.Destination, optionalData ...interface{}) {
	ctx := r.getContext()
	runtime.EventsEmit(ctx, string(destination), optionalData)
}
func (r *eventsStruct) SendErrorEvent(message string, optionalErrs ...error) {
	ctx := r.getContext()
	runtime.EventsEmit(ctx, string(api.EventOnInternalError), message, optionalErrs)
}
func (r *eventsStruct) SendWarnEvent(message string, optionalErrs ...error) {
	ctx := r.getContext()
	runtime.EventsEmit(ctx, string(api.EventOnInternalWarning), message, optionalErrs)
}

func CreateEvents(contextProvider api.ContextProvider) api.IEvents {
	if contextProvider == nil {
		panic("Context is not provided")
	}
	return &eventsStruct{
		getContext: contextProvider,
	}
}
