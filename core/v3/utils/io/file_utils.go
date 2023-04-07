package io

import (
	"errors"
	"github.com/labstack/gommon/log"
	"os"
)

func GetTextFromFile(filePath string) (string, error) {
	log.Info("GetTextFromFile", filePath)
	if len(filePath) == 0 {
		log.Error("GetTextFromFile", "file path is empty string or incorrect")
		return "", errors.New("file path is empty string or incorrect")
	}

	fileByteContent, err := os.ReadFile(filePath)

	if err != nil {
		log.Error("GetTextFromFile", "error happened during reading file")
		return "", errors.New("error happened during reading file")
	}

	convertedBytesIntoString := string(fileByteContent)
	return convertedBytesIntoString, nil
}

func SaveTextToFile(filePath string, text string) error {
	log.Info("SaveTextToFile", filePath, text)
	if len(filePath) == 0 || len(text) == 0 {
		log.Error("SaveTextToFile", "text or file name is empty")
		return errors.New("text or file name is empty")
	}

	err := os.WriteFile(filePath, []byte(text), 0644)

	if err != nil {
		log.Error("SaveTextToFile", "the error happened with writing text to file")
		return errors.New("the error happened with writing text to file")
	}

	return nil
}
