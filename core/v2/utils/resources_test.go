package utils

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"simple_text_editor/core/v2/api"
	"testing"
)

const TestPath = "../../../fileTypes.json"

func TestLoadFileTypesJsonWithEmptyPath(t *testing.T) {
	res := LoadFileTypesJson("")

	if res == nil {
		t.Errorf("File Types slice should not be nil")
	}

	if len(res) > 0 {
		t.Errorf("File Types slice should have length = 0")
	}
}

func TestLoadFileTypesJsonWithCorrectPath(t *testing.T) {
	println(os.Getwd())

	res := LoadFileTypesJson(TestPath)

	if res == nil {
		t.Errorf("File Types slice should not be nil")
	}

	if len(res) == 0 {
		t.Errorf("File Types slice should have length > 0")
	}
}

func TestCreateFileFilter(t *testing.T) {
	fileType := api.FileTypesJsonStruct{
		Key:        "cpp",
		Name:       "C++",
		Extensions: []string{"cpp", "hh", "c++"},
	}
	filter := createFileFilter(&fileType)
	if filter.DisplayName != fileType.Name {
		t.Errorf("Filter name is not correct. Expected: %s, Actual: %s", fileType.Name, filter.DisplayName)
	}
	if filter.Pattern != "*.cpp;*.hh;*.c++;" {
		t.Errorf("Filter pattern is not correct. Expected: %s, Actual: %s", "*.cpp;*.hh;*.c++;", filter.Pattern)
	}
}

func TestGetSupportedFileFilters(t *testing.T) {
	extensions := make(map[string]api.FileTypesJsonStruct, 1)
	extensions["py"] = api.FileTypesJsonStruct{
		Key:        "python",
		Name:       "python",
		Extensions: []string{"py"},
	}
	filters := GetSupportedFileFilters(extensions)

	if len(filters) != 3 {
		t.Errorf("Size of created filters slice is not equear 3. Actual: %d", len(filters))
	}
	python := findFilterInSlice("python", filters)

	if python == nil {
		t.Errorf("Not found in returned slice Python")
	}
	pyPattern := "*.py;"
	if python.Pattern != pyPattern {
		t.Errorf("Python pattern is not valid. Expected: %s Actual: %s", pyPattern, python.Pattern)
	}

	txt := findFilterInSlice("Plain Text", filters)
	txtPattern := "*.txt"
	if txt.Pattern != txtPattern {
		t.Errorf("txt pattern is not valid. Expected: %s Actual: %s", txtPattern, txt.Pattern)
	}
	if txt == nil {
		t.Errorf("Not found in returned slice TXT")
	}

	anyFile := findFilterInSlice("Any File", filters)
	anyPattern := ""
	if anyFile.Pattern != anyPattern {
		t.Errorf("anyFile pattern is not valid. Expected: %s Actual: %s", anyPattern, anyFile.Pattern)
	}
	if anyFile == nil {
		t.Errorf("Not found in returned slice any file type")
	}
}

func findFilterInSlice(name string, slice []runtime.FileFilter) *runtime.FileFilter {
	for _, filter := range slice {
		if filter.DisplayName == name {
			return &filter
		}
	}
	return nil
}
