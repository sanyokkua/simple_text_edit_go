package types

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type IUniqueIdGenerator interface {
	GenerateId() int64
}

type ITypeManager interface {
	GetTypeStructByKey(key FileTypeKey) (*FileTypesJsonStruct, error)
	GetTypeStructByExt(extension FileTypeExtension) (*FileTypesJsonStruct, error)

	GetTypeKeyByExtension(extension FileTypeExtension) (FileTypeKey, error)
	GetExtensionsForType(key FileTypeKey) ([]FileTypeExtension, error)

	GetSupportedFileFilters() []runtime.FileFilter

	BuildFileTypeMappingKeyToName() ([]KeyValuePairStruct, error)
	BuildFileTypeMappingExtToExt(fileTypeKey FileTypeKey) ([]KeyValuePairStruct, error)
}

type IFileHelper interface {
	GetFileNameFromPath(filePath string) (string, error)
	GetFileExtensionFromPath(filePath string) (FileTypeExtension, error)
	GetFileTypeFromExtension(fileExtension FileTypeExtension) (FileTypeKey, error)

	CreateNewFileEmpty() (*FileStruct, error)
	CreateNewFileWithData(path string, originalContent string) (*FileStruct, error)
}

type IEventSender interface {
	SendNotificationEvent(destination Destination, optionalData ...interface{})
	SendErrorEvent(message string, optionalErrs ...error)
	SendWarnEvent(message string, optionalErrs ...error)
}
type IDialogHelper interface {
	OpenFileDialog() (filePath string, err error)
	SaveFileDialog(defaultFileNameWithExt string) (filePath string, err error)
	OkCancelMessageDialog(title string, message string) (Button, error)
}
type IMenuHelper interface {
	CreateMenu() *menu.Menu
}
type IMenuOpsHelper interface {
	NewFile()
	OpenFile()
	SaveFile()
	SaveFileAs()
	ShowFileInfoModal()
	CloseFile()
	CloseApplication()
}
type IFrontendApi interface {
	GetFilesShortInfo() FrontendFileInfoArrayContainerStruct
	GetOpenedFile() FrontendFileContainerStruct
	GetFileTypes() FrontendKeyValueArrayContainerStruct
	GetFileTypeExtension(fileTypeKey string) FrontendKeyValueArrayContainerStruct

	SwitchOpenedFileTo(fileId int64)
	UpdateFileContent(fileId int64, content string)
	UpdateFileInformation(fileId int64, updateStruct FileTypeUpdateStruct)
	CloseFile(fileId int64)
}

type IEditor interface {
	GetFilesShortInfo() ([]FileInfoStruct, error)
	GetOpenedFile() (*FileStruct, error)

	CreateFileAndShow() error
	OpenFileAndShow(filePath string) error

	SaveFile(fileId int64) error
	CloseFile(fileId int64) error
	ShowFile(fileId int64) error

	GetFileById(fileId int64) (*FileStruct, error)

	HideAllFiles()

	UpdateFileContent(fileId int64, content string) error
	UpdateFileInformation(fileId int64, updateStruct FileTypeUpdateStruct) error
}

type IApplication interface {
	GetContext() context.Context

	OnStartup(ctx context.Context)
	OnDomReady(ctx context.Context)
	OnShutdown(ctx context.Context)
	OnBeforeClose(ctx context.Context) (prevent bool)

	GetMenuApi() IMenuHelper
	GetFrontendApi() IFrontendApi
}
