package constants

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	EventOnNewFileCreate string = "EventOnNewFileCreate"
	EventOnFileOpened    string = "EventOnFileOpened"
	EventOnErrorHappened string = "EventOnErrorHappened"
	EventOnFileSaved     string = "EventOnFileSaved"
	EventOnFileClosed    string = "EventOnFileClosed"
)

type FileTypeInformation struct {
	Key        string
	Name       string
	Extensions []string
}

func createFileTypeInformation(key string, name string, extensions ...string) FileTypeInformation {
	log.Info("createFileTypeInformation", key, name, extensions)
	return FileTypeInformation{
		Key:        key,
		Name:       name,
		Extensions: extensions,
	}
}

func getFileTypeInformation() []FileTypeInformation {
	log.Info("getFileTypeInformation")
	return []FileTypeInformation{
		createFileTypeInformation("c", "C", "c", "h", "cc", "hh", "C", "H"),
		createFileTypeInformation("cmake", "Cmake", "cmake"),
		createFileTypeInformation("cpp", "Cpp", "cpp", "cxx", "c++", "hpp", "hxx", "h++"),
		createFileTypeInformation("csharp", "Csharp", "cs", "csx"),
		createFileTypeInformation("css", "Css", "css"),
		createFileTypeInformation("dockerfile", "Dockerfile", "dockerfile"),
		createFileTypeInformation("go", "Golang", "go"),
		createFileTypeInformation("html", "Html", "html"),
		createFileTypeInformation("json", "Json", "json"),
		createFileTypeInformation("jsx", "Jsx", "jsx"),
		createFileTypeInformation("lua", "Lua", "lua"),
		createFileTypeInformation("groovy", "Groovy", "groovy", "gvy", "gy", "gsh"),
		createFileTypeInformation("java", "Java", "java", "jmod"),
		createFileTypeInformation("javascript", "Javascript", "js", "cjs", "mjs"),
		createFileTypeInformation("jinja2", "Jinja_2", "jinja2", "j2"),
		createFileTypeInformation("kotlin", "Kotlin", "kt", "kts", "ktm"),
		createFileTypeInformation("markdown", "Markdown", "markdown", "md"),
		createFileTypeInformation("objectivec", "Objective C", "m", "M", "mm"),
		createFileTypeInformation("php", "Php", "php"),
		createFileTypeInformation("powershell", "Powershell", "ps1"),
		createFileTypeInformation("properties", "Properties", "properties"),
		createFileTypeInformation("puppet", "Puppet", "pp"),
		createFileTypeInformation("python", "Python", "py", "pyi", "pyw"),
		createFileTypeInformation("ruby", "Ruby", "rb"),
		createFileTypeInformation("rust", "Rust", "rs"),
		createFileTypeInformation("scala", "Scala", "scala", "sc"),
		createFileTypeInformation("shell", "Shell", "sh"),
		createFileTypeInformation("sql", "Sql", "sql"),
		createFileTypeInformation("swift", "Swift", "swift", "SWIFT"),
		createFileTypeInformation("toml", "Toml", "toml"),
		createFileTypeInformation("tsx", "Tsx", "tsx"),
		createFileTypeInformation("typescript", "Typescript", "ts"),
		createFileTypeInformation("velocity", "Velocity", "vm", "vt"),
		createFileTypeInformation("xml", "Xml", "xml"),
		createFileTypeInformation("yaml", "Yaml", "yaml", "yml"),
	}
}

func createFileFilter(fileTypeInfo *FileTypeInformation) runtime.FileFilter {
	log.Info("createFileFilter", *fileTypeInfo)
	typePattern := "*.%s;"
	pattern := ""
	for _, value := range fileTypeInfo.Extensions {
		pattern += fmt.Sprintf(typePattern, value)
	}
	log.Info("createFileFilter", pattern)
	return runtime.FileFilter{
		DisplayName: fileTypeInfo.Name,
		Pattern:     pattern,
	}
}

func GetSupportedFileFilters() []runtime.FileFilter {
	log.Info("GetSupportedFileFilters")
	fileTypes := getFileTypeInformation()
	fileFilters := make([]runtime.FileFilter, 0, 10)
	for _, value := range fileTypes {
		fileFilters = append(fileFilters, createFileFilter(&value))
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

func GetExtToLangMapping() *map[string]string {
	log.Info("GetExtToLangMapping")
	info := getFileTypeInformation()
	mapping := make(map[string]string)
	for _, fileTypeInfo := range info {
		for _, ext := range fileTypeInfo.Extensions {
			mapping[ext] = fileTypeInfo.Key
		}
	}
	log.Info("GetExtToLangMapping, return", mapping)
	return &mapping
}
