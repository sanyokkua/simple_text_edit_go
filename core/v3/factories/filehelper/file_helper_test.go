package filehelper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateIFileHelper(t *testing.T) {

}

func TestCreateITypeManagerPanic(t *testing.T) {
	defer func() {
		r := recover()
		require.NotNil(t, r, "Code didn't panic")
	}()

	CreateIFileHelper(nil, nil)
}

func TestFileHelperStruct_CreateNewFileEmpty(t *testing.T) {

}

func TestFileHelperStruct_CreateNewFileWithData(t *testing.T) {

}

func TestFileHelperStruct_GetFileExtensionFromPath(t *testing.T) {

}

func TestFileHelperStruct_GetFileNameFromPath(t *testing.T) {

}

func TestFileHelperStruct_GetFileTypeFromExtension(t *testing.T) {

}
