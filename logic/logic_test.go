package logic

import (
	"os"
	"simple_text_editor/constants"
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

func TestMakeLines(t *testing.T) {
	line1 := "The First Line."
	line2 := "The Second Line."
	line3 := "The Third LINE!"
	givenText := line1 + "\n" + line2 + "\n" + line3

	lines, err := MakeLines(givenText)
	if err != nil {
		t.Errorf("Error happened. err: %s", err)
	}
	if len(lines) != 3 {
		t.Errorf("Number of lines is incorrect. Expected: %d, Actual: %d", 3, len(lines))
	}

	expectedLines := []string{
		line1, line2, line3,
	}
	for i, exp := range expectedLines {
		if lines[i] != exp {
			t.Errorf("Line %d is not valid. Expected: %s, Actual: %s", i, exp, lines[i])
		}
	}
}

func TestSortLines(t *testing.T) {
	linesToSort := []string{
		"w",
		"A",
		"c",
		"C",
	}
	expectedLines := []string{
		"A",
		"C",
		"c",
		"w",
	}

	err := SortLines(linesToSort, constants.SORT_ASC)
	if err != nil {
		t.Errorf("Error happened during sorting. error: %s", err)
	}

	for i, val := range expectedLines {
		if val != linesToSort[i] {
			t.Errorf("Line is not valid. Expected: %s, Actual: %s", val, linesToSort[i])
		}
	}
}
