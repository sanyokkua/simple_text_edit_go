package logic

import (
	"errors"
	"github.com/labstack/gommon/log"
	"os"
	"simple_text_editor/constants"
	"strings"
)

func GetTextFromFile(filePath string) (string, error) {
	if len(filePath) == 0 {
		errorMessage := "file path is empty string or incorrect"
		log.Error(errorMessage, filePath)
		return "", errors.New(errorMessage)
	}

	fileByteContent, err := os.ReadFile(filePath)
	if err != nil {
		errorMessage := "read file function finished with error"
		log.Error(errorMessage, err)
	}

	convertedBytesIntoString := string(fileByteContent)
	return convertedBytesIntoString, nil
}

func SaveTextToFile(filePath string, text string) error {
	if len(filePath) == 0 || len(text) == 0 {
		errorMessage := "text or file name is empty"
		log.Error(errorMessage, filePath, text)
		err := errors.New(errorMessage)
		return err
	}

	err := os.WriteFile(filePath, []byte(text), 0644)
	if err != nil {
		errorMessage := "the error happened with writing text to file"
		log.Error(errorMessage, filePath, text)
		return err
	}

	return nil
}

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

func SortLines(lines []string, order int) error {
	if len(lines) == 0 {
		errorMessage := "text lines are empty array, nothing to sort"
		err := errors.New(errorMessage)
		log.Error(err)
		return err
	}
	if len(lines) == 1 {
		return nil
	}

	switch order {
	case constants.SORT_ASC:
		sortAsc(lines)
	case constants.SORT_ASC_IGNORE_CASE:
		sortAscIgnoreCase(lines)
	case constants.SORT_DESC:
		sortDesc(lines)
	case constants.SORT_DESC_IGNORE_CASE:
		sortDescIgnoreCase(lines)
	default:
		errorMessage := "passed sorting params are not valid"
		err := errors.New(errorMessage)
		log.Error(err, order)
		return err
	}
	return nil
}
