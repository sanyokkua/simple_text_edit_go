package files

import (
	"github.com/labstack/gommon/log"
	"os"
	"simple_text_editor/core/apperrors"
)

func GetTextFromFile(filePath string) (string, error) {
	log.Info("GetTextFromFile", filePath)
	if len(filePath) == 0 {
		log.Error("GetTextFromFile", "file path is empty string or incorrect")
		return "", apperrors.CreateError("file path is empty string or incorrect")
	}

	fileByteContent, err := os.ReadFile(filePath)

	if err != nil {
		log.Error("GetTextFromFile", "error happened during reading file")
		return "", apperrors.WrapError(err, "error happened during reading file")
	}

	convertedBytesIntoString := string(fileByteContent)
	return convertedBytesIntoString, nil
}

func SaveTextToFile(filePath string, text string) error {
	log.Info("SaveTextToFile", filePath, text)
	if len(filePath) == 0 || len(text) == 0 {
		log.Error("SaveTextToFile", "text or file name is empty")
		return apperrors.CreateError("text or file name is empty")
	}

	err := os.WriteFile(filePath, []byte(text), 0644)

	if err != nil {
		log.Error("SaveTextToFile", "the error happened with writing text to file")
		return apperrors.WrapError(err, "the error happened with writing text to file")
	}

	return nil
}
