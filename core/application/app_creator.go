package application

import (
	"context"
	"simple_text_editor/core/api"
	"simple_text_editor/core/structures/appmenu"
	"simple_text_editor/core/structures/dialogs"
	"simple_text_editor/core/structures/editor"
	"simple_text_editor/core/structures/internals"
)

type ContextHolder struct {
	appContext           *context.Context
	flaskAppApi          *api.FlaskApplicationApi
	dialogsApi           *api.DialogsApi
	applicationMenu      *api.ApplicationMenu
	editorApplicationApi *api.EditorApplicationApi
	jsApi                *api.JsApi
}

func (receiver *ContextHolder) OnStartup(ctx context.Context) {
	receiver.appContext = &ctx
	flaskApp := *receiver.flaskAppApi
	flaskApp.OnStartup(ctx)
}
func (receiver *ContextHolder) OnDomReady(ctx context.Context) {
	flaskApp := *receiver.flaskAppApi
	flaskApp.OnDomReady(ctx)
}
func (receiver *ContextHolder) OnShutdown(ctx context.Context) {
	flaskApp := *receiver.flaskAppApi
	flaskApp.OnShutdown(ctx)
}
func (receiver *ContextHolder) OnBeforeClose(ctx context.Context) (prevent bool) {
	flaskApp := *receiver.flaskAppApi
	return flaskApp.OnBeforeClose(ctx)
}
func (receiver *ContextHolder) GetContext() (ctx context.Context) {
	return *receiver.appContext
}
func (receiver *ContextHolder) GetFlaskApplicationApi() api.FlaskApplicationApi {
	return *receiver.flaskAppApi
}
func (receiver *ContextHolder) GetDialogsApi() api.DialogsApi {
	return *receiver.dialogsApi
}
func (receiver *ContextHolder) GetApplicationMenu() api.ApplicationMenu {
	return *receiver.applicationMenu
}
func (receiver *ContextHolder) GetEditorApplicationApi() api.EditorApplicationApi {
	return *receiver.editorApplicationApi
}
func (receiver *ContextHolder) GetJsApi() api.JsApi {
	return *receiver.jsApi
}

func CreateApplicationContextHolderApi() api.ApplicationContextHolderApi {
	ctxHolder := &ContextHolder{}
	dialogsApi := dialogs.CreateDialogApi(ctxHolder.GetContext)
	flaskAppApi := editor.CreateFlaskApplicationApi(&dialogsApi)

	filesMap := (*flaskAppApi).GetFilesMap()

	editorApplicationApi := internals.CreateEditorInternals(&filesMap)
	applicationMenu := appmenu.CreateApplicationMenu(
		ctxHolder.GetContext,
		&filesMap,
		editorApplicationApi,
		dialogsApi,
	)

	var jsApi api.JsApi = editorApplicationApi
	ctxHolder.dialogsApi = &dialogsApi
	ctxHolder.flaskAppApi = flaskAppApi
	ctxHolder.editorApplicationApi = &editorApplicationApi
	ctxHolder.applicationMenu = &applicationMenu
	ctxHolder.jsApi = &jsApi

	createdFile := editorApplicationApi.CreateEmptyFile()
	editorApplicationApi.AddEmptyFile(createdFile)
	editorApplicationApi.ChangeFileStatusToOpened((*(*createdFile).GetInformationRef()).GetOpenTimeStamp())

	return ctxHolder
}
