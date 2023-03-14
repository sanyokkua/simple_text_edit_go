package logic

import "testing"

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
