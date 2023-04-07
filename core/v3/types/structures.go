package types

// Frontend structures

type FrontendFileContainerStruct struct {
	HasError bool
	Error    string
	File     FileStruct
}

type FrontendFileInfoArrayContainerStruct struct {
	HasError bool
	Error    string
	Files    []FileInfoStruct
}

type FrontendKeyValueArrayContainerStruct struct {
	HasError bool
	Error    string
	Pairs    []KeyValuePairStruct
}

// Internal structures

type FileTypesJsonStruct struct {
	Key        FileTypeKey         `json:"key"`
	Name       string              `json:"name"`
	Extensions []FileTypeExtension `json:"extensions"`
}
type KeyValuePairStruct struct {
	Key   string
	Value string
}
type FileTypeUpdateStruct struct {
	Id        int64
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

// FileStruct defines structure to keep all the information about file that currently opened in memory of the app.
// Id - is used to access files. 27648617239328 (any number)
// Path - Full path to the file if it was opened. Ex: /home/username/AwesomeFile.txt
// Name - Name of the file. Example AwesomeFile
// Type - key that represent type in
// Extension - .txt
// InitialExtension - "" or ".txt" or any other
// InitialContent - "" or initial text
// ActualContent - "" or initial text
// New - True if the file was just created or extension changed and file should be saved as new file
// Opened - True if file currently displayed on the screen (opened in the editor widget)
// Changed - True if InitialContent is not equal to ActualContent
type FileStruct struct {
	Id               int64
	Path             string
	Name             string
	Type             string
	Extension        string
	InitialExtension string
	InitialContent   string
	ActualContent    string
	New              bool
	Opened           bool
	Changed          bool
}
