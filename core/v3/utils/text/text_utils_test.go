package text

import (
	"testing"
)

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

	err := SortLines(linesToSort, SortAsc)
	if err != nil {
		t.Errorf("Error happened during sorting. error: %s", err)
	}

	for i, val := range expectedLines {
		if val != linesToSort[i] {
			t.Errorf("Line is not valid. Expected: %s, Actual: %s", val, linesToSort[i])
		}
	}
}

func TestSortAsc(t *testing.T) {
	originalText := [7]string{
		"The first string that needs to be sorted",
		"the second string",
		"Awesome golang",
		"Work In Progress",
		"witcher is a great game",
		"boring tests",
		"Core i5 was my CPU",
	}
	sortedText := [7]string{
		"Awesome golang",
		"Core i5 was my CPU",
		"The first string that needs to be sorted",
		"Work In Progress",
		"boring tests",
		"the second string",
		"witcher is a great game",
	}

	sortAsc(originalText[:])

	for i := 0; i < 7; i++ {
		if originalText[i] != sortedText[i] {
			t.Errorf("Sorting is not working, expected: %s, actual: %s", sortedText[i], originalText[i])
			break
		}
	}
}

func TestSortAscIgnoreCase(t *testing.T) {
	originalText := [7]string{
		"The first string that needs to be sorted",
		"the second string",
		"Awesome golang",
		"Work In Progress",
		"witcher is a great game",
		"boring tests",
		"Core i5 was my CPU",
	}
	sortedText := [7]string{
		"Awesome golang",
		"boring tests",
		"Core i5 was my CPU",
		"The first string that needs to be sorted",
		"the second string",
		"witcher is a great game",
		"Work In Progress",
	}

	sortAscIgnoreCase(originalText[:])

	for i := 0; i < 7; i++ {
		if originalText[i] != sortedText[i] {
			t.Errorf("Sorting is not working, expected: %s, actual: %s", sortedText[i], originalText[i])
			break
		}
	}
}

func TestSortDesc(t *testing.T) {
	originalText := [7]string{
		"The first string that needs to be sorted",
		"the second string",
		"Awesome golang",
		"Work In Progress",
		"witcher is a great game",
		"boring tests",
		"Core i5 was my CPU",
	}
	sortedText := [7]string{
		"witcher is a great game",
		"the second string",
		"boring tests",
		"Work In Progress",
		"The first string that needs to be sorted",
		"Core i5 was my CPU",
		"Awesome golang",
	}

	sortDesc(originalText[:])

	for i := 0; i < 7; i++ {
		if originalText[i] != sortedText[i] {
			t.Errorf("Sorting is not working, expected: %s, actual: %s", sortedText[i], originalText[i])
			break
		}
	}
}

func TestSortDescIgnoreCase(t *testing.T) {
	originalText := [7]string{
		"The first string that needs to be sorted",
		"the second string",
		"Awesome golang",
		"Work In Progress",
		"witcher is a great game",
		"boring tests",
		"Core i5 was my CPU",
	}
	sortedText := [7]string{
		"Work In Progress",
		"witcher is a great game",
		"the second string",
		"The first string that needs to be sorted",
		"Core i5 was my CPU",
		"boring tests",
		"Awesome golang",
	}

	sortDescIgnoreCase(originalText[:])

	for i := 0; i < 7; i++ {
		if originalText[i] != sortedText[i] {
			t.Errorf("Sorting is not working, expected: %s, actual: %s", sortedText[i], originalText[i])
			break
		}
	}
}
