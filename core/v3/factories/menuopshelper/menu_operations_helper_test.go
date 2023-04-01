package menuopshelper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateIMenuOpsHelper(t *testing.T) {

}

func TestCreateIMenuOpsHelperPanic(t *testing.T) {
	defer func() {
		r := recover()
		require.NotNil(t, r, "Code didn't panic")
	}()

	CreateIMenuOpsHelper(nil, nil, nil, nil)
}

func TestMenuHelperOperationsStruct_CloseApplication(t *testing.T) {

}

func TestMenuHelperOperationsStruct_CloseFile(t *testing.T) {

}

func TestMenuHelperOperationsStruct_NewFile(t *testing.T) {

}

func TestMenuHelperOperationsStruct_OpenFile(t *testing.T) {

}

func TestMenuHelperOperationsStruct_SaveFile(t *testing.T) {

}

func TestMenuHelperOperationsStruct_SaveFileAs(t *testing.T) {

}

func TestMenuHelperOperationsStruct_ShowFileInfoModal(t *testing.T) {

}

func TestMenuHelperOperationsStruct_saveFile(t *testing.T) {

}
