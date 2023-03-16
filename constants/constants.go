package constants

import "github.com/wailsapp/wails/v2/pkg/runtime"

const EVENT_FILE_IS_CHOSEN = "EVENT_FILE_IS_CHOSEN"
const EVENT_IS_FILE_OPENED = "EVENT_IS_FILE_OPENED"
const GENERIC_ERROR_HAPPENED = "GENERIC_ERROR_HAPPENED"
const EVENT_FILE_SHOULD_BE_SAVED = "EVENT_FILE_SHOULD_BE_SAVED"
const EVENT_TEXT_SHOULD_BE_SORTED = "EVENT_FILE_SHOULD_BE_SAVED"

const (
	SORT_ASC              int = 0
	SORT_DESC             int = 1
	SORT_ASC_IGNORE_CASE  int = 2
	SORT_DESC_IGNORE_CASE int = 3
)

func GetSupportedFileFilters() []runtime.FileFilter {
	return []runtime.FileFilter{
		{
			DisplayName: "Plain Text",
			Pattern:     "*.txt;",
		},
		{
			DisplayName: "Go",
			Pattern:     "*.go;",
		},
		{
			DisplayName: "Java",
			Pattern:     "*.java;",
		},
		{
			DisplayName: "TypeScript",
			Pattern:     "*.ts;*.tsx",
		},
		{
			DisplayName: "Javascript",
			Pattern:     "*.js;*.jsx",
		},
		{
			DisplayName: "Python",
			Pattern:     "*.py;",
		},
		{
			DisplayName: "XML",
			Pattern:     "*.xml;*.pom;",
		},
		{
			DisplayName: "Json",
			Pattern:     "*.json;",
		},
		{
			DisplayName: "Other",
			Pattern:     "*.toml;*.md;*.mod;*.sum;",
		},
	}
}
