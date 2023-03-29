package utils

import (
	"simple_text_editor/core/v2/api"
	"testing"
)

func TestGetFileNameFromPathWhereNameIsEmpty(t *testing.T) {
	res := GetFileNameFromPath("")

	if res != "" {
		t.Errorf("FileName for empty path should be empty string, not result value = %s", res)
	}
}

func TestGetFileNameFromPathWhereNameWithoutExtension(t *testing.T) {
	res := GetFileNameFromPath("fileName1")

	if res != "fileName1" {
		t.Errorf("FileName is not valid, expected: %s, actual: %s", "fileName1", res)
	}
}
func TestGetFileNameFromPathWhereNameWithExtension(t *testing.T) {
	res := GetFileNameFromPath("fileName1.txt")

	if res != "fileName1.txt" {
		t.Errorf("FileName is not valid, expected: %s, actual: %s", "fileName1.txt", res)
	}
}
func TestGetFileNameFromPathWhereNameWithExtensionInFolder(t *testing.T) {
	res := GetFileNameFromPath("/var/folder/1/2/3/fileName2.py")

	if res != "fileName2.py" {
		t.Errorf("FileName is not valid, expected: %s, actual: %s", "fileName2.py", res)
	}
}

func TestGetFileExtensionFromPathWherePathIsEmpty(t *testing.T) {
	res := GetFileExtensionFromPath("")

	if res != "" {
		t.Errorf("File Extension for empty path should be empty string, not result value = %s", res)
	}
}

func TestGetFileExtensionFromPathWherePathWithoutExtension(t *testing.T) {
	res := GetFileExtensionFromPath("fileName1")

	if res != "" {
		t.Errorf("File Extension is not valid, expected: %s, actual: %s", "", res)
	}
}
func TestGetFileExtensionFromPathWherePathWithExtension(t *testing.T) {
	res := GetFileExtensionFromPath("fileName1.txt")

	if res != ".txt" {
		t.Errorf("File Extension is not valid, expected: %s, actual: %s", ".txt", res)
	}
}
func TestGetFileExtensionFromPathWherePathWithExtensionInFolder(t *testing.T) {
	res := GetFileExtensionFromPath("/var/folder/1/2/3/fileName2.py")

	if res != ".py" {
		t.Errorf("File Extension is not valid, expected: %s, actual: %s", ".py", res)
	}
}

func TestGetFileTypeForEmptyLine(t *testing.T) {
	extensions := make(map[string]api.FileTypesJsonStruct, 1)
	extensions["txt"] = api.FileTypesJsonStruct{
		Key:        "txt",
		Name:       "Plain Text",
		Extensions: []string{"txt"},
	}
	res := GetFileType("", extensions)

	if res != "" {
		t.Errorf("File Type is not valid, expected: %s, actual: %s", "", res)
	}
}

func TestGetFileTypeForExtensionWithoutDotInPath(t *testing.T) {
	extensions := make(map[string]api.FileTypesJsonStruct, 1)
	extensions["txt"] = api.FileTypesJsonStruct{
		Key:        "txt",
		Name:       "Plain Text",
		Extensions: []string{"txt"},
	}
	res := GetFileType("ext", extensions)

	if res != "ext" {
		t.Errorf("File Type is not valid, expected: %s, actual: %s", "", res)
	}
}

func TestGetFileTypeForExtensionWithCorrectExt(t *testing.T) {
	extensions := make(map[string]api.FileTypesJsonStruct, 1)
	extensions["py"] = api.FileTypesJsonStruct{
		Key:        "python",
		Name:       "python",
		Extensions: []string{"py"},
	}
	res := GetFileType(".py", extensions)

	if res != "python" {
		t.Errorf("File Type is not valid, expected: %s, actual: %s", "", res)
	}
}
