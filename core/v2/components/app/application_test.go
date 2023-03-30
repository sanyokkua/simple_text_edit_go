package app

import (
	"fmt"
	"os"
	"simple_text_editor/core/v2/api"
	"simple_text_editor/core/v2/components/typemngr"
	"simple_text_editor/core/v2/utils"
	"strconv"
	"testing"
	"time"
)

func TestApplicationCreation(t *testing.T) {
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

	createdEditor := CreateEditor(manager)

	files := createdEditor.GetAllFilesInfo()
	if files == nil {
		t.Fatalf("Files are nil")
	}
	if len(files) != 1 {
		t.Fatalf("New app has incorrect number of opened files")
	}
	if !files[0].Opened {
		t.Fatalf("Newly created app has file wich is not opened")
	}
	if !files[0].New {
		t.Fatalf("Newly created app has file wich is not new")
	}
}

func TestApplicationCreationPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	CreateEditor(nil)
}

func TestApplicationCreateFile(t *testing.T) {
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

	createdEditor := CreateEditor(manager)
	firstFile, _ := createdEditor.GetOpenedFile()

	path := createTestFile()
	openErr := createdEditor.OpenFile(path)
	if openErr != nil {
		t.Fatalf("Error happened during opening file. Path: %s", path)
	}

	filesInfo := createdEditor.GetAllFilesInfo()
	if len(filesInfo) != 2 {
		t.Fatalf("Number of files is wrong. After opening file should be 2 files")
	}

	file, getOpenedFileErr := createdEditor.GetOpenedFile()
	if getOpenedFileErr != nil {
		t.Fatalf("Opened file was not returned. Error happened")
	}

	if path != file.Path {
		t.Fatalf("Opened file should be active in editor. Expected: %s, Actual: %s", "", file.Path)
	}

	switchErr := createdEditor.SwitchOpenedFileTo(firstFile.Id)
	if switchErr != nil {
		t.Fatalf("Error during switching active (opened) file")
	}

	if !firstFile.Opened {
		t.Fatalf("File should be opened")
	}

	if !createdEditor.IsFileOpenedInEditor(path) {
		t.Fatalf("Check for opened file is failed")
	}

	newFileErr := createdEditor.CreateNewFileInEditor()
	if newFileErr != nil {
		t.Fatalf("New file failed creation")
	}

	if len(createdEditor.GetAllFilesInfo()) != 3 {
		t.Fatalf("Number of created (opened) file in editor is wrong")
	}
}

func TestApplicationFileContentChanges(t *testing.T) {
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
	createdEditor := CreateEditor(manager)

	file, getOpenedFileErr := createdEditor.GetOpenedFile()
	if getOpenedFileErr != nil {
		t.Fatalf("Error on retrieving opened file")
	}
	if file.ActualContent != "" && file.InitialContent != "" {
		t.Fatalf("Error with initial data in the new file")
	}

	content := "Updated content"
	updErr := createdEditor.UpdateFileContent(file.Id, content)
	if updErr != nil {
		t.Fatalf("Error during updating content")
	}
	if file.InitialContent != "" {
		t.Fatalf("Content was changed for initial data")
	}
	if file.ActualContent != content {
		t.Fatalf("Content for file was not updated. Expected: %s, Actual: %s", content, file.ActualContent)
	}
}

func TestApplicationInactivateAllFiles(t *testing.T) {
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

	createdEditor := CreateEditor(manager)
	time.Sleep(5 * time.Millisecond)
	// Creation of file can be too fast so file with the same id (time used for id) will not be created. Sleep is used to fix it
	crErr1 := createdEditor.CreateNewFileInEditor()
	if crErr1 != nil {
		t.Fatalf("Error creating file in editor")
	}
	time.Sleep(5 * time.Millisecond)
	crErr2 := createdEditor.CreateNewFileInEditor()
	if crErr2 != nil {
		t.Fatalf("Error creating file in editor")
	}
	time.Sleep(5 * time.Millisecond)
	crErr3 := createdEditor.CreateNewFileInEditor()
	if crErr3 != nil {
		t.Fatalf("Error creating file in editor")
	}
	time.Sleep(5 * time.Millisecond)

	all := createdEditor.GetAllFilesInfo()

	if len(all) != 4 {
		t.Fatalf("Incorrect number of files in editor")
	}

	var counter1 = 0
	for _, infoStruct := range all {
		if infoStruct.Opened {
			counter1++
		}
	}
	if counter1 == 0 || counter1 > 1 {
		t.Fatalf("incorrect state of app, number of opened files not equal 1. Actual: %d", counter1)
	}

	createdEditor.InactivateAllFiles()

	all2 := createdEditor.GetAllFilesInfo()
	var counter2 = 0
	for _, infoStruct := range all2 {
		if infoStruct.Opened {
			counter2++
		}
	}

	if counter2 > 0 {
		t.Fatalf("incorrect state of app, number of opened files not equal 0 after inactivation. Actual: %d", counter2)
	}
}

