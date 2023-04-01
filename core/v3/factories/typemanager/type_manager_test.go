package typemanager

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateITypeManager(t *testing.T) {

}

func TestCreateITypeManagerPanic(t *testing.T) {
	defer func() {
		r := recover()
		require.NotNil(t, r, "Code didn't panic")
	}()

	CreateITypeManager(nil)
}

func TestTypeManagerStruct_BuildFileTypeMappingExtToExt(t *testing.T) {

}

func TestTypeManagerStruct_BuildFileTypeMappingKeyToName(t *testing.T) {

}

func TestTypeManagerStruct_GetExtensionsForType(t *testing.T) {

}

func TestTypeManagerStruct_GetSupportedFileFilters(t *testing.T) {

}

func TestTypeManagerStruct_GetTypeKeyByExtension(t *testing.T) {

}

func TestTypeManagerStruct_GetTypeStructByExt(t *testing.T) {

}

func TestTypeManagerStruct_GetTypeStructByKey(t *testing.T) {

}

func Test_createFileFilter(t *testing.T) {

}
