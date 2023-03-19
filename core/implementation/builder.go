package implementation

import (
	"context"
	"simple_text_editor/core/api"
	"simple_text_editor/core/implementation/dialogs"
	"simple_text_editor/core/implementation/editor"
	"simple_text_editor/core/implementation/jsapi"
	"simple_text_editor/core/implementation/menu"
)

type ApplicationContextHolderStruct struct {
	AppContext        context.Context
	EditorApplication *api.EditorApplication
	DialogsApi        *api.DialogsApi
	ApplicationMenu   *api.ApplicationMenu
	JsApi             *api.JsApi
}

func (r *ApplicationContextHolderStruct) GetContext() context.Context {
	return r.AppContext
}
func (r *ApplicationContextHolderStruct) OnStartup(ctx context.Context) {
	r.AppContext = ctx
	app := *r.EditorApplication
	app.CreateEmptyFileAndMakeItOpened()
}
func (r *ApplicationContextHolderStruct) OnDomReady(ctx context.Context) {

}
func (r *ApplicationContextHolderStruct) OnShutdown(ctx context.Context) {

}
func (r *ApplicationContextHolderStruct) OnBeforeClose(ctx context.Context) (prevent bool) {
	return false
}
func (r *ApplicationContextHolderStruct) GetDialogsApi() api.DialogsApi {
	return *r.DialogsApi
}
func (r *ApplicationContextHolderStruct) GetApplicationMenu() api.ApplicationMenu {
	return *r.ApplicationMenu
}
func (r *ApplicationContextHolderStruct) GetEditorApplicationApi() api.EditorApplication {
	return *r.EditorApplication
}
func (r *ApplicationContextHolderStruct) GetJsApi() api.JsApi {
	return *r.JsApi
}

func CreateApplicationContextHolderApi() *api.ApplicationContextHolderApi {
	appContextHolder := ApplicationContextHolderStruct{
		EditorApplication: nil,
		DialogsApi:        nil,
		ApplicationMenu:   nil,
		JsApi:             nil,
	}

	var retriever api.ContextRetriever
	retriever = appContextHolder.GetContext

	dialogApi := dialogs.CreateDialogApi(&retriever)
	app := editor.CreateEditorApplication(&retriever, &dialogApi)
	jsApi := jsapi.CreateJsApi(&app)
	appMenu := menu.CreateApplicationMenu(&retriever, &app, &dialogApi)

	appContextHolder.DialogsApi = &dialogApi
	appContextHolder.EditorApplication = &app
	appContextHolder.JsApi = &jsApi
	appContextHolder.ApplicationMenu = &appMenu

	var holder api.ApplicationContextHolderApi
	holder = &appContextHolder
	return &holder
}
