package file

import "simple_text_editor/core/api"

type OpenedFileStruct struct {
	FileInfo        api.FileInformation // FileInfo - contains base file information
	OriginalContent string              // OriginalContent - Content that was read after open
	ActualContent   string              // ActualContent - Content that can be changed during the time
}

func (r *OpenedFileStruct) GetInformation() *api.FileInformation {
	return &r.FileInfo
}
func (r *OpenedFileStruct) GetOriginalContent() string {
	return r.OriginalContent
}
func (r *OpenedFileStruct) GetActualContent() string {
	return r.ActualContent
}
func (r *OpenedFileStruct) HasChanges() bool {
	return r.ActualContent == r.OriginalContent
}
func (r *OpenedFileStruct) SetOriginalContent(content string) {
	r.OriginalContent = content
}
func (r *OpenedFileStruct) SetActualContent(content string) {
	r.ActualContent = content
}
