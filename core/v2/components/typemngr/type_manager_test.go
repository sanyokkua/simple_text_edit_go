package typemngr

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/v2/api"
	"testing"
)

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
	f1 := api.FileTypesJsonStruct{
		Key:        "python",
		Name:       "python",
		Extensions: []string{"py"},
	}
	manager := CreateTypeManager([]api.FileTypesJsonStruct{f1})
	filters := manager.GetSupportedFileFilters()

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

func TestGetTypeInformationByKey(t *testing.T) {
	f1 := api.FileTypesJsonStruct{
		Key:        "python",
		Name:       "Python",
		Extensions: []string{"py"},
	}
	f2 := api.FileTypesJsonStruct{
		Key:        "java",
		Name:       "Java",
		Extensions: []string{"java", "jdk"},
	}
	manager := CreateTypeManager([]api.FileTypesJsonStruct{f1, f2})

	keyPython := manager.GetTypeStructByKey("python")
	if keyPython == nil {
		t.Fatalf("Python struct is not found")
	}
	if keyPython.Key != "python" {
		t.Fatalf("Key is not correct. Expected 'python', Actual: %s", keyPython.Key)
	}
	if keyPython.Name != "Python" {
		t.Fatalf("Name is not correct. Expected 'Python', Actual: %s", keyPython.Name)
	}
	if len(keyPython.Extensions) != 1 {
		t.Fatalf("Extensions amount is not correct. Expected: 1, Actual: %d", len(keyPython.Extensions))
	}

	keyJava := manager.GetTypeStructByKey("java")
	if keyJava == nil {
		t.Fatalf("Java struct is not found")
	}
	if keyJava.Key != "java" {
		t.Fatalf("Key is not correct. Expected 'java', Actual: %s", keyJava.Key)
	}
	if keyJava.Name != "Java" {
		t.Fatalf("Name is not correct. Expected 'Java', Actual: %s", keyJava.Name)
	}
	if len(keyJava.Extensions) != 2 {
		t.Fatalf("Extensions amount is not correct. Expected: 2, Actual: %d", len(keyJava.Extensions))
	}

}

//TODO: add more tests

func findFilterInSlice(name string, slice []runtime.FileFilter) *runtime.FileFilter {
	for _, filter := range slice {
		if filter.DisplayName == name {
			return &filter
		}
	}
	return nil
}
