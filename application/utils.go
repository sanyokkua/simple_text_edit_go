package application

import (
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"path/filepath"
	"simple_text_editor/constants"
	"time"
)

func createAppFileForNewFile() *AppFile {
	fileId := time.Now().UnixNano()

	openedFileStruct := &AppFile{
		Descriptor: FileDescriptor{
			FileId:   fileId,
			FileName: "New*",
			FilePath: "",
			FileType: "txt",
			IsActive: true,
		},
		FileContent:     "",
		SelectedContent: "",
		ContentHistory:  make(map[int64]string),
	}

	return openedFileStruct
}

func createAppFileForExistingFile(filePath string, fileContent string) *AppFile {
	fileId := time.Now().UnixNano()
	name := getFileNameFromPath(filePath)
	extension := getFileExtensionFromPath(filePath)

	openedFileStruct := &AppFile{
		Descriptor: FileDescriptor{
			FileId:   fileId,
			FileName: name,
			FileType: extension,
			FilePath: filePath,
			IsActive: true,
		},
		FileContent:     fileContent,
		SelectedContent: "",
		ContentHistory:  make(map[int64]string),
	}

	return openedFileStruct
}

func getFileNameFromPath(filePath string) string {
	if len(filePath) == 0 {
		return filePath
	}
	fileName := filepath.Base(filePath)
	return fileName
}

func getFileExtensionFromPath(filePath string) string {
	if len(filePath) == 0 {
		return filePath
	}
	fileName := filepath.Ext(filePath)
	return fileName
}

func inactivateAllFiles(app *EditorApplication) {
	log.Info("inactivateAllFiles")
	for _, openedFile := range app.AppFiles {
		openedFile.Descriptor.IsActive = false
		log.Info("inactivateAllFiles", openedFile)
	}
	log.Info("inactivateAllFiles result", app.AppFiles)
}

func processGenericError(app *EditorApplication, err error, message string) {
	log.Warn(err)
	runtime.EventsEmit(app.AppContext, constants.GENERIC_ERROR_HAPPENED, message)
}
