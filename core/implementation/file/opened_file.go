package file

import (
	"github.com/labstack/gommon/log"
	"simple_text_editor/core/api"
)

type OpenedFileStruct struct {
	FileInfo        *api.FileInformation // FileInfo - contains base file information
	OriginalContent string               // OriginalContent - Content that was read after open
	ActualContent   string               // ActualContent - Content that can be changed during the time
}

func (r *OpenedFileStruct) GetInformation() *api.FileInformation {
	log.Info("GetInformation")
	return r.FileInfo
}
func (r *OpenedFileStruct) GetOriginalContent() string {
	log.Info("GetOriginalContent", r.OriginalContent)
	return r.OriginalContent
}
func (r *OpenedFileStruct) GetActualContent() string {
	log.Info("GetActualContent", r.ActualContent)
	return r.ActualContent
}
func (r *OpenedFileStruct) HasChanges() bool {
	log.Info("HasChanges ->", r.ActualContent, r.OriginalContent, r.ActualContent != r.OriginalContent)
	return r.ActualContent != r.OriginalContent
}
func (r *OpenedFileStruct) SetOriginalContent(content string) {
	log.Info("SetOriginalContent ->", content)
	r.OriginalContent = content
	(*r.FileInfo).SetHasChanges(r.HasChanges())
}
func (r *OpenedFileStruct) SetActualContent(content string) {
	log.Info("SetActualContent ->", content)
	r.ActualContent = content
	(*r.FileInfo).SetHasChanges(r.HasChanges())
}
