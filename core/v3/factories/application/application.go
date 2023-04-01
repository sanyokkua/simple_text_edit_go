package application

import (
	"context"
	"simple_text_editor/core/v3/factories/dialoghelper"
	"simple_text_editor/core/v3/factories/editor"
	"simple_text_editor/core/v3/factories/eventsender"
	"simple_text_editor/core/v3/factories/filehelper"
	"simple_text_editor/core/v3/factories/frontendapi"
	"simple_text_editor/core/v3/factories/menuhelper"
	"simple_text_editor/core/v3/factories/menuopshelper"
	"simple_text_editor/core/v3/factories/typemanager"
	"simple_text_editor/core/v3/factories/uniqueidgen"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/validators"
)

type ApplicationStruct struct {
	Context     context.Context
	MenuHelper  types.IMenuHelper
	FrontendApi types.IFrontendApi
}

func (r *ApplicationStruct) GetContext() context.Context {
	return r.Context
}

func (r *ApplicationStruct) OnStartup(ctx context.Context) {
	r.Context = ctx
}

func (r *ApplicationStruct) OnDomReady(ctx context.Context) {
	//TODO implement me
}

func (r *ApplicationStruct) OnShutdown(ctx context.Context) {
	//TODO implement me
}

func (r *ApplicationStruct) OnBeforeClose(ctx context.Context) (prevent bool) {
	//TODO implement me
	return false
}

func (r *ApplicationStruct) GetMenuApi() types.IMenuHelper {
	return r.MenuHelper
}

func (r *ApplicationStruct) GetFrontendApi() types.IFrontendApi {
	return r.FrontendApi
}

func CreateIApplication(typesMap types.TypesMap) types.IApplication {
	validators.PanicOnNil(typesMap, "TypesMap")

	application := ApplicationStruct{}

	iUniqueIdGenerator := uniqueidgen.CreateIUniqueIdGenerator()
	iTypeManager := typemanager.CreateITypeManager(typesMap)
	iFileHelper := filehelper.CreateIFileHelper(iUniqueIdGenerator, iTypeManager)
	iEditor := editor.CreateIEditor(iFileHelper, iTypeManager)
	iEventSender := eventsender.CreateIEventSender(application.GetContext)
	iDialogHelper := dialoghelper.CreateIDialogHelper(application.GetContext, iTypeManager)
	iMenuOpsHelper := menuopshelper.CreateIMenuOpsHelper(application.GetContext, iEventSender, iDialogHelper, iEditor)
	iMenuHelper := menuhelper.CreateIMenuHelper(iMenuOpsHelper)
	frontendApi := frontendapi.CreateIFrontendApi(iEditor, iEventSender, iDialogHelper, iTypeManager)

	application.MenuHelper = iMenuHelper
	application.FrontendApi = frontendApi

	return &application
}
