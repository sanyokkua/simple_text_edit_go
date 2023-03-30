package files

import (
	"simple_text_editor/core/v2/api"
	"simple_text_editor/core/v2/components/typemngr"
	"testing"
)

func TestCreateNewFileEmpty(t *testing.T) {
	empty := CreateNewFileEmpty()
	if !empty.New {
		t.Errorf("New file should have true value in New field")
	}
	if empty.Id == 0 {
		t.Errorf("New file doesn't have generated id value")
	}
	if len(empty.Name) == 0 || empty.Name != "New" {
		t.Errorf("Name for new file should not be empty and should be equal to 'New'")
	}
	if empty.Opened {
		t.Errorf("Opened field should be false")
	}
	if empty.Changed {
		t.Errorf("Changed field should be false")
	}
	if len(empty.Path) > 0 {
		t.Errorf("Path field should be empty")
	}
	if len(empty.InitialContent) > 0 {
		t.Errorf("Initial content field should be empty")
	}
	if len(empty.ActualContent) > 0 {
		t.Errorf("Actual content field should be empty")
	}
}

func TestCreateNewFileWithData(t *testing.T) {
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
	manager := typemngr.CreateTypeManager([]api.FileTypesJsonStruct{f1, f2})

	originalContent := "Original Content of file"
	fileName := "file.py"
	filePath := "/test/path/" + fileName

	fileWithData := CreateNewFileWithData(filePath, originalContent, manager)

	if fileWithData.New {
		t.Errorf("New file should have false value in New field for existing file")
	}
	if fileWithData.Id == 0 {
		t.Errorf("File doesn't have generated id value")
	}
	if len(fileWithData.Name) == 0 || fileWithData.Name != fileName {
		t.Errorf("Name for new file should not be empty and should be equal to '%s', actual: %s",
			fileName, fileWithData.Name)
	}
	if fileWithData.Opened {
		t.Errorf("Opened field should be false")
	}
	if fileWithData.Changed {
		t.Errorf("Changed field should be false")
	}
	if len(fileWithData.Path) == 0 {
		t.Errorf("Path field should not be empty")
	}
	if fileWithData.Path != filePath {
		t.Errorf("Path field should not be equal to original. Expected: %s, Actual: %s",
			filePath, fileWithData.Path)
	}
	if len(fileWithData.InitialContent) == 0 {
		t.Errorf("Initial content field should NOT be empty")
	}
	if len(fileWithData.ActualContent) == 0 {
		t.Errorf("Actual content field should NOT be empty")
	}
	if fileWithData.InitialContent != originalContent {
		t.Errorf("Initial content field should NOT be empty, Expected: '%s', Actual: '%s'",
			originalContent, fileWithData.InitialContent)
	}
	if fileWithData.ActualContent != originalContent {
		t.Errorf("Actual content field should NOT be empty, Expected: '%s', Actual: '%s'",
			originalContent, fileWithData.ActualContent)
	}
}
