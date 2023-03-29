package api

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

type IIdProvider interface {
	GetId() int64
}
type IEvents interface {
	SendEvent(destination Destination, optionalData ...interface{})
	SendErrorEvent(message string, optionalErrs ...error)
	SendWarnEvent(message string, optionalErrs ...error)
}
type IDialogs interface {
	OpenFileDialog() (filePath string, err error)
	SaveFileDialog(defaultFileNameWithExt string) (filePath string, err error)
	OkCancelMessageDialog(title string, message string) (Button, error)
	InfoMessageDialog(title string, message string) error
}
type IMenuApi interface {
	CreateMenu() *menu.Menu
}
type IFilesOperations interface {
	CreateNewFile()
	OpenFile()
	SaveCurrentFile()
	SaveCurrentFileAs()
	OpenCurrentFileInfo()
	EditCurrentFileInfo()
	CloseCurrentFile()
	CloseApplication()
}
type IFrontendApi interface {
	GetFilesShortInfo() []FileInfoStruct
	GetOpenedFile() FileStruct
	SwitchOpenedFileTo(fileId int64)
	UpdateFileContent(fileId int64, content string)
	UpdateFileInformation(fileId int64, information FileInfoUpdateStruct)
	GetFileTypes() []KeyValuePairStruct
	GetFileTypeExtension(fileTypeKey string) []KeyValuePairStruct
	CloseCurrentFile(fileId int64)
}

type IEditor interface {
	GetAllFilesInfo() []FileInfoStruct
	GetOpenedFile() (*FileStruct, error)
	CreateNewFileInEditor() error
	OpenFile(filePath string) error
	SaveFile(fileId int64) error
	CloseFile(fileId int64) error
	SwitchOpenedFileTo(fileId int64) error
	GetFileById(fileId int64) (*FileStruct, error)
	InactivateAllFiles()
	IsFileOpenedInEditor(filePath string) bool
	UpdateFileContent(fileId int64, content string) error
	UpdateFileInformation(fileId int64, information FileInfoUpdateStruct) error
}

type IApplication interface {
	GetContext() context.Context

	OnStartup(ctx context.Context)
	OnDomReady(ctx context.Context)
	OnShutdown(ctx context.Context)
	OnBeforeClose(ctx context.Context) (prevent bool)

	GetMenuApi() IMenuApi
	GetFrontendApi() IFrontendApi
}
