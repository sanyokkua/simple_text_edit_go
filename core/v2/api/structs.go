package api

type FileTypesJsonStruct struct {
	Key        string   `json:"key"`
	Name       string   `json:"name"`
	Extensions []string `json:"extensions"`
}
type KeyValuePairStruct struct {
	Key   string
	Value string
}
type FileInfoUpdateStruct struct {
	Id        int64
	Name      string
	Type      string
	Extension string
}
type FileInfoStruct struct {
	Id        int64
	Path      string
	Name      string
	Type      string
	Extension string
	New       bool
	Opened    bool
	Changed   bool
}
type FileStruct struct {
	Id             int64
	Path           string
	Name           string
	Type           string
	Extension      string
	InitialContent string
	ActualContent  string
	New            bool
	Opened         bool
	Changed        bool
}
