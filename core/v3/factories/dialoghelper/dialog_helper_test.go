package dialoghelper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateIDialogHelper(t *testing.T) {

}

func TestCreateIDialogHelperPanic(t *testing.T) {
	defer func() {
		r := recover()
		require.NotNil(t, r, "Code didn't panic")
	}()

	CreateIDialogHelper(nil, nil)
}

func TestDialogHelperStruct_OkCancelMessageDialog(t *testing.T) {

}

func TestDialogHelperStruct_OpenFileDialog(t *testing.T) {

}

func TestDialogHelperStruct_SaveFileDialog(t *testing.T) {

}
