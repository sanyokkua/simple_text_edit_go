package application

import (
	"errors"
	"github.com/labstack/gommon/log"
	"strconv"
	"time"
)

func (application *EditorApplication) GetAllOpenedFileDescriptors() []FileDescriptor {
	allOpenedFiles := make([]FileDescriptor, 0, len(application.AppFiles))

	for _, file := range application.AppFiles {
		descriptor := file.Descriptor
		allOpenedFiles = append(allOpenedFiles, descriptor)
	}
	log.Info("GetAllOpenedFileDescriptors", allOpenedFiles)
	return allOpenedFiles
}

func (application *EditorApplication) GetAllOpenedFile() map[int64]*AppFile {
	log.Info("GetAllOpenedFile", application.AppFiles)
	return application.AppFiles
}

func (application *EditorApplication) GetActiveFile() (*AppFile, error) {
	for _, file := range application.AppFiles {
		if file.Descriptor.IsActive {
			log.Info("GetActiveFile", file)
			return file, nil
		}
	}
	log.Warn("GetActiveFile", "Not FOUND. Error")
	return &AppFile{}, errors.New("no active files")
}

func (application *EditorApplication) MakeFileActive(fileId int64) error {
	openedFiles := application.GetAllOpenedFile()
	file, ok := openedFiles[fileId]

	if !ok {
		return errors.New("passed file path is not found in opened files map. path:" + strconv.FormatInt(fileId, 10))
	}

	inactivateAllFiles(application)
	file.Descriptor.IsActive = true
	log.Info("MakeFileActive", file)
	return nil
}

func (application *EditorApplication) UpdateFileContent(fileId int64, fileContent string) error {
	log.Info("UpdateFileContent", fileId)
	openedFiles := application.GetAllOpenedFile()
	file := openedFiles[fileId]

	if !file.Descriptor.IsActive {
		return errors.New("there was a try to update file that is not active")
	}

	log.Info("UpdateFileContent", file)
	changeTime := time.Now().UnixNano()
	file.ContentHistory[changeTime] = file.FileContent
	file.FileContent = fileContent
	log.Info("UpdateFileContent", changeTime)
	return nil
}
