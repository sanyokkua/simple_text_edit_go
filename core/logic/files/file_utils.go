package files

import (
	"os"
	"simple_text_editor/core/apperrors"
)

func GetTextFromFile(filePath string) (string, error) {
	if len(filePath) == 0 {
		return "", apperrors.CreateError("file path is empty string or incorrect")
	}

	fileByteContent, err := os.ReadFile(filePath)

	if err != nil {
		return "", apperrors.WrapError(err, "error happened during reading file")
	}

	convertedBytesIntoString := string(fileByteContent)
	return convertedBytesIntoString, nil
}

func SaveTextToFile(filePath string, text string) error {
	if len(filePath) == 0 || len(text) == 0 {
		return apperrors.CreateError("text or file name is empty")
	}

	err := os.WriteFile(filePath, []byte(text), 0644)

	if err != nil {
		return apperrors.WrapError(err, "the error happened with writing text to file")
	}

	return nil
}
