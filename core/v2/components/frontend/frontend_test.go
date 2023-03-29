package frontend

import (
	"context"
	"simple_text_editor/core/v2/api"
	"simple_text_editor/core/v2/components/app"
	"simple_text_editor/core/v2/components/dialogs"
	"simple_text_editor/core/v2/components/events"
	"testing"
)

func TestCreateFrontendApiPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	CreateFrontendApi(nil, nil, nil, nil)
}

func TestGetFileTypes(t *testing.T) {
	extensions := make(map[string]api.FileTypesJsonStruct, 1)
	extensions["py"] = api.FileTypesJsonStruct{
		Key:        "python",
		Name:       "python",
		Extensions: []string{"py"},
	}
	extensions["txt"] = api.FileTypesJsonStruct{
		Key:        "txt",
		Name:       "Plain Text",
		Extensions: []string{"txt"},
	}
	createdEditor := app.CreateEditor(&extensions)
	var provider api.ContextProvider = func() (ctx context.Context) {
		return nil
	}
	eventsApi := events.CreateEvents(provider)
	dialogsApi := dialogs.CreateDialogs(provider, &extensions)
	frontendApi := CreateFrontendApi(createdEditor, eventsApi, dialogsApi, &extensions)

	types := frontendApi.GetFileTypes()
	if len(types) != 2 {
		t.Fatalf("Incorrect number of created types")
	}

	for i, pairStruct := range types {
		if pairStruct.Key == "python" {
			if types[i].Value != "python" {
				t.Fatalf("Incorrect VALUE in type. Expected: %s, Actual: %s", "python", types[i].Value)
			}
		} else if pairStruct.Key == "txt" {
			if types[i].Value != "Plain Text" {
				t.Fatalf("Incorrect VALUE in type. Expected: %s, Actual: %s", "Plain Text", types[i].Value)
			}
		} else {
			t.Fatalf("Key is not correct. Actual: %s", pairStruct.Key)
		}
	}
}

func TestGetFileTypeExtension(t *testing.T) {
	extensions := make(map[string]api.FileTypesJsonStruct, 1)
	extensions["py"] = api.FileTypesJsonStruct{
		Key:        "python",
		Name:       "python",
		Extensions: []string{"py"},
	}
	extensions["txt"] = api.FileTypesJsonStruct{
		Key:        "txt",
		Name:       "Plain Text",
		Extensions: []string{"txt", "rtf"},
	}
	createdEditor := app.CreateEditor(&extensions)
	var provider api.ContextProvider = func() (ctx context.Context) {
		return nil
	}
	eventsApi := events.CreateEvents(provider)
	dialogsApi := dialogs.CreateDialogs(provider, &extensions)
	frontendApi := CreateFrontendApi(createdEditor, eventsApi, dialogsApi, &extensions)

	extensions1 := frontendApi.GetFileTypeExtension("py")
	if len(extensions1) != 1 {
		t.Fatalf("Number of extensions for py should be 1. Actual: %d", len(extensions1))
	}
	extensions2 := frontendApi.GetFileTypeExtension("txt")
	if len(extensions2) != 2 {
		t.Fatalf("Number of extensions for txt should be 2. Actual: %d", len(extensions2))
	}
	extensions3 := frontendApi.GetFileTypeExtension("nonEx")
	if len(extensions3) != 0 {
		t.Fatalf("Number of extensions for non existing should be 0. Actual: %d", len(extensions3))
	}
}
