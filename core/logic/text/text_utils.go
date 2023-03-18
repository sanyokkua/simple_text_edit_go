package text

import (
	"errors"
	"github.com/labstack/gommon/log"
	"sort"
	"strings"
)

type SortOrder = int

const (
	SortAsc SortOrder = iota + 1
	SortDesc
	SortAscIgnoreCase
	SortDescIgnoreCase
)

func MakeLines(text string) ([]string, error) {
	if len(text) == 0 {
		errorMessage := "text if empty and can't be split"
		log.Error(errorMessage)
		err := errors.New(errorMessage)
		return []string{}, err
	}

	splitText := strings.Split(text, "\n")
	return splitText, nil
}

func SortLines(lines []string, order SortOrder) error {
	if len(lines) == 0 {
		err := errors.New("text lines are empty array, nothing to sort")
		return err
	}

	if len(lines) == 1 {
		return nil
	}

	switch order {
	case SortAsc:
		sortAsc(lines)
	case SortAscIgnoreCase:
		sortAscIgnoreCase(lines)
	case SortDesc:
		sortDesc(lines)
	case SortDescIgnoreCase:
		sortDescIgnoreCase(lines)
	default:
		err := errors.New("passed sorting params are not valid")
		return err
	}
	return nil
}

func sortAsc(text []string) {
	sort.Strings(text)
}

func sortAscIgnoreCase(text []string) {
	sort.Slice(text, func(i, j int) bool {
		return strings.ToLower(text[i]) < strings.ToLower(text[j])
	})
}

func sortDesc(text []string) {
	sort.Slice(text, func(i, j int) bool {
		return text[i] > text[j]
	})
}

func sortDescIgnoreCase(text []string) {
	sort.Slice(text, func(i, j int) bool {
		return strings.ToLower(text[i]) > strings.ToLower(text[j])
	})
}
