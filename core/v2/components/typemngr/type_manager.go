package typemngr

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"simple_text_editor/core/v2/api"
	"strings"
)

type typeManager struct {
	mapping map[string]*api.FileTypesJsonStruct
}

func (r *typeManager) GetTypeStructByKey(key string) *api.FileTypesJsonStruct {
	jsonStruct, ok := r.mapping[key]

	if !ok {
		return nil
	}

	return jsonStruct
}
func (r *typeManager) GetTypeStructByExt(extension string) *api.FileTypesJsonStruct {
	reMapped := make(map[string]*api.FileTypesJsonStruct, len(r.mapping))

	for _, jsonStruct := range r.mapping {
		for _, ext := range jsonStruct.Extensions {
			if strings.HasPrefix(ext, ".") {
				reMapped[ext] = jsonStruct
			} else {
				reMapped[fmt.Sprintf(".%s", ext)] = jsonStruct
			}
		}
	}

	return reMapped[extension]
}
func (r *typeManager) GetTypeKeyByExtension(extension string) string {
	typeInfo := r.GetTypeStructByExt(extension)
	if typeInfo == nil {
		return ""
	}

	return typeInfo.Key
}
func (r *typeManager) GetExtensionsForType(key string) []string {
	typeStructByKey := r.GetTypeStructByKey(key)
	if typeStructByKey == nil {
		return []string{}
	}

	extensions := make([]string, 0, len(typeStructByKey.Extensions)) //TODO:
	for _, ext := range r.mapping[key].Extensions {
		if strings.HasPrefix(ext, ".") {
			extensions = append(extensions, ext)
		} else {
			extensions = append(extensions, fmt.Sprintf(".%s", ext))
		}
	}
	return extensions
}
func (r *typeManager) GetSupportedFileFilters() []runtime.FileFilter {
	fileFilters := make([]runtime.FileFilter, 0, len(r.mapping))

	for _, value := range r.mapping {
		fileFilters = append(fileFilters, createFileFilter(value))
	}

	fileFilters = append(fileFilters, runtime.FileFilter{
		DisplayName: "Plain Text",
		Pattern:     "*.txt",
	})

	fileFilters = append(fileFilters, runtime.FileFilter{
		DisplayName: "Any File",
		Pattern:     "",
	})
	log.Info("GetSupportedFileFilters, return", fileFilters)
	return fileFilters
}
func (r *typeManager) BuildFileTypeMappingKeyToName() []api.KeyValuePairStruct {
	fileTypes := make([]api.KeyValuePairStruct, 0, len(r.mapping))

	for _, jsonStruct := range r.mapping {
		fileTypes = append(fileTypes, api.KeyValuePairStruct{
			Key:   jsonStruct.Key,
			Value: jsonStruct.Name,
		})
	}

	return fileTypes
}
func (r *typeManager) BuildFileTypeMappingExtToExt(fileTypeKey string) []api.KeyValuePairStruct {
	extensions := r.GetExtensionsForType(fileTypeKey)
	if extensions == nil || len(extensions) == 0 {
		return []api.KeyValuePairStruct{}
	}

	extToReturn := make([]api.KeyValuePairStruct, 0, len(extensions))
	for _, extension := range extensions {

		var ext string
		if strings.HasPrefix(extension, ".") {
			ext = extension
		} else {
			ext = fmt.Sprintf(".%s", extension)
		}

		extToReturn = append(extToReturn, api.KeyValuePairStruct{
			Key:   ext,
			Value: ext,
		})
	}

	return extToReturn
}
func createFileFilter(fileTypeInfo *api.FileTypesJsonStruct) runtime.FileFilter {
	typePattern := "*.%s;"
	pattern := ""

	for _, value := range fileTypeInfo.Extensions {
		pattern += fmt.Sprintf(typePattern, value)
	}

	return runtime.FileFilter{
		DisplayName: fileTypeInfo.Name,
		Pattern:     pattern,
	}
}

func CreateTypeManager(config []api.FileTypesJsonStruct) api.ITypeManager {
	mapping := make(map[string]*api.FileTypesJsonStruct, len(config))
	for _, jsonStruct := range config {
		key := jsonStruct.Key
		value := jsonStruct
		mapping[key] = &value
	}
	return &typeManager{mapping: mapping}
}
