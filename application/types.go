package application

import "context"

type EditorApplication struct {
	AppContext context.Context
	AppFiles   map[int64]*AppFile
}

type FileDescriptor struct {
	FileId   int64
	FilePath string
	FileName string
	FileType string
	IsActive bool
}

type AppFile struct {
	Descriptor      FileDescriptor
	FileContent     string
	SelectedContent string
	ContentHistory  map[int64]string
}
