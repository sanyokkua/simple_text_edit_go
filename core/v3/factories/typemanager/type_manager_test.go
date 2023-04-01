package typemanager

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/utils"
	"testing"
)

func TestCreateITypeManager(t *testing.T) {
	typesMap := createTestData()
	manager := CreateITypeManager(typesMap)

	require.NotNil(t, manager, "Manager should be created")
}

func TestCreateITypeManagerPanic(t *testing.T) {
	defer func() {
		r := recover()
		require.NotNil(t, r, "Code didn't panic")
	}()

	CreateITypeManager(nil)
}

func TestTypeManagerStruct_BuildFileTypeMappingExtToExt(t *testing.T) {
	typesMap := createTestData()
	manager := CreateITypeManager(typesMap)

	typesSlice, err := manager.BuildFileTypeMappingKeyToName()
	require.Nil(t, err, "Should not return error")
	require.NotNil(t, typesSlice, "Should not be nil")
	assert.Equal(t, 2, len(typesSlice))

	utils.ForEach(typesSlice, func(_ int, data *types.KeyValuePairStruct) {
		if "python" == string(data.Key) {
			assert.Equal(t, "Python", data.Value)
		} else if "java" == string(data.Key) {
			assert.Equal(t, "Java", data.Value)
		} else {
			assert.Fail(t, "Incorrect value of key found")
		}
	})
}

func TestTypeManagerStruct_BuildFileTypeMappingKeyToName(t *testing.T) {
	typesMap := createTestData()
	manager := CreateITypeManager(typesMap)

	extensions, err := manager.BuildFileTypeMappingExtToExt("java")
	require.Nil(t, err, "Should not return error")
	require.NotNil(t, extensions, "Should not be nil")
	assert.Equal(t, 2, len(extensions))

	utils.ForEach(extensions, func(_ int, data *types.KeyValuePairStruct) {
		if ".java" == string(data.Key) {
			assert.Equal(t, ".java", data.Value)
		} else if ".jdk" == string(data.Key) {
			assert.Equal(t, ".jdk", data.Value)
		} else {
			assert.Fail(t, "Incorrect value of key found")
		}
	})
}

func TestTypeManagerStruct_GetExtensionsForType(t *testing.T) {
	typesMap := createTestData()
	manager := CreateITypeManager(typesMap)

	extensions1, err := manager.GetExtensionsForType("python")
	require.Nil(t, err, "Should not return error")
	assert.Equal(t, 1, len(extensions1), "Incorrect number of extensions")

	extensions2, err2 := manager.GetExtensionsForType("java")
	require.Nil(t, err2, "Should not return error")
	assert.Equal(t, 2, len(extensions2), "Incorrect number of extensions")

	_, err3 := manager.GetExtensionsForType("php")
	require.NotNil(t, err3, "Should return error")
}

func TestTypeManagerStruct_GetSupportedFileFilters(t *testing.T) {
	typesMap := createTestData()
	manager := CreateITypeManager(typesMap)

	filters := manager.GetSupportedFileFilters()
	require.NotNil(t, filters, "Filters shouldn't be nil")
	assert.Equal(t, 3, len(filters), "Number of filters is incorrect")
}

func TestTypeManagerStruct_GetTypeKeyByExtension(t *testing.T) {
	typesMap := createTestData()
	manager := CreateITypeManager(typesMap)

	type1, err := manager.GetTypeKeyByExtension(".py")
	require.Nil(t, err, "Should not return error")
	assert.Equal(t, "python", string(type1), "Incorrect number of extensions")

	type2, err2 := manager.GetTypeKeyByExtension(".java")
	require.Nil(t, err2, "Should not return error")
	assert.Equal(t, "java", string(type2), "Incorrect number of extensions")

	_, err3 := manager.GetExtensionsForType(".php")
	require.NotNil(t, err3, "Should return error")

}

func TestTypeManagerStruct_GetTypeStructByExt(t *testing.T) {
	typesMap := createTestData()

	manager := CreateITypeManager(typesMap)

	keyPython, err1 := manager.GetTypeStructByExt(".py")
	require.Nil(t, err1, "Python struct is not found")
	require.Equal(t, "python", string(keyPython.Key), "Key is not correct")
	require.Equal(t, "Python", keyPython.Name, "Name is not correct")
	require.Equal(t, 1, len(keyPython.Extensions), "Key is not correct")

	keyJava, err2 := manager.GetTypeStructByExt(".java")
	require.Nil(t, err2, "Java struct is not found")
	require.Equal(t, "java", string(keyJava.Key), "Key is not correct")
	require.Equal(t, "Java", keyJava.Name, "Name is not correct")
	require.Equal(t, 2, len(keyJava.Extensions), "Key is not correct")

	_, err3 := manager.GetTypeStructByExt(".php")
	require.NotNil(t, err3, "PHP struct should not be found")
}

func TestTypeManagerStruct_GetTypeStructByKey(t *testing.T) {
	typesMap := createTestData()

	manager := CreateITypeManager(typesMap)

	keyPython, err1 := manager.GetTypeStructByKey("python")
	require.Nil(t, err1, "Python struct is not found")
	require.Equal(t, "python", string(keyPython.Key), "Key is not correct")
	require.Equal(t, "Python", keyPython.Name, "Name is not correct")
	require.Equal(t, 1, len(keyPython.Extensions), "Key is not correct")

	keyJava, err2 := manager.GetTypeStructByKey("java")
	require.Nil(t, err2, "Java struct is not found")
	require.Equal(t, "java", string(keyJava.Key), "Key is not correct")
	require.Equal(t, "Java", keyJava.Name, "Name is not correct")
	require.Equal(t, 2, len(keyJava.Extensions), "Key is not correct")

	_, err3 := manager.GetTypeStructByKey("PHP")
	require.NotNil(t, err3, "PHP struct should not be found")
}

func Test_createFileFilter(t *testing.T) {
	fileType := types.FileTypesJsonStruct{
		Key:        "cpp",
		Name:       "C++",
		Extensions: []types.FileTypeExtension{"cpp", "hh", "c++"},
	}
	filter := createFileFilter(&fileType)
	if filter.DisplayName != fileType.Name {
		t.Errorf("Filter name is not correct. Expected: %s, Actual: %s", fileType.Name, filter.DisplayName)
	}
	if filter.Pattern != "*.cpp;*.hh;*.c++;" {
		t.Errorf("Filter pattern is not correct. Expected: %s, Actual: %s", "*.cpp;*.hh;*.c++;", filter.Pattern)
	}
}

func createTestData() types.TypesMap {
	f1 := types.FileTypesJsonStruct{
		Key:        "python",
		Name:       "Python",
		Extensions: []types.FileTypeExtension{".py"},
	}
	f2 := types.FileTypesJsonStruct{
		Key:        "java",
		Name:       "Java",
		Extensions: []types.FileTypeExtension{".java", ".jdk"},
	}
	typesMap := make(types.TypesMap)
	typesMap[f1.Key] = &f1
	typesMap[f2.Key] = &f2
	return typesMap
}
