package info

import "github.com/labstack/gommon/log"

type FileInformationStruct struct {
	OpenTimeStamp  int64  // OpenTimeStamp - Used to place tab in right order
	FilePath       string // FilePath - Full path to the file (empty for new file)
	FileName       string // FileName of the file (last item from the path without extension)
	FileExtension  string // FileExtension - Extension of the file (last item after . in path if available) (empty for new file)
	FileType       string // FileType based on the extension. Can be the same or different
	FileExists     bool   // FileExists equal true if this file was opened and not just created
	FileIsOpened   bool   // FileIsOpened equal true if this file should be shown on the UI now
	FileHasChanges bool   // FileHasChanges equal true if this the actual content is not equal to original content
}

func (r *FileInformationStruct) GetOpenTimeStamp() int64 {
	log.Info("GetOpenTimeStamp", r.OpenTimeStamp)
	return r.OpenTimeStamp
}
func (r *FileInformationStruct) GetPath() string {
	log.Info("GetPath", r.FilePath)
	return r.FilePath
}
func (r *FileInformationStruct) GetName() string {
	log.Info("GetName", r.FileName)
	return r.FileName
}
func (r *FileInformationStruct) GetExt() string {
	log.Info("GetExt", r.FileExtension)
	return r.FileExtension
}
func (r *FileInformationStruct) GetType() string {
	log.Info("GetType", r.FileType)
	return r.FileType
}
func (r *FileInformationStruct) Exists() bool {
	log.Info("Exists", r.FileExists)
	return r.FileExists
}
func (r *FileInformationStruct) IsOpened() bool {
	log.Info("IsOpened", r.FileIsOpened)
	return r.FileIsOpened
}
func (r *FileInformationStruct) HasChanges() bool {
	log.Info("HasChanges", r.FileHasChanges)
	return r.FileHasChanges
}
func (r *FileInformationStruct) SetOpenTimeStamp(timestamp int64) {
	log.Info("SetOpenTimeStamp", timestamp)
	r.OpenTimeStamp = timestamp
}
func (r *FileInformationStruct) SetPath(path string) {
	log.Info("SetPath", path)
	r.FilePath = path
}
func (r *FileInformationStruct) SetName(name string) {
	log.Info("SetName", name)
	r.FileName = name
}
func (r *FileInformationStruct) SetExt(extension string) {
	log.Info("SetExt", extension)
	r.FileExtension = extension
}
func (r *FileInformationStruct) SetType(typeVal string) {
	log.Info("SetType", typeVal)
	r.FileType = typeVal
}
func (r *FileInformationStruct) SetExists(exists bool) {
	log.Info("SetExists", exists)
	r.FileExists = exists
}
func (r *FileInformationStruct) SetIsOpened(opened bool) {
	log.Info("SetIsOpened", opened)
	r.FileIsOpened = opened
}
func (r *FileInformationStruct) SetHasChanges(changed bool) {
	log.Info("SetHasChanges", changed)
	r.FileHasChanges = changed
}
