package eventsender

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateIEventSender(t *testing.T) {

}

func TestCreateITypeManagerPanic(t *testing.T) {
	defer func() {
		r := recover()
		require.NotNil(t, r, "Code didn't panic")
	}()

	CreateIEventSender(nil)
}

func TestEventSenderStruct_SendErrorEvent(t *testing.T) {

}

func TestEventSenderStruct_SendNotificationEvent(t *testing.T) {

}

func TestEventSenderStruct_SendWarnEvent(t *testing.T) {

}
