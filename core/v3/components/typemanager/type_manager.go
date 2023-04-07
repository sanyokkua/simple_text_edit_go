package typemanager

import (
	"errors"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/v3/types"
	"simple_text_editor/core/v3/utils"
	"simple_text_editor/core/v3/validators"
)

const FilterExtensionPattern string = "*%s;"

var AnyFileFilter = runtime.FileFilter{
	DisplayName: "Any File",
	Pattern:     "",
}

func createFileFilter(typeInfo *types.FileTypesJsonStruct) runtime.FileFilter {
	pattern := ""

	utils.ForEach(typeInfo.Extensions, func(_ int, data *types.FileTypeExtension) {
		pattern += fmt.Sprintf(FilterExtensionPattern, *data)
	})

	return runtime.FileFilter{
		DisplayName: typeInfo.Name,
		Pattern:     pattern,
	}
}

type TypeManagerStruct struct {
	Mapping types.TypesMap
}

func (r *TypeManagerStruct) GetTypeStructByKey(key types.FileTypeKey) (*types.FileTypesJsonStruct, error) {
	if !validators.IsValidFileTypeKey(string(key)) {
		return nil, errors.New("passed key is not valid")
	}

	return r.Mapping.GetByTypeKey(key)
}

func (r *TypeManagerStruct) GetTypeStructByExt(extension types.FileTypeExtension) (*types.FileTypesJsonStruct, error) {
	if !validators.IsValidExtension(string(extension)) {
		return nil, errors.New("extension is not valid")
	}

	reMapped := make(types.ExtensionsMap, len(r.Mapping))

	for _, jsonStruct := range r.Mapping {
		jsStruct := jsonStruct

		utils.ForEach(jsStruct.Extensions, func(_ int, data *types.FileTypeExtension) {
			reMapped[*data] = jsStruct
		})
	}

	return reMapped.GetByExtension(extension)
}

func (r *TypeManagerStruct) GetTypeKeyByExtension(extension types.FileTypeExtension) (types.FileTypeKey, error) {
	typeStructByExt, err := r.GetTypeStructByExt(extension)
	if validators.HasError(err) {
		return "", err
	}

	return typeStructByExt.Key, nil
}

func (r *TypeManagerStruct) GetExtensionsForType(key types.FileTypeKey) ([]types.FileTypeExtension, error) {
	typeStructByKey, err := r.GetTypeStructByKey(key)
	if validators.HasError(err) {
		return nil, err
	}

	return typeStructByKey.Extensions, err
}

func (r *TypeManagerStruct) GetSupportedFileFilters() []runtime.FileFilter {
	fileFilters := make([]runtime.FileFilter, 0, len(r.Mapping))

	for _, value := range r.Mapping {
		fileFilters = append(fileFilters, createFileFilter(value))
	}

	fileFilters = append(fileFilters, AnyFileFilter)

	return fileFilters
}

func (r *TypeManagerStruct) BuildFileTypeMappingKeyToName() ([]types.KeyValuePairStruct, error) {
	fileTypes := make([]types.KeyValuePairStruct, 0, len(r.Mapping))

	for _, jsonStruct := range r.Mapping {
		fileTypes = append(fileTypes, types.KeyValuePairStruct{
			Key:   string(jsonStruct.Key),
			Value: jsonStruct.Name,
		})
	}

	return fileTypes, nil
}

func (r *TypeManagerStruct) BuildFileTypeMappingExtToExt(fileTypeKey types.FileTypeKey) ([]types.KeyValuePairStruct, error) {
	if validators.IsEmptyString(string(fileTypeKey)) {
		return nil, errors.New("file type key is empty string")
	}

	extensions, err := r.GetExtensionsForType(fileTypeKey)
	if validators.HasError(err) {
		return nil, err
	}

	extToReturn := make([]types.KeyValuePairStruct, 0, len(extensions))
	for _, extension := range extensions {
		extToReturn = append(extToReturn, types.KeyValuePairStruct{
			Key:   string(extension),
			Value: string(extension),
		})
	}

	return extToReturn, nil
}

func CreateITypeManager(typesMap types.TypesMap) types.ITypeManager {
	validators.PanicOnNil(typesMap, "TypesMap")

	return &TypeManagerStruct{Mapping: typesMap}
}
