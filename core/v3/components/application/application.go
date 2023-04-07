package application

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/v3/components/dialoghelper"
	"simple_text_editor/core/v3/components/editor"
	"simple_text_editor/core/v3/components/eventsender"
	"simple_text_editor/core/v3/components/filehelper"
	"simple_text_editor/core/v3/components/frontendapi"
	"simple_text_editor/core/v3/components/menuhelper"
	"simple_text_editor/core/v3/components/menuopshelper"
	"simple_text_editor/core/v3/components/typemanager"
	"simple_text_editor/core/v3/components/uniqueidgen"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/validators"
)

type ApplicationStruct struct {
	Context     context.Context
	MenuHelper  types.IMenuHelper
	FrontendApi types.IFrontendApi
	iEditor     types.IEditor
}

func (r *ApplicationStruct) switchWinTitle() {
	file, err := r.iEditor.GetOpenedFile()
	if err != nil {
		return
	}
	if file.New {
		runtime.WindowSetTitle(r.Context, fmt.Sprintf("Simple Text Editor"))
	} else {
		runtime.WindowSetTitle(r.Context, fmt.Sprintf("Simple Text Editor. File Path: %s", file.Path))
	}

}
func (r *ApplicationStruct) GetContext() context.Context {
	return r.Context
}

func (r *ApplicationStruct) OnStartup(ctx context.Context) {
	r.Context = ctx
	runtime.EventsOn(ctx, string(eventsender.EventOnFileOpened), func(optionalData ...interface{}) {
		r.switchWinTitle()
	})
	runtime.EventsOn(ctx, string(eventsender.EventOnFileSaved), func(optionalData ...interface{}) {
		r.switchWinTitle()
	})
	runtime.EventsOn(ctx, string(eventsender.EventOnFileIsSwitched), func(optionalData ...interface{}) {
		r.switchWinTitle()
	})
	runtime.EventsOn(ctx, string(eventsender.EventOnNewFileCreated), func(optionalData ...interface{}) {
		r.switchWinTitle()
	})
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
	iMenuHelper := menuhelper.CreateIMenuHelper(iMenuOpsHelper, iTypeManager)
	frontendApi := frontendapi.CreateIFrontendApi(iEditor, iEventSender, iDialogHelper, iTypeManager)

	application.MenuHelper = iMenuHelper
	application.FrontendApi = frontendApi
	application.iEditor = iEditor

	return &application
}
