package files

import (
	"os"
	"testing"
)

func TestGetTextFromFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "get_text_from_file_test")
	if err != nil {
		t.Fatalf("File was not created")
	}
	expectedText := `hbsr ub 3fbfbjvhb se\n
rver af
 af
aw f
ew
f a434iu buy3wyufbube`
	_, err2 := tmpFile.WriteString(expectedText)
	if err2 != nil {
		t.Fatalf("File data was not written")
	}

	textFromFile, err3 := GetTextFromFile(tmpFile.Name())

	if err3 != nil {
		t.Fatalf("error happen during getting text from file")
	}

	if expectedText != textFromFile {
		t.Errorf("Expected text is not the same with text from file. Expected: %s. actual: %s", expectedText, textFromFile)
	}

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Log(err)
		}
	}(tmpFile.Name())
}

func TestSaveTextToFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "get_text_from_file_test")
	if err != nil {
		t.Fatalf("File was not created. Err: %s", err)
	}

	expectedText := `Here it is a test text.
It is multiline text.
It is awesome.`

	err2 := SaveTextToFile(tmpFile.Name(), expectedText)
	if err2 != nil {
		t.Errorf("Text was not saved. Err: %s", err2)
	}

	actualText, err3 := GetTextFromFile(tmpFile.Name())
	if err3 != nil {
		t.Errorf("Text was not read. Err: %s", err3)
	}

	if expectedText != actualText {
		t.Errorf("Expected and Actual text are different. Expected %s\n Actual: %s", expectedText, actualText)
	}

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Log(err)
		}
	}(tmpFile.Name())
}