func TestApplicationCloseFile(t *testing.T) {
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

	createdEditor := CreateEditor(manager)
	time.Sleep(5 * time.Millisecond)

	crErr1 := createdEditor.CreateNewFileInEditor()
	if crErr1 != nil {
		t.Fatalf("Error creating file in editor")
	}
	time.Sleep(5 * time.Millisecond)
	crErr2 := createdEditor.CreateNewFileInEditor()
	if crErr2 != nil {
		t.Fatalf("Error creating file in editor")
	}
	time.Sleep(5 * time.Millisecond)
	crErr3 := createdEditor.CreateNewFileInEditor()
	if crErr3 != nil {
		t.Fatalf("Error creating file in editor")
	}
	time.Sleep(5 * time.Millisecond)

	all := createdEditor.GetAllFilesInfo()

	if len(all) != 4 {
		t.Fatalf("Incorrect number of files in editor")
	}

	fileId1 := all[0].Id
	fileId2 := all[1].Id
	fileId3 := all[2].Id
	fileId4 := all[3].Id

	if fileId1 == fileId2 || fileId2 == fileId3 || fileId3 == fileId4 {
		t.Fatalf("Files can't have same ids")
	}

	fId1, errId1 := createdEditor.GetFileById(fileId1)
	if errId1 != nil {
		t.Fatalf("File1 is not found by ID, id: %d", fileId1)
	}
	if fId1.Id != fileId1 {
		t.Fatalf("Returned file has different id. Expected: %d, Actual: %d", fileId1, fId1.Id)
	}

	fId2, errId2 := createdEditor.GetFileById(fileId2)
	if errId2 != nil {
		t.Fatalf("File2 is not found by ID, id: %d", fileId2)
	}
	if fId2.Id != fileId2 {
		t.Fatalf("Returned file has different id. Expected: %d, Actual: %d", fileId2, fId2.Id)
	}

	fId3, errId3 := createdEditor.GetFileById(fileId3)
	if errId3 != nil {
		t.Fatalf("File3 is not found by ID, id: %d", fileId3)
	}
	if fId3.Id != fileId3 {
		t.Fatalf("Returned file has different id. Expected: %d, Actual: %d", fileId3, fId3.Id)
	}

	fId4, errId4 := createdEditor.GetFileById(fileId4)
	if errId4 != nil {
		t.Fatalf("File4 is not found by ID, id: %d", fileId4)
	}
	if fId4.Id != fileId4 {
		t.Fatalf("Returned file has different id. Expected: %d, Actual: %d", fileId4, fId3.Id)
	}
	switchErr1 := createdEditor.SwitchOpenedFileTo(fileId4)
	if switchErr1 != nil {
		t.Fatalf("Failed to switch to 4th file")
	}
	closeErr1 := createdEditor.CloseFile(fileId1)
	if closeErr1 != nil {
		t.Fatalf("Failed to close file")
	}
	if len(createdEditor.GetAllFilesInfo()) != 3 {
		t.Fatalf("Incorrect number of files in memory after close of 1 file")
	}
	_, err := createdEditor.GetFileById(fileId1)
	if err == nil {
		t.Fatalf("File 1 should not be found. It should be deleted")
	}

	closeErr2 := createdEditor.CloseFile(fileId4)
	if closeErr2 != nil {
		t.Fatalf("File4 should be closed without errors")
	}
	if len(createdEditor.GetAllFilesInfo()) != 2 {
		t.Fatalf("Incorrect number of files in memory after close of 2 files")
	}

	_, openedFileErr := createdEditor.GetOpenedFile()
	if openedFileErr != nil {
		t.Fatalf("Opened file should be found after closing opened file")
	}

	_, getFileByIdErr := createdEditor.GetFileById(fileId4)
	if getFileByIdErr == nil {
		t.Fatalf("File by closed ID should NOT be found")
	}

	err3 := createdEditor.CloseFile(fileId2)
	if err3 != nil {
		t.Fatalf("File 2 should be closed")
	}

	err4 := createdEditor.CloseFile(fileId3)
	if err4 != nil {
		t.Fatalf("File 3 should be closed")
	}

	if len(createdEditor.GetAllFilesInfo()) != 1 {
		t.Fatalf("When all files were closed, new one should be created and number of files should be 1")
	}

	newFile, err5 := createdEditor.GetOpenedFile()
	if err5 != nil {
		t.Fatalf("Opened file should be found")
	}
	if newFile.Id == fileId1 || newFile.Id == fileId2 || newFile.Id == fileId3 || newFile.Id == fileId4 {
		t.Fatalf("ID of new file should not be the same as IDs of closed files erlier")
	}

}

