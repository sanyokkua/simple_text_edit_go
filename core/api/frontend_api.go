package api

type JsApi interface {
	GetFilesInformation() []FileInformation
	FindOpenedFile() OpenedFile
	ChangeFileStatusToOpened(uniqueIdentifier int64)
	ChangeFileContent(uniqueIdentifier int64, content string) bool
}
