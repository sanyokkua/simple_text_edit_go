package api

type JsApi interface {
	GetFilesInformation() []InformationApi
	FindOpenedFile() FileApi
	ChangeFileStatusToOpened(uniqueIdentifier int64)
	ChangeFileContent(uniqueIdentifier int64, content string) bool
}
