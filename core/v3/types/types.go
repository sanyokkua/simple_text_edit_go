package types

import (
	"context"
	"errors"
	"simple_text_editor/core/v3/validators"
	"strconv"
)

//go:generate mockery --name ContextProvider
type ContextProvider func() (ctx context.Context)
type Destination string

type Button string

const (
	BtnOk     Button = "Ok"
	BtnCancel Button = "Cancel"
)

const NewFileName = "*No name*"

func (r Button) EqualTo(another Button) bool {
	return r == another
}

type FilesMap map[int64]*FileStruct

func (r *FilesMap) GetById(fileId int64) (*FileStruct, error) {
	deref := *r
	fileStruct, ok := deref[fileId]
	if !ok {
		msg := "File not found with id: " + strconv.FormatInt(fileId, 10)
		return nil, errors.New(msg)
	}
	return fileStruct, nil
}

func (r *FilesMap) Remove(fileId int64) error {
	deref := *r
	_, ok := deref[fileId]
	if !ok {
		msg := "File not found with id: " + strconv.FormatInt(fileId, 10)
		return errors.New(msg)
	}
	delete(*r, fileId)
	return nil
}

func (r *FilesMap) Add(fileStruct *FileStruct) error {
	if fileStruct == nil {
		return errors.New("file Struct is NIL")
	}

	deref := *r
	_, ok := deref[fileStruct.Id]
	if ok {
		return errors.New("file with passed ID already exists")
	}

	deref[fileStruct.Id] = fileStruct
	return nil
}

func (r *FilesMap) IsPathPresentInMap(path string) bool {
	if validators.IsEmptyString(path) {
		return false
	}

	deref := *r
	for _, file := range deref {
		if file.Path == path {
			return true
		}
	}

	return false
}

func (r *FilesMap) GetSize() int {
	deref := *r
	return len(deref)
}

func (r *FilesMap) GetSlice() []*FileStruct {
	deref := *r
	files := make([]*FileStruct, 0, deref.GetSize())
	for _, fileStruct := range deref {
		f := fileStruct
		files = append(files, f)
	}
	return files
}

func (r *FilesMap) Find(predicate func(file *FileStruct) bool) *FileStruct {
	deref := *r
	for _, fileStruct := range deref {
		f := fileStruct
		if predicate(f) {
			return f
		}
	}
	return nil
}

func (r *FilesMap) Foreach(action func(file *FileStruct)) {
	deref := *r
	for _, fileStruct := range deref {
		f := fileStruct
		action(f)
	}
}

type FileTypeKey string
type FileTypeExtension string

type ExtensionsMap map[FileTypeExtension]*FileTypesJsonStruct

func (r *ExtensionsMap) GetByExtension(ext FileTypeExtension) (*FileTypesJsonStruct, error) {
	deref := *r
	fileStruct, ok := deref[ext]
	if !ok {
		msg := "Extension not found with id: " + string(ext)
		return nil, errors.New(msg)
	}
	return fileStruct, nil
}

type TypesMap map[FileTypeKey]*FileTypesJsonStruct

func (r *TypesMap) GetByTypeKey(ext FileTypeKey) (*FileTypesJsonStruct, error) {
	deref := *r
	fileStruct, ok := deref[ext]
	if !ok {
		msg := "Type not found with id: " + string(ext)
		return nil, errors.New(msg)
	}
	return fileStruct, nil
}
