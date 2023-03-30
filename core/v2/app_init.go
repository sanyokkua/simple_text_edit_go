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
	"simple_text_editor/core/v2/components/typemngr"
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

func CreateApplication(extensions []api.FileTypesJsonStruct) api.IApplication {
	application := appStruct{}

	typeManager := typemngr.CreateTypeManager(extensions)
	editor := app.CreateEditor(typeManager)
	eventsApi := events.CreateEvents(application.GetContext)
	dialogsApi := dialogs.CreateDialogs(application.GetContext, typeManager)
	filesOperations := ops.CreateFilesOperations(application.GetContext, editor, eventsApi, dialogsApi)
	menuApi := menu.CreateMenuApi(filesOperations)
	frontendApi := frontend.CreateFrontendApi(editor, eventsApi, dialogsApi, typeManager)

	application.menuApi = menuApi
	application.frontApi = frontendApi

	return &application
}
