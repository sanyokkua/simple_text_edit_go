package validators

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"regexp"
	"strings"
)

func IsEmptyString(value string) bool {
	return len(value) == 0
}

func IsNil(value interface{}) bool {
	return value == nil
}

func IsNotNil(value interface{}) bool {
	return !IsNil(value)
}

func IsNilOrEmptySlice[T interface{}](slice []T) bool {
	return slice == nil || len(slice) == 0
}

func HasError(err error) bool {
	return err != nil
}

func IsValidExtension(value string) bool {
	if IsEmptyString(value) {
		return false
	}
	matched, err := regexp.MatchString(`^\.(\d\w)+$`, value)
	if err != nil {
		log.Warn("Passed value failed matching. Value: ", value)
		return false
	}
	return matched
}

func PanicOnNil(argument interface{}, argumentName string) {
	if IsNil(argument) {
		panic(fmt.Sprintf("%s should not be NIL", argumentName))
	}
}

func IsValidFileTypeKey(value string) bool {
	if IsEmptyString(value) {
		return false
	}

	if strings.HasPrefix(value, ".") || strings.HasSuffix(value, ".") {
		return false
	}

	matched, err := regexp.MatchString(`^(\d\w)+$`, value)
	if err != nil {
		log.Warn("Passed value failed matching. Value: ", value)
		return false
	}
	return matched
}
