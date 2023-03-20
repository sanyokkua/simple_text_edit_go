package implementation

import (
	"context"
	"github.com/labstack/gommon/log"
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
	log.Info("(r *ApplicationContextHolderStruct) GetContext()", r.AppContext)
	return r.AppContext
}
func (r *ApplicationContextHolderStruct) OnStartup(ctx context.Context) {
	log.Info("(r *ApplicationContextHolderStruct) OnStartup(ctx context.Context)")
	r.AppContext = ctx
	app := *r.EditorApplication
	app.CreateEmptyFileAndMakeItOpened()
}
func (r *ApplicationContextHolderStruct) OnDomReady(context.Context) {
	log.Info("(r *ApplicationContextHolderStruct) OnDomReady(ctx context.Context)")
}
func (r *ApplicationContextHolderStruct) OnShutdown(context.Context) {
	log.Info("(r *ApplicationContextHolderStruct) OnShutdown(ctx context.Context)")
}
func (r *ApplicationContextHolderStruct) OnBeforeClose(context.Context) (prevent bool) {
	log.Info("(r *ApplicationContextHolderStruct) OnBeforeClose(ctx context.Context) (prevent bool)")
	return false
}
func (r *ApplicationContextHolderStruct) GetDialogsApi() api.DialogsApi {
	log.Info("(r *ApplicationContextHolderStruct) GetDialogsApi() api.DialogsApi", *r.DialogsApi)
	return *r.DialogsApi
}
func (r *ApplicationContextHolderStruct) GetApplicationMenu() api.ApplicationMenu {
	log.Info("(r *ApplicationContextHolderStruct) GetApplicationMenu() api.ApplicationMenu", *r.ApplicationMenu)
	return *r.ApplicationMenu
}
func (r *ApplicationContextHolderStruct) GetEditorApplicationApi() api.EditorApplication {
	log.Info("(r *ApplicationContextHolderStruct) GetEditorApplicationApi() api.EditorApplication", *r.EditorApplication)
	return *r.EditorApplication
}
func (r *ApplicationContextHolderStruct) GetJsApi() api.JsApi {
	log.Info("(r *ApplicationContextHolderStruct) GetJsApi() api.JsApi", *r.JsApi)
	return *r.JsApi
}

func CreateApplicationContextHolderApi() *api.ApplicationContextHolderApi {
	log.Info("CreateApplicationContextHolderApi() *api.ApplicationContextHolderApi")
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
	log.Info("CreateApplicationContextHolderApi() *api.ApplicationContextHolderApi, return", holder)
	return &holder
}