func TestApplicationGetFilesInfo(t *testing.T) {
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

	createdEditor := CreateEditor(manager)
	newFile, newFileErr := createdEditor.GetOpenedFile()
	if newFileErr != nil {
		t.Fatalf("Failed to get new file from editor")
	}
	time.Sleep(5 * time.Millisecond)

	testFilePath := createTestFile()
	err := createdEditor.OpenFile(testFilePath)
	if err != nil {
		t.Fatalf("Failed to load test data. Failed to open test file")
	}

	openedTestFile, errOpenTestFile := createdEditor.GetOpenedFile()
	if errOpenTestFile != nil {
		t.Fatalf("Error with getting Opened file in editor")
	}

	idOfTheNewFile := newFile.Id
	idOfTheTstFile := openedTestFile.Id

	if idOfTheNewFile == idOfTheTstFile {
		t.Fatalf("Incorrect state of the app. Id of the empty file and test file are the same")
	}

	updateErr := createdEditor.UpdateFileContent(idOfTheTstFile, "Update File Content")
	if updateErr != nil {
		t.Fatalf("Failed to update test file content")
	}

	allFilesInfo := createdEditor.GetAllFilesInfo()

	for _, fileInfoStruct := range allFilesInfo {
		if fileInfoStruct.Id == idOfTheNewFile {
			if "" != fileInfoStruct.Path {
				t.Fatalf("Path should be empty for new file")
			}
			if "New" != fileInfoStruct.Name {
				t.Fatalf("Name should be 'New' for new file")
			}
			if "" != fileInfoStruct.Type {
				t.Fatalf("Type should be empty for new file")
			}
			if "" != fileInfoStruct.Extension {
				t.Fatalf("Extension should be empty for new file")
			}
			if !fileInfoStruct.New {
				t.Fatalf("New should be true for new file")
			}
			if fileInfoStruct.Opened {
				t.Fatalf("Opened should be false")
			}
			if fileInfoStruct.Changed {
				t.Fatalf("Changed should be false")
			}
		} else if fileInfoStruct.Id == idOfTheTstFile {
			name := utils.GetFileNameFromPath(testFilePath)
			ext := utils.GetFileExtensionFromPath(testFilePath)
			fileType := manager.GetTypeKeyByExtension(ext)

			if testFilePath != fileInfoStruct.Path {
				t.Fatalf("Path should not be empty for opened test file")
			}
			if name != fileInfoStruct.Name {
				t.Fatalf("Name should NOT be 'New' for opened test file")
			}
			if fileType != fileInfoStruct.Type {
				t.Fatalf("Type should NOT be empty for opened test file")
			}
			if ext != fileInfoStruct.Extension {
				t.Fatalf("Extension should NOT be empty for opened test file")
			}
			if fileInfoStruct.New {
				t.Fatalf("New should be false")
			}
			if !fileInfoStruct.Opened {
				t.Fatalf("Opened should be true")
			}
			if !fileInfoStruct.Changed {
				t.Fatalf("Changed should be true")
			}
		} else {
			t.Fatalf("Id is unknown")
		}
	}
}

