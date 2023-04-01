package frontendapi

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateIFrontendApi(t *testing.T) {

}

func TestCreateIFrontendApiPanic(t *testing.T) {
	defer func() {
		r := recover()
		require.NotNil(t, r, "Code didn't panic")
	}()

	CreateIFrontendApi(nil, nil, nil, nil)
}

func TestFrontendApiStruct_CloseFile(t *testing.T) {

}

func TestFrontendApiStruct_GetFileTypeExtension(t *testing.T) {

}

func TestFrontendApiStruct_GetFileTypes(t *testing.T) {

}

func TestFrontendApiStruct_GetFilesShortInfo(t *testing.T) {

}

func TestFrontendApiStruct_GetOpenedFile(t *testing.T) {

}

func TestFrontendApiStruct_SwitchOpenedFileTo(t *testing.T) {

}

func TestFrontendApiStruct_UpdateFileContent(t *testing.T) {

}

func TestFrontendApiStruct_UpdateFileInformation(t *testing.T) {

}
