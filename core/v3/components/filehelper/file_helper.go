package filehelper

import (
	"path/filepath"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/validators"
	"strings"
)

type FileHelperStruct struct {
	UniqueIdGenerator types.IUniqueIdGenerator
	TypeManager       types.ITypeManager
}

func (r *FileHelperStruct) GetFileNameFromPath(filePath string) (string, error) {
	if validators.IsEmptyString(filePath) {
		return "", nil
	}
	fileName := filepath.Base(filePath)
	return fileName, nil
}

func (r *FileHelperStruct) GetFileExtensionFromPath(filePath string) (types.FileTypeExtension, error) {
	if validators.IsEmptyString(filePath) {
		return "", nil
	}
	fileName := filepath.Ext(filePath)
	return types.FileTypeExtension(fileName), nil
}

func (r *FileHelperStruct) GetFileTypeFromExtension(fileExtension types.FileTypeExtension) (types.FileTypeKey, error) {
	fileType, fileTypeErr := r.TypeManager.GetTypeKeyByExtension(fileExtension)
	if validators.HasError(fileTypeErr) {
		return "", fileTypeErr
	}
	return fileType, nil
}

func (r *FileHelperStruct) CreateNewFileEmpty() (*types.FileStruct, error) {
	id := r.UniqueIdGenerator.GenerateId()
	return &types.FileStruct{
		Id:   id,
		Name: types.NewFileName,
		New:  true,
	}, nil
}

func (r *FileHelperStruct) CreateNewFileWithData(path string, originalContent string) (*types.FileStruct, error) {
	id := r.UniqueIdGenerator.GenerateId()

	fileStruct := types.FileStruct{
		Id:             id,
		Path:           path,
		InitialContent: originalContent,
		ActualContent:  originalContent,
		Opened:         false,
		Changed:        false,
	}

	return r.UpdateFileDataOnFilePath(path, &fileStruct)
}

func (r *FileHelperStruct) UpdateFileDataOnFilePath(path string, file *types.FileStruct) (*types.FileStruct, error) {
	fileName, fileNameErr := r.GetFileNameFromPath(path)
	if validators.HasError(fileNameErr) {
		return nil, fileNameErr
	}

	fileExtension, fileExtErr := r.GetFileExtensionFromPath(path)
	if validators.HasError(fileExtErr) {
		return nil, fileExtErr
	}

	fileType, fileTypeErr := r.GetFileTypeFromExtension(fileExtension)
	if validators.HasError(fileTypeErr) {
		return nil, fileTypeErr
	}

	isNew := validators.IsEmptyString(path)

	if len(fileName) > 0 && len(fileExtension) > 0 && strings.HasSuffix(fileName, string(fileExtension)) {
		index := strings.LastIndex(fileName, string(fileExtension))
		fileName = fileName[:index]
	}

	file.New = isNew
	file.Name = fileName
	file.Type = string(fileType)
	file.Extension = string(fileExtension)

	return file, nil
}

func CreateIFileHelper(uniqueIdGenerator types.IUniqueIdGenerator, typeManager types.ITypeManager) types.IFileHelper {
	validators.PanicOnNil(uniqueIdGenerator, "IUniqueIdGenerator")
	validators.PanicOnNil(typeManager, "ITypeManager")

	return &FileHelperStruct{
		UniqueIdGenerator: uniqueIdGenerator,
		TypeManager:       typeManager,
	}
}
