package menuopshelper

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"simple_text_editor/core/v3/components/eventsender"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/types/mocks"
	"testing"
)

type TestDataStruct struct {
	mockContextProvider types.ContextProvider
	mockedContext       *mocks.MockContextInterface
	mockEventSender     *mocks.IEventSender
	mockDialogHelper    *mocks.IDialogHelper
	mockEditor          *mocks.IEditor
	menuOpsHelper       *types.IMenuOpsHelper
}

func createTestData(t *testing.T) *TestDataStruct {
	mockedContext := mocks.NewMockContextInterface(t)
	var mockContextProvider types.ContextProvider = func() (ctx context.Context) {
		return mockedContext
	}
	mockEventSender := mocks.NewIEventSender(t)
	mockDialogHelper := mocks.NewIDialogHelper(t)
	mockEditor := mocks.NewIEditor(t)
	RuntimeQuit = func(ctx context.Context) {
	}

	helper := CreateIMenuOpsHelper(mockContextProvider, mockEventSender, mockDialogHelper, mockEditor)

	return &TestDataStruct{
		mockedContext:       mockedContext,
		mockContextProvider: mockContextProvider,
		mockEventSender:     mockEventSender,
		mockDialogHelper:    mockDialogHelper,
		mockEditor:          mockEditor,
		menuOpsHelper:       &helper,
	}
}

func TestCreateIMenuOpsHelper(t *testing.T) {
	data := createTestData(t)
	helper := data.menuOpsHelper

	require.NotNil(t, helper, "Should not be nil")
}

func TestCreateIMenuOpsHelperPanic(t *testing.T) {
	defer func() {
		r := recover()
		require.NotNil(t, r, "Code didn't panic")
	}()

	CreateIMenuOpsHelper(nil, nil, nil, nil)
}

func TestMenuHelperOperationsStruct_CloseApplicationCanceled(t *testing.T) {
	data := createTestData(t)
	helper := *data.menuOpsHelper

	data.mockEditor.On("GetFilesShortInfo").Return([]types.FileInfoStruct{
		{
			Id:      1,
			Name:    "New",
			New:     true,
			Opened:  true,
			Changed: true,
		},
		{
			Id:      2,
			Name:    "New",
			New:     true,
			Opened:  false,
			Changed: false,
		},
		{
			Id:      3,
			Name:    "New",
			New:     true,
			Opened:  false,
			Changed: false,
		},
	}, nil)

	data.mockDialogHelper.
		On("OkCancelMessageDialog", mock.Anything, mock.Anything).
		Return(types.Button("Cancel"), nil)

	helper.CloseApplication()

	data.mockEditor.AssertExpectations(t)
	data.mockDialogHelper.AssertExpectations(t)
	data.mockEventSender.AssertExpectations(t)
	data.mockedContext.AssertExpectations(t)
}

func TestMenuHelperOperationsStruct_CloseApplicationErr(t *testing.T) {
	data := createTestData(t)
	helper := *data.menuOpsHelper

	data.mockEditor.
		On("GetFilesShortInfo").
		Return([]types.FileInfoStruct{
			{
				Id:      1,
				Name:    "New",
				New:     true,
				Opened:  true,
				Changed: true,
			},
			{
				Id:      2,
				Name:    "New",
				New:     true,
				Opened:  false,
				Changed: false,
			},
			{
				Id:      3,
				Name:    "New",
				New:     true,
				Opened:  false,
				Changed: false,
			},
		}, errors.New("err"))
	data.mockEventSender.
		On("SendErrorEvent", "Failed to get all files")

	helper.CloseApplication()

	data.mockEditor.AssertExpectations(t)
	data.mockDialogHelper.AssertExpectations(t)
	data.mockEventSender.AssertExpectations(t)
	data.mockedContext.AssertExpectations(t)
}

func TestMenuHelperOperationsStruct_CloseApplicationNoChanges(t *testing.T) {
	data := createTestData(t)
	helper := *data.menuOpsHelper

	data.mockEditor.
		On("GetFilesShortInfo").
		Return([]types.FileInfoStruct{
			{
				Id:      1,
				Name:    "New",
				New:     true,
				Opened:  true,
				Changed: false,
			},
			{
				Id:      2,
				Name:    "New",
				New:     true,
				Opened:  false,
				Changed: false,
			},
			{
				Id:      3,
				Name:    "New",
				New:     true,
				Opened:  false,
				Changed: false,
			},
		}, nil)

	helper.CloseApplication()

	data.mockEditor.AssertExpectations(t)
	data.mockDialogHelper.AssertExpectations(t)
	data.mockEventSender.AssertExpectations(t)
	data.mockedContext.AssertExpectations(t)
}

func TestMenuHelperOperationsStruct_CloseApplicationDialErr(t *testing.T) {
	data := createTestData(t)
	helper := *data.menuOpsHelper

	data.mockEditor.
		On("GetFilesShortInfo").
		Return([]types.FileInfoStruct{
			{
				Id:      1,
				Name:    "New",
				New:     true,
				Opened:  true,
				Changed: true,
			},
			{
				Id:      2,
				Name:    "New",
				New:     true,
				Opened:  false,
				Changed: false,
			},
			{
				Id:      3,
				Name:    "New",
				New:     true,
				Opened:  false,
				Changed: false,
			},
		}, nil)
	data.mockDialogHelper.
		On("OkCancelMessageDialog",
			"Close application?",
			"Files in editor have changes. Do you want proceed and close all files? (Changes will not be saved)",
		).
		Return(types.Button("Ok"), errors.New("error"))
	data.mockEventSender.On("SendErrorEvent", "Failed to process Message dialog", mock.Anything)

	helper.CloseApplication()

	data.mockEditor.AssertExpectations(t)
	data.mockDialogHelper.AssertExpectations(t)
	data.mockEventSender.AssertExpectations(t)
	data.mockedContext.AssertExpectations(t)
}

func TestMenuHelperOperationsStruct_CloseFile(t *testing.T) {
	data := createTestData(t)
	helper := *data.menuOpsHelper

	data.mockEditor.
		On("GetOpenedFile").
		Return(&types.FileStruct{
			Id:               1,
			Path:             "path/file.txt",
			Name:             "file.txt",
			Type:             "txt",
			Extension:        ".txt",
			InitialExtension: ".txt",
			InitialContent:   "content",
			ActualContent:    "content",
			New:              false,
			Opened:           true,
			Changed:          true,
		}, nil)
	data.mockDialogHelper.
		On("OkCancelMessageDialog", "Close file? (file.txt)", "File has changes. Do you want proceed and close file? (Changes will not be saved)").
		Return(types.BtnOk, nil)
	data.mockEditor.
		On("CloseFile", int64(1)).
		Return(nil)
	data.mockEventSender.
		On("SendNotificationEvent", eventsender.EventOnFileClosed)

	helper.CloseFile()

	data.mockEditor.AssertExpectations(t)
	data.mockDialogHelper.AssertExpectations(t)
	data.mockEventSender.AssertExpectations(t)
	data.mockedContext.AssertExpectations(t)
}

func TestMenuHelperOperationsStruct_NewFile(t *testing.T) {

}

func TestMenuHelperOperationsStruct_OpenFile(t *testing.T) {

}

func TestMenuHelperOperationsStruct_SaveFile(t *testing.T) {

}

func TestMenuHelperOperationsStruct_SaveFileAs(t *testing.T) {

}

func TestMenuHelperOperationsStruct_saveFile(t *testing.T) {

}
