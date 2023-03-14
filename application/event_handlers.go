package application

import "github.com/wailsapp/wails/v2/pkg/runtime"

const (
	JsEventFileContentUpdated = "JS_EVENT_FILE_CONTENT_UPDATED"
)

func (application *EditorApplication) registerEventHandlers() {
	runtime.EventsOn(application.AppContext, JsEventFileContentUpdated, updateFileContent)
}
func updateFileContent(data ...interface{}) {}
