package types

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestButton_EqualTo(t *testing.T) {
	var btn1 Button = "Ok"
	var btn2 Button = "Ok"
	var btn3 Button = "Cancel"

	assert.True(t, btn1.EqualTo(btn1))
	assert.True(t, btn2.EqualTo(btn2))
	assert.True(t, btn1.EqualTo(btn2))
	assert.True(t, btn2.EqualTo(btn1))
	assert.False(t, btn3.EqualTo(btn1))
	assert.False(t, btn3.EqualTo(btn2))
	assert.False(t, btn1.EqualTo(btn3))
	assert.False(t, btn2.EqualTo(btn3))
}

func TestExtensionsMap_GetByExtension(t *testing.T) {
	extensionsMap := make(ExtensionsMap)
	extensionsMap[FileTypeExtension(".jar")] = &FileTypesJsonStruct{
		Key:        "java",
		Name:       "Java",
		Extensions: []FileTypeExtension{".jar", ".java"},
	}

	_, err := extensionsMap.GetByExtension("ff")
	assert.NotNil(t, err, "Error should be returned on non existing mapping")

	value, err2 := extensionsMap.GetByExtension(".jar")
	require.Nil(t, err2, "Error should not be returned on existing mapping")
	require.NotNil(t, value, "Values should not be nil")
	assert.Equal(t, FileTypeKey("java"), value.Key)
	assert.Equal(t, "Java", value.Name)
	assert.Equal(t, 2, len(value.Extensions))
}

func TestFilesMap_Add(t *testing.T) {
	filesMap := make(FilesMap)
	file1 := FileStruct{
		Id:   0,
		Name: "New",
		New:  true,
	}
	file2 := FileStruct{
		Id:   1,
		Name: "New",
		New:  true,
	}

	err := filesMap.Add(&file1)
	require.Nil(t, err, "File should be added without errors")
	assert.Equal(t, 1, len(filesMap), "Incorrect number of files")

	err2 := filesMap.Add(&file2)
	require.Nil(t, err2, "File should be added without errors")
	assert.Equal(t, 2, len(filesMap), "Incorrect number of files")

	err3 := filesMap.Add(&file2)
	require.NotNil(t, err3, "File should not be added, file with id already exists")
	assert.Equal(t, 2, len(filesMap), "Incorrect number of files")
}

func TestFilesMap_GetById(t *testing.T) {
	filesMap := make(FilesMap)
	file1 := FileStruct{
		Id:   2222,
		Name: "New",
		New:  true,
	}
	file2 := FileStruct{
		Id:   3333,
		Name: "New",
		New:  true,
	}

	err := filesMap.Add(&file1)
	require.Nil(t, err, "File should be added without errors")

	err2 := filesMap.Add(&file2)
	require.Nil(t, err2, "File should be added without errors")

	_, errGet1 := filesMap.GetById(-100)
	require.NotNil(t, errGet1, "Should return error on non existing id in map")

	_, errGet2 := filesMap.GetById(999999)
	require.NotNil(t, errGet2, "Should return error on non existing id in map")

	actualFile1, errGet3 := filesMap.GetById(2222)
	require.Nil(t, errGet3, "Should return nil error on existing id in map")
	assert.Equal(t, file1, *actualFile1, "Files should be the same")

	actualFile2, errGet4 := filesMap.GetById(3333)
	require.Nil(t, errGet4, "Should return nil error on existing id in map")
	assert.Equal(t, file2, *actualFile2, "Files should be the same")
}

func TestFilesMap_IsPathPresentInMap(t *testing.T) {
	filesMap := make(FilesMap)
	file1 := FileStruct{
		Id:   2222,
		Name: "New",
		Path: "path/example/1",
		New:  true,
	}
	file2 := FileStruct{
		Id:   3333,
		Name: "New",
		Path: "diff/example/file.txt",
		New:  true,
	}
	file3 := FileStruct{
		Id:   5555,
		Name: "New",
		Path: "",
		New:  true,
	}

	err := filesMap.Add(&file1)
	require.Nil(t, err, "File should be added without errors")

	err2 := filesMap.Add(&file2)
	require.Nil(t, err2, "File should be added without errors")

	err3 := filesMap.Add(&file3)
	require.Nil(t, err3, "File should be added without errors")

	assert.False(t, filesMap.IsPathPresentInMap("NonExisting/path"), "Should return false on non existing path in map")
	assert.False(t, filesMap.IsPathPresentInMap(""), "Should return false on non empty path in map")
	assert.True(t, filesMap.IsPathPresentInMap("path/example/1"), "Should return false on non existing path in map")
	assert.True(t, filesMap.IsPathPresentInMap("diff/example/file.txt"), "Should return false on non existing path in map")
}

func TestFilesMap_Remove(t *testing.T) {
	filesMap := make(FilesMap)
	file1 := FileStruct{
		Id:   2222,
		Name: "New",
		Path: "path/example/1",
		New:  true,
	}
	file2 := FileStruct{
		Id:   3333,
		Name: "New",
		Path: "diff/example/file.txt",
		New:  true,
	}
	file3 := FileStruct{
		Id:   5555,
		Name: "New",
		Path: "",
		New:  true,
	}

	err := filesMap.Add(&file1)
	require.Nil(t, err, "File should be added without errors")

	err2 := filesMap.Add(&file2)
	require.Nil(t, err2, "File should be added without errors")

	err3 := filesMap.Add(&file3)
	require.Nil(t, err3, "File should be added without errors")

	require.Equal(t, 3, len(filesMap), "Initial state of the map before test is incorrect")

	removeErr1 := filesMap.Remove(-100)
	assert.NotNil(t, removeErr1, "For non existing ID error should be returned")
	require.Equal(t, 3, len(filesMap), "Amount of objects in map is incorrect")

	removeErr2 := filesMap.Remove(100)
	assert.NotNil(t, removeErr2, "For non existing ID error should be returned")
	require.Equal(t, 3, len(filesMap), "Amount of objects in map is incorrect")

	removeErr3 := filesMap.Remove(2222)
	assert.Nil(t, removeErr3, "For existing ID error should NOT be returned")
	require.Equal(t, 2, len(filesMap), "Amount of objects in map is incorrect")

	removeErr4 := filesMap.Remove(3333)
	assert.Nil(t, removeErr4, "For existing ID error should NOT be returned")
	require.Equal(t, 1, len(filesMap), "Amount of objects in map is incorrect")

	removeErr5 := filesMap.Remove(5555)
	assert.Nil(t, removeErr5, "For existing ID error should NOT be returned")
	require.Equal(t, 0, len(filesMap), "Amount of objects in map is incorrect")

	removeErr6 := filesMap.Remove(2222)
	assert.NotNil(t, removeErr6, "For non existing ID error should be returned")
	require.Equal(t, 0, len(filesMap), "Amount of objects in map is incorrect")
}

func TestTypesMap_GetByTypeKey(t *testing.T) {
	typesMap := make(TypesMap)
	typesMap[FileTypeKey("java")] = &FileTypesJsonStruct{
		Key:        "java",
		Name:       "Java",
		Extensions: []FileTypeExtension{".jar", ".java"},
	}

	_, err := typesMap.GetByTypeKey("ff")
	assert.NotNil(t, err, "Error should be returned on non existing mapping")

	value, err2 := typesMap.GetByTypeKey("java")
	require.Nil(t, err2, "Error should not be returned on existing mapping")
	require.NotNil(t, value, "Values should not be nil")
	assert.Equal(t, FileTypeKey("java"), value.Key)
	assert.Equal(t, "Java", value.Name)
	assert.Equal(t, 2, len(value.Extensions))
}
