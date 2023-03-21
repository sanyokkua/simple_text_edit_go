package api

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"simple_text_editor/core/constants"
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
	HasChanges() bool

	SetOpenTimeStamp(timestamp int64)
	SetPath(path string)
	SetName(name string)
	SetExt(extension string)
	SetType(typeVal string)
	SetExists(exists bool)
	SetIsOpened(opened bool)
	SetHasChanges(changed bool)
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
	SaveFileDialog(defaultFileName string) (filePath string, err error)
	OkCancelMessageDialog(title string, message string) (clickedBtnName string, err error)
	InfoMessageDialog(title string, message string) (err error)
}
type ApplicationMenu interface {
	GetContext() *context.Context
	CreateMenu() *menu.Menu
	SendEvent(destination string, optionalData ...interface{})
}
type JsApi interface {
	GetFilesInformation() []FileInformation
	FindOpenedFile() OpenedFile
	ChangeFileStatusToOpened(uniqueIdentifier int64)
	ChangeFileContent(uniqueIdentifier int64, content string) bool
	ChangeFileInformation(dialResults DialogResult)
	GetFileTypeInformation() []constants.FileTypeInformation
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

type DialogResult struct {
	FileName string
	FileType string
	FileExt  string
}
