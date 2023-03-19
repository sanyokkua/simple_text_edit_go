package api

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

type ContextRetriever func() (ctx context.Context)
type FileInformation interface {
	GetOpenTimeStamp() int64
	GetPath() string
	GetName() string
	GetExt() string
	GetType() string
	Exists() bool
	IsOpened() bool

	SetOpenTimeStamp(timestamp int64)
	SetPath(path string)
	SetName(name string)
	SetExt(extension string)
	SetType(typeVal string)
	SetExists(exists bool)
	SetIsOpened(opened bool)
}
type OpenedFile interface {
	GetInformation() *FileInformation
	GetOriginalContent() string
	GetActualContent() string
	HasChanges() bool

	SetOriginalContent(content string)
	SetActualContent(content string)
}
type EditorApplication interface {
	GetFilesMap() *map[int64]*OpenedFile // Required for JS API
	InactivateAllFiles()                 // Required for JS API
	ChangeFileStatusToOpened(uniqueIdentifier int64)
	FindOpenedFile() OpenedFile

	CreateEmptyFileAndMakeItOpened()
	AddFileToMemory(file *OpenedFile) *OpenedFile
	IsFileAlreadyOpened(filePath string) bool

	CloseFile(uniqueIdentifier int64) bool

	FindAnyFileInMemory() *OpenedFile
}

type DialogsApi interface {
	GetContext() *context.Context
	OpenFileDialog() (filePath string, err error)
	SaveFileDialog() (filePath string, err error)
	OkCancelMessageDialog(title string, message string) (clickedBtnName string, err error)
}

type ApplicationMenu interface {
	GetContext() *context.Context
	CreateMenu() *menu.Menu
	SendEvent(destination string, optionalData ...interface{})
}

type ApplicationContextHolderApi interface {
	GetContext() context.Context

	OnStartup(ctx context.Context)
	OnDomReady(ctx context.Context)
	OnShutdown(ctx context.Context)
	OnBeforeClose(ctx context.Context) (prevent bool)

	GetDialogsApi() DialogsApi
	GetApplicationMenu() ApplicationMenu
	GetEditorApplicationApi() EditorApplication
	GetJsApi() JsApi
}
