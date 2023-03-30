package utils

import (
	"os"
	"testing"
)

const TestPath = "../../../fileTypes.json"

func TestLoadFileTypesJsonWithEmptyPath(t *testing.T) {
	res := LoadFileTypesJson("")

	if res == nil {
		t.Errorf("File Types slice should not be nil")
	}

	if len(res) > 0 {
		t.Errorf("File Types slice should have length = 0")
	}
}

func TestLoadFileTypesJsonWithCorrectPath(t *testing.T) {
	println(os.Getwd())

	res := LoadFileTypesJson(TestPath)

	if res == nil {
		t.Errorf("File Types slice should not be nil")
	}

	if len(res) == 0 {
		t.Errorf("File Types slice should have length > 0")
	}
}
