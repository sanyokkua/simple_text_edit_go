package app

import (
	"errors"
	"fmt"
	"simple_text_editor/core/v2/api"
	"simple_text_editor/core/v2/files"
	"simple_text_editor/core/v2/utils"
	"simple_text_editor/core/v2/utils/io"
)

type editor struct {
	files      map[int64]*api.FileStruct
	extensions map[string]api.FileTypesJsonStruct
}

func (r *editor) GetAllFilesInfo() []api.FileInfoStruct {
	shortInfo := make([]api.FileInfoStruct, 0, len(r.files))

	for _, file := range r.files {
		shortInfo = append(shortInfo, api.FileInfoStruct{
			Id:        file.Id,
			Path:      file.Path,
			Name:      file.Name,
			Type:      file.Type,
			Extension: file.Extension,
			New:       file.New,
			Opened:    file.Opened,
			Changed:   file.Changed,
		})
	}

	return shortInfo
}
func (r *editor) GetOpenedFile() (*api.FileStruct, error) {
	for _, file := range r.files {
		f := *file
		if f.Opened {
			return file, nil
		}
	}

	return nil, errors.New("opened file not found")
}
func (r *editor) CreateNewFileInEditor() error {
	newFileEmpty := files.CreateNewFileEmpty()
	r.addFileToMemory(&newFileEmpty)
	return r.SwitchOpenedFileTo(newFileEmpty.Id)
}
func (r *editor) OpenFile(filePath string) error {
	if len(filePath) == 0 {
		return errors.New("file path is empty, nothing to open")
	}

	isOpened := r.IsFileOpenedInEditor(filePath)
	if isOpened {
		return fmt.Errorf("file with path %s already is opened", filePath)
	}

	text, err := io.GetTextFromFile(filePath)
	if err != nil {
		return err
	}

	fileWithData := files.CreateNewFileWithData(filePath, text, r.extensions)
	r.addFileToMemory(&fileWithData)
	return r.SwitchOpenedFileTo(fileWithData.Id)
}
func (r *editor) SaveFile(fileId int64) error {
	file, err := r.GetFileById(fileId)
	if err != nil {
		return err
	}

	saveErr := io.SaveTextToFile(file.Path, file.ActualContent)
	if saveErr != nil {
		return saveErr
	}

	name := utils.GetFileNameFromPath(file.Path)
	ext := utils.GetFileExtensionFromPath(file.Path)
	ftype := utils.GetFileType(ext, r.extensions)

	file.InitialContent = file.ActualContent
	file.New = false
	file.Changed = false
	file.Name = name
	file.Type = ftype
	file.Extension = ext

	return nil
}
func (r *editor) CloseFile(fileId int64) error {
	file, err := r.GetFileById(fileId)
	if err != nil {
		return err
	}

	delete(r.files, file.Id)

	var nextFile *api.FileStruct
	for _, fileStruct := range r.files {
		nextFile = fileStruct
		break
	}

	if nextFile == nil {
		empty := files.CreateNewFileEmpty()
		r.addFileToMemory(&empty)
		nextFile = &empty
	}

	return r.SwitchOpenedFileTo(nextFile.Id)
}
func (r *editor) SwitchOpenedFileTo(fileId int64) error {
	file, err := r.GetFileById(fileId)
	if err != nil {
		return err
	}

	r.InactivateAllFiles()
	file.Opened = true

	return nil
}

func (r *editor) GetFileById(fileId int64) (*api.FileStruct, error) {
	for _, file := range r.files {
		f := *file
		if f.Id == fileId {
			return file, nil
		}
	}

	return nil, fmt.Errorf("file by provided if %d is not found", fileId)
}
func (r *editor) InactivateAllFiles() {
	for _, file := range r.files {
		file.Opened = false
	}
}
func (r *editor) IsFileOpenedInEditor(filePath string) bool {
	for _, file := range r.files {
		f := *file
		if f.Path == filePath {
			return true
		}
	}

	return false
}
func (r *editor) UpdateFileContent(fileId int64, content string) error {
	file, err := r.GetFileById(fileId)
	if err != nil {
		return err
	}

	file.ActualContent = content
	file.Changed = file.InitialContent != file.ActualContent

	return nil
}
func (r *editor) UpdateFileInformation(fileId int64, information api.FileInfoUpdateStruct) error {
	id := information.Id
	if fileId != id {
		return fmt.Errorf("file id passed to function is differenf with id from UI. %d!=%d", fileId, id)
	}

	file, err := r.GetFileById(fileId)
	if err != nil {
		return err
	}

	name := information.Name
	fileType := information.Type
	extension := information.Extension

	if file.New {
		file.Name = name
		file.Type = fileType
		file.Extension = extension
		return nil
	}

	file.Name = name
	file.Path = ""
	file.Type = fileType
	file.Extension = extension
	file.New = true
	return nil
}

func (r *editor) addFileToMemory(file *api.FileStruct) {
	r.files[file.Id] = file
}
func CreateEditor(extensions *map[string]api.FileTypesJsonStruct) api.IEditor {
	if extensions == nil {
		panic("Extensions passed to editor factory method is nil")
	}

	editorFilesStorage := make(map[int64]*api.FileStruct)
	editorApp := editor{
		files:      editorFilesStorage,
		extensions: *extensions,
	}

	file := files.CreateNewFileEmpty()
	editorApp.addFileToMemory(&file)
	err := editorApp.SwitchOpenedFileTo(file.Id)
	if err != nil {
		return nil
	}

	return &editorApp
}
