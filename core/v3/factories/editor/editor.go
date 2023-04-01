package editor

import (
	"errors"
	"fmt"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/utils/io"
	"simple_text_editor/core/v3/validators"
)

type EditorStruct struct {
	Files       types.FilesMap
	FileHelper  types.IFileHelper
	TypeManager types.ITypeManager
}

func mapFileStructToInfoStruct(fileStruct *types.FileStruct) types.FileInfoStruct {
	return types.FileInfoStruct{
		Id:        fileStruct.Id,
		Path:      fileStruct.Path,
		Name:      fileStruct.Name,
		Type:      fileStruct.Type,
		Extension: fileStruct.Extension,
		New:       fileStruct.New,
		Opened:    fileStruct.Opened,
		Changed:   fileStruct.Changed,
	}
}

func (r *EditorStruct) GetFilesShortInfo() ([]types.FileInfoStruct, error) {
	filesInfoSlice := make([]types.FileInfoStruct, 0, len(r.Files))

	for _, fileStruct := range r.Files {
		infoStruct := mapFileStructToInfoStruct(fileStruct)
		filesInfoSlice = append(filesInfoSlice, infoStruct)
	}

	return filesInfoSlice, nil
}

func (r *EditorStruct) GetOpenedFile() (*types.FileStruct, error) {
	for _, fileStruct := range r.Files {
		if fileStruct.Opened {
			return fileStruct, nil
		}
	}
	return nil, errors.New("opened file is not found")
}

func (r *EditorStruct) CreateFileAndShow() error {
	newFileEmpty, fileCreationErr := r.FileHelper.CreateNewFileEmpty()
	if validators.HasError(fileCreationErr) {
		return fileCreationErr
	}
	fileAddErr := r.Files.Add(newFileEmpty)
	if validators.HasError(fileAddErr) {
		return fileAddErr
	}
	return r.ShowFile(newFileEmpty.Id)
}

func (r *EditorStruct) OpenFileAndShow(filePath string) error {
	if len(filePath) == 0 {
		return errors.New("file path is empty, nothing to open")
	}

	isOpened := r.Files.IsPathPresentInMap(filePath)
	if isOpened {
		return fmt.Errorf("file with path %s already is opened", filePath)
	}

	text, ioErr := io.GetTextFromFile(filePath)
	if validators.HasError(ioErr) {
		return ioErr
	}

	fileWithData, fileCreationErr := r.FileHelper.CreateNewFileWithData(filePath, text)
	if validators.HasError(fileCreationErr) {
		return fileCreationErr
	}

	fileAddErr := r.Files.Add(fileWithData)
	if validators.HasError(fileAddErr) {
		return fileAddErr
	}

	return r.ShowFile(fileWithData.Id)
}

func (r *EditorStruct) SaveFile(fileId int64) error {
	file, getFileErr := r.GetFileById(fileId)
	if validators.HasError(getFileErr) {
		return getFileErr
	}

	saveErr := io.SaveTextToFile(file.Path, file.ActualContent)
	if validators.HasError(saveErr) {
		return saveErr
	}

	name, getNameErr := r.FileHelper.GetFileNameFromPath(file.Path)
	if validators.HasError(getNameErr) {
		return getNameErr
	}

	ext, getExtErr := r.FileHelper.GetFileExtensionFromPath(file.Path)
	if validators.HasError(getExtErr) {
		return getExtErr
	}

	fileType, getTypeErr := r.FileHelper.GetFileTypeFromExtension(ext)
	if validators.HasError(getTypeErr) {
		return getTypeErr
	}

	file.InitialContent = file.ActualContent
	file.New = false
	file.Changed = false
	file.Name = name
	file.Type = string(fileType)
	file.Extension = string(ext)

	return nil
}

func (r *EditorStruct) CloseFile(fileId int64) error {
	file, getFileErr := r.GetFileById(fileId)
	if validators.HasError(getFileErr) {
		return getFileErr
	}

	removeErr := r.Files.Remove(file.Id)
	if validators.HasError(removeErr) {
		return removeErr
	}

	var nextFile *types.FileStruct

	for _, fileStruct := range r.Files {
		nextFile = fileStruct
		break
	}

	if validators.IsNil(nextFile) {
		empty, fileCreationErr := r.FileHelper.CreateNewFileEmpty()
		if validators.HasError(fileCreationErr) {
			return fileCreationErr
		}

		fileAddErr := r.Files.Add(empty)
		if validators.HasError(fileAddErr) {
			return fileAddErr
		}

		nextFile = empty
	}

	return r.ShowFile(nextFile.Id)
}

func (r *EditorStruct) ShowFile(fileId int64) error {
	file, getFileErr := r.GetFileById(fileId)
	if validators.HasError(getFileErr) {
		return getFileErr
	}

	r.HideAllFiles()

	file.Opened = true

	return nil
}

func (r *EditorStruct) GetFileById(fileId int64) (*types.FileStruct, error) {
	return r.Files.GetById(fileId)
}

func (r *EditorStruct) HideAllFiles() {
	for _, file := range r.Files {
		file.Opened = false
	}
}

func (r *EditorStruct) UpdateFileContent(fileId int64, content string) error {
	file, getFileErr := r.GetFileById(fileId)
	if validators.HasError(getFileErr) {
		return getFileErr
	}

	file.ActualContent = content
	file.Changed = file.InitialContent != file.ActualContent

	return nil
}

func (r *EditorStruct) UpdateFileInformation(fileId int64, updateStruct types.FileTypeUpdateStruct) error {
	id := updateStruct.Id
	if fileId != id {
		return fmt.Errorf("file id passed to function is different to id from UI. %d!=%d", fileId, id)
	}

	file, getFileErr := r.GetFileById(fileId)
	if validators.HasError(getFileErr) {
		return getFileErr
	}

	fileType := updateStruct.Type
	extension := updateStruct.Extension

	if file.New && file.InitialExtension == "" {
		file.Type = fileType
		file.Extension = extension
		return nil
	}

	var isNewFile = false
	if file.Type != fileType || file.InitialExtension != extension {
		isNewFile = true
	}

	file.Type = fileType
	file.Extension = extension
	file.New = isNewFile
	return nil
}

func CreateIEditor(fileHelper types.IFileHelper, typeManager types.ITypeManager) types.IEditor {
	validators.PanicOnNil(fileHelper, "IFileHelper")
	validators.PanicOnNil(typeManager, "ITypeManager")

	filesMap := make(types.FilesMap, 1)

	editorStruct := EditorStruct{
		Files:       filesMap,
		FileHelper:  fileHelper,
		TypeManager: typeManager,
	}

	fileCreateErr := editorStruct.CreateFileAndShow()
	if validators.HasError(fileCreateErr) {
		panic("Failed to create empty file during editor creation")
	}

	return &editorStruct
}
