package menuhelper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateIMenuHelper(t *testing.T) {

}

func TestCreateIMenuHelperPanic(t *testing.T) {
	defer func() {
		r := recover()
		require.NotNil(t, r, "Code didn't panic")
	}()

	CreateIMenuHelper(nil, nil)
}

func TestMenuHelperStruct_CreateMenu(t *testing.T) {

}

func TestMenuHelperStruct_menuEditFileInfo(t *testing.T) {

}

func TestMenuHelperStruct_menuEditSort(t *testing.T) {

}

func TestMenuHelperStruct_menuFileCloseApp(t *testing.T) {

}

func TestMenuHelperStruct_menuFileCloseFile(t *testing.T) {

}

func TestMenuHelperStruct_menuFileNew(t *testing.T) {

}

func TestMenuHelperStruct_menuFileOpen(t *testing.T) {

}

func TestMenuHelperStruct_menuFileSave(t *testing.T) {

}

func TestMenuHelperStruct_menuFileSaveAs(t *testing.T) {

}
