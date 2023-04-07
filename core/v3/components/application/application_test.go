package application

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateIApplication(t *testing.T) {

}

func TestCreateIApplicationPanic(t *testing.T) {
	defer func() {
		r := recover()
		require.NotNil(t, r, "Code didn't panic")
	}()

	CreateIApplication(nil)
}

func TestApplicationStruct_GetContext(t *testing.T) {

}

func TestApplicationStruct_GetFrontendApi(t *testing.T) {

}

func TestApplicationStruct_GetMenuApi(t *testing.T) {

}

func TestApplicationStruct_OnBeforeClose(t *testing.T) {

}

func TestApplicationStruct_OnDomReady(t *testing.T) {

}

func TestApplicationStruct_OnShutdown(t *testing.T) {

}

func TestApplicationStruct_OnStartup(t *testing.T) {

}
