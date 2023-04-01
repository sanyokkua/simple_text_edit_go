package editor

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateIEditor(t *testing.T) {

}

func TestCreateIEditorPanic(t *testing.T) {
	defer func() {
		r := recover()
		require.NotNil(t, r, "Code didn't panic")
	}()

	CreateIEditor(nil, nil)
}

func TestEditorStruct_CloseFile(t *testing.T) {

}

func TestEditorStruct_CreateFileAndShow(t *testing.T) {

}

func TestEditorStruct_GetFileById(t *testing.T) {

}

func TestEditorStruct_GetFilesShortInfo(t *testing.T) {

}

func TestEditorStruct_GetOpenedFile(t *testing.T) {

}

func TestEditorStruct_HideAllFiles(t *testing.T) {

}

func TestEditorStruct_OpenFileAndShow(t *testing.T) {

}

func TestEditorStruct_SaveFile(t *testing.T) {

}

func TestEditorStruct_ShowFile(t *testing.T) {

}

func TestEditorStruct_UpdateFileContent(t *testing.T) {

}

func TestEditorStruct_UpdateFileInformation(t *testing.T) {

}

func Test_mapFileStructToInfoStruct(t *testing.T) {

}
