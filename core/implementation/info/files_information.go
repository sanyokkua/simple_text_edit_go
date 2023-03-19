package info

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
	return r.OpenTimeStamp
}
func (r *FileInformationStruct) GetPath() string {
	return r.FilePath
}
func (r *FileInformationStruct) GetName() string {
	return r.FileName
}
func (r *FileInformationStruct) GetExt() string {
	return r.FileExtension
}
func (r *FileInformationStruct) GetType() string {
	return r.FileType
}
func (r *FileInformationStruct) Exists() bool {
	return r.FileExists
}
func (r *FileInformationStruct) IsOpened() bool {
	return r.FileIsOpened
}
func (r *FileInformationStruct) HasChanges() bool {
	return r.FileHasChanges
}
func (r *FileInformationStruct) SetOpenTimeStamp(timestamp int64) {
	r.OpenTimeStamp = timestamp
}
func (r *FileInformationStruct) SetPath(path string) {
	r.FilePath = path
}
func (r *FileInformationStruct) SetName(name string) {
	r.FileName = name
}
func (r *FileInformationStruct) SetExt(extension string) {
	r.FileExtension = extension
}
func (r *FileInformationStruct) SetType(typeVal string) {
	r.FileType = typeVal
}
func (r *FileInformationStruct) SetExists(exists bool) {
	r.FileExists = exists
}
func (r *FileInformationStruct) SetIsOpened(opened bool) {
	r.FileIsOpened = opened
}
func (r *FileInformationStruct) SetHasChanges(changed bool) {
	r.FileHasChanges = changed
}
