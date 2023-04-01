package utils

import (
	"github.com/stretchr/testify/assert"
	"simple_text_editor/core/v3/types"
	"testing"
)

const TestPath = "../../../fileTypes.json"

func TestForEach(t *testing.T) {
	var emptySlice []string
	wasExecuted := false
	ForEach(emptySlice, func(index int, data *string) {
		wasExecuted = true
	})
	assert.False(t, wasExecuted, "For empty slice it is not expected that any executions of callback happen")

	notEmptySlice := []string{"Value1", "Value2", "Value3"}
	var countOfExecutions = 0
	var index1IsIncorrect = false
	var index2IsIncorrect = false
	var index3IsIncorrect = false
	ForEach(notEmptySlice, func(index int, data *string) {
		countOfExecutions++
		if index == 0 && *data != "Value1" {
			index1IsIncorrect = true
		}
		if index == 1 && *data != "Value2" {
			index2IsIncorrect = true
		}
		if index == 2 && *data != "Value3" {
			index3IsIncorrect = true
		}
	})

	assert.Equal(t, 3, countOfExecutions, "Number of executions is not valid")
	assert.False(t, index1IsIncorrect, "Indexed value is not correct")
	assert.False(t, index2IsIncorrect, "Indexed value is not correct")
	assert.False(t, index3IsIncorrect, "Indexed value is not correct")
}

func TestFindInSlice(t *testing.T) {
	var emptySlice []string
	wasExecuted := false
	indexEmpty := FindInSlice(emptySlice, func(value string) bool {
		wasExecuted = true
		return value == ""
	})
	assert.False(t, wasExecuted, "For empty slice it is not expected that any executions of callback happen")
	assert.Equal(t, -1, indexEmpty, "For empty slice it is not expected that any executions of callback happen")

	notEmptySlice := []string{"Value1", "Value2", "Value3"}

	index0 := FindInSlice(notEmptySlice, func(value string) bool {
		return "Value1" == value
	})
	index1 := FindInSlice(notEmptySlice, func(value string) bool {
		return "Value2" == value
	})
	index2 := FindInSlice(notEmptySlice, func(value string) bool {
		return "Value3" == value
	})

	assert.Equal(t, 0, index0, "Index of found entity is invalid")
	assert.Equal(t, 1, index1, "Index of found entity is invalid")
	assert.Equal(t, 2, index2, "Index of found entity is invalid")

}

func TestMapFileTypesJsonStructToTypesMap(t *testing.T) {
	structSlice := []types.FileTypesJsonStruct{{
		Key:        "key",
		Name:       "Type",
		Extensions: []types.FileTypeExtension{".txt", ".py"},
	}}
	typesMap, err := MapFileTypesJsonStructToTypesMap(structSlice)
	assert.Nil(t, err, "Should not return error")
	assert.NotNil(t, typesMap, "Should return map")
	jsonStruct := typesMap[types.FileTypeKey("key")]
	assert.Equal(t, "key", string(jsonStruct.Key))
	assert.Equal(t, "Type", jsonStruct.Name)
	assert.Equal(t, 2, len(jsonStruct.Extensions))
}

func TestReadFileTypesJsonWithEmptyPath(t *testing.T) {
	_, err := ReadFileTypesJson("")
	assert.NotNil(t, err, "Should return error")
}

func TestLoadFileTypesJsonWithCorrectPath(t *testing.T) {
	json, err := ReadFileTypesJson(TestPath)
	assert.Nil(t, err, "Should not return error")
	assert.NotNil(t, json, "Should return json struct")
	assert.True(t, len(json) > 0, "Should have several values in slice")
}
