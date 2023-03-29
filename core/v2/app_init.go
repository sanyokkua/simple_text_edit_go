package v2

import (
	"context"
	"simple_text_editor/core/v2/api"
	"simple_text_editor/core/v2/components/app"
	"simple_text_editor/core/v2/components/dialogs"
	"simple_text_editor/core/v2/components/events"
	"simple_text_editor/core/v2/components/frontend"
	"simple_text_editor/core/v2/components/menu"
	"simple_text_editor/core/v2/components/ops"
)

type appStruct struct {
	ctx      context.Context
	menuApi  api.IMenuApi
	frontApi api.IFrontendApi
}

func (r *appStruct) GetContext() context.Context {
	return r.ctx
}

func (r *appStruct) OnStartup(ctx context.Context) {
	r.ctx = ctx
}
func (r *appStruct) OnDomReady(context.Context) {

}
func (r *appStruct) OnShutdown(context.Context) {

}
func (r *appStruct) OnBeforeClose(context.Context) (prevent bool) {
	return false
}

func (r *appStruct) GetMenuApi() api.IMenuApi {
	return r.menuApi
}

func (r *appStruct) GetFrontendApi() api.IFrontendApi {
	return r.frontApi
}

func CreateApplication(extensions *map[string]api.FileTypesJsonStruct) api.IApplication {
	application := appStruct{}

	editor := app.CreateEditor(extensions)
	eventsApi := events.CreateEvents(application.GetContext)
	dialogsApi := dialogs.CreateDialogs(application.GetContext, extensions)
	filesOperations := ops.CreateFilesOperations(application.GetContext, editor, eventsApi, dialogsApi)
	menuApi := menu.CreateMenuApi(filesOperations)
	frontendApi := frontend.CreateFrontendApi(editor, eventsApi, dialogsApi, extensions)

	application.menuApi = menuApi
	application.frontApi = frontendApi

	return &application
}
