package api

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

type InformationApi interface {
	GetOpenTimeStamp() int64
	GetPath() string
	GetName() string
	GetExt() string
	GetType() string
	GetExists() bool
	GetIsOpenedNow() bool
	GetIsChanged() bool

	SetOpenTimeStamp(timestamp int64)
	SetPath(path string)
	SetName(name string)
	SetExt(extension string)
	SetType(typeVal string)
	SetExists(exists bool)
	SetIsOpenedNow(opened bool)
	SetIsChanged(changed bool)
}

type FileApi interface {
	GetInformationRef() *InformationApi
	GetOriginalContent() string
	GetActualContent() string
	HasChanges() bool

	SetOriginalContent(content string)
	SetActualContent(content string)
}

type FlaskApplicationApi interface {
	OnStartup(ctx context.Context)
	OnDomReady(ctx context.Context)
	OnShutdown(ctx context.Context)
	OnBeforeClose(ctx context.Context) (prevent bool)

	GetContext() *context.Context
	GetFilesMap() map[int64]*FileApi
}

type EditorApplicationApi interface {
	CreateEmptyFile() *FileApi
	CreateExistingFile(filePath string, fileContent string) *FileApi

	AddEmptyFile(file *FileApi) *FileApi
	AddExistingFile(file *FileApi) *FileApi

	InactivateAllFiles()
	IsFileAlreadyOpened(filePath string) bool
	CloseFile(uniqueIdentifier int64) bool
	FindAnyFileInMemory() *FileApi

	GetFilesInformation() []InformationApi
	FindOpenedFile() FileApi
	ChangeFileStatusToOpened(uniqueIdentifier int64)
	ChangeFileContent(uniqueIdentifier int64, content string) bool
}

type DialogsApi interface {
	OpenFileDialog() (filePath string, err error)
	SaveFileDialog() (filePath string, err error)
	OkCancelMessageDialog(title string, message string) (clickedBtnName string, err error)
}

type ApplicationMenu interface {
	CreateMenu() *menu.Menu
	GetContext() *context.Context
	SendEvent(destination string, optionalData ...interface{})
}

type GetContext func() (ctx context.Context)

type ApplicationContextHolderApi interface {
	OnStartup(ctx context.Context)
	OnDomReady(ctx context.Context)
	OnShutdown(ctx context.Context)
	OnBeforeClose(ctx context.Context) (prevent bool)

	GetContext() (ctx context.Context)

	GetFlaskApplicationApi() FlaskApplicationApi
	GetDialogsApi() DialogsApi
	GetApplicationMenu() ApplicationMenu
	GetEditorApplicationApi() EditorApplicationApi
	GetJsApi() JsApi
}