func TestApplicationSaveFile(t *testing.T) {
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

	createdEditor := CreateEditor(manager)

	openedFile, err := createdEditor.GetOpenedFile()
	if err != nil {
		t.Fatalf("Error getting opened file")
	}

	fileContent := "New Content"
	updErr := createdEditor.UpdateFileContent(openedFile.Id, fileContent)
	if updErr != nil {
		t.Fatalf("Error updating file content")
	}

	tempDir := os.TempDir()
	newFilePath := fmt.Sprintf("%s/%s.%s", tempDir, "test", "py")
	openedFile.Path = newFilePath

	updInfErr := createdEditor.UpdateFileInformation(openedFile.Id, api.FileInfoUpdateStruct{
		Id:        openedFile.Id,
		Type:      "python",
		Extension: "py",
	})
	if updInfErr != nil {
		t.Fatalf("Failed to update file information")
	}

	saveErr := createdEditor.SaveFile(openedFile.Id)
	if saveErr != nil {
		t.Fatalf("Save error")
	}

	bytes, readErr := os.ReadFile(newFilePath)
	if readErr != nil {
		t.Fatalf("Failed to read saved file")
	}
	readFileContent := string(bytes)
	if fileContent != readFileContent {
		t.Fatalf("Content wasn't written right")
	}
	if openedFile.New {
		t.Fatalf("File should not be new after success save")
	}
	if openedFile.Changed {
		t.Fatalf("File should not be changed after success save")
	}
	if openedFile.Name != "test.py" {
		t.Fatalf("File name should be equal to test.py, Actual: %s", openedFile.Name)
	}
	if openedFile.Path != newFilePath {
		t.Fatalf("File path should be equal to %s, Actual: %s", newFilePath, openedFile.Path)
	}
	if openedFile.Type != "Python" {
		t.Fatalf("File type should be equal to python, Actual: %s", openedFile.Type)
	}
	if openedFile.Extension != ".py" {
		t.Fatalf("File extension should be equal to py, Actual: %s", openedFile.Extension)
	}
	if openedFile.InitialContent != fileContent {
		t.Fatalf("File initial content should be equal to Expected: %s, Actual: %s", fileContent, openedFile.InitialContent)
	}
	if openedFile.ActualContent != fileContent {
		t.Fatalf("File actual content should be equal to Expected: %s, Actual: %s", fileContent, openedFile.ActualContent)
	}
	if !openedFile.Opened {
		t.Fatalf("File after save should still be opened")
	}
}

func TestApplicationUpdateFileInfo(t *testing.T) {
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

	createdEditor := CreateEditor(manager)

	testFilePath := createTestFile()
	openErr := createdEditor.OpenFile(testFilePath)
	if openErr != nil {
		t.Fatalf("Open of the test file failed")
	}

	openedFile, getOpErr := createdEditor.GetOpenedFile()
	if getOpErr != nil {
		t.Fatalf("Failed to get opened file")
	}
	if openedFile.Path != testFilePath {
		t.Fatalf("Opened file is wrong. Expected: %s, Actual: %s", testFilePath, openedFile.Path)
	}
	if openedFile.New {
		t.Fatalf("File that was opened from existing file can't be new")
	}

	updInfErr := createdEditor.UpdateFileInformation(openedFile.Id, api.FileInfoUpdateStruct{
		Id:        openedFile.Id,
		Type:      "Plain Text",
		Extension: ".txt",
	})
	if updInfErr != nil {
		t.Fatalf("Failed to update file information")
	}
	if !openedFile.New {
		t.Fatalf("File should be new after changing extension or type")
	}
	if openedFile.Type != "Plain Text" {
		t.Fatalf("File type wasn't changed to proper one")
	}
	if openedFile.Extension != ".txt" {
		t.Fatalf("File extension wasn't changed to proper one")
	}
}
func createTestFile() string {
	filePrefix := strconv.FormatInt(time.Now().UnixNano(), 10)
	name := filePrefix + "*test_1.py"
	tmpFile, _ := os.CreateTemp("", name)
	expectedText := `import logging
import flask
from app.config import configure_flask

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s %(levelname)s [%(name)s %(funcName)s] %(message)s",
)
`
	_, _ = tmpFile.WriteString(expectedText)

	defer func(tmpFile *os.File) {
		err := tmpFile.Close()
		if err != nil {

		}
	}(tmpFile)
	return tmpFile.Name()
}
