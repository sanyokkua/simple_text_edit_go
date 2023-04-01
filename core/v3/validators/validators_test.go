package validators

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasError(t *testing.T) {
	hasErrorNil := HasError(nil)
	assert.False(t, hasErrorNil, "Should return false for nil value")

	err := errors.New("test error")
	hasError := HasError(err)
	assert.True(t, hasError, "Should return true for error obj")
}

func TestIsEmptyString(t *testing.T) {
	assert.False(t, IsEmptyString("Not empty"), "Should return false on non empty string")
	assert.True(t, IsEmptyString("   "), "Should return true on empty string")
	assert.True(t, IsEmptyString(""), "Should return true on empty string")
}

func TestIsNil(t *testing.T) {
	type test struct {
		field string
	}

	assert.True(t, IsNil(nil), "Should return true on nil value")
	assert.False(t, IsNil(test{
		field: "",
	}), "Should return false on non nil value")
}

func TestIsNilOrEmptySlice(t *testing.T) {
	var slice []string = nil
	assert.True(t, IsNilOrEmptySlice(slice), "Should return true on nil slice")
	assert.True(t, IsNilOrEmptySlice([]string{}), "Should return true on empty slice")
	assert.False(t, IsNilOrEmptySlice([]string{"Hello"}), "Should return false on slice with values")
}

func TestIsNotNil(t *testing.T) {
	type test struct {
		field string
	}

	assert.False(t, IsNotNil(nil), "Should return false on value")
	assert.True(t, IsNotNil(test{
		field: "",
	}), "Should return true on nil value")
}

func TestIsValidExtension(t *testing.T) {
	assert.True(t, IsValidExtension(".txt"), "Should return true on valid extension")
	assert.True(t, IsValidExtension(".c++"), "Should return true on valid extension")
	assert.True(t, IsValidExtension(".Java"), "Should return true on valid extension")
	assert.True(t, IsValidExtension(".C"), "Should return true on valid extension")
	assert.False(t, IsValidExtension(""), "Should return false on empty string")
	assert.False(t, IsValidExtension("    "), "Should return false on empty string")
	assert.False(t, IsValidExtension("    \n"), "Should return false on empty string")
	assert.False(t, IsValidExtension("    \t"), "Should return false on empty string")
	assert.False(t, IsValidExtension("txt"), "Should return false on string without '.' ")
	assert.False(t, IsValidExtension("._txt"), "Should return false on string with underscore")
}

func TestIsValidFileTypeKey(t *testing.T) {
	assert.True(t, IsValidFileTypeKey("txt"), "Should return true on valid type key")
	assert.True(t, IsValidFileTypeKey("c++"), "Should return true on valid type key")
	assert.True(t, IsValidFileTypeKey("Java"), "Should return true on valid type key")
	assert.True(t, IsValidFileTypeKey("C"), "Should return true on valid type key")
	assert.False(t, IsValidFileTypeKey(""), "Should return false on empty string")
	assert.False(t, IsValidFileTypeKey("    "), "Should return false on empty string")
	assert.False(t, IsValidFileTypeKey("    \n"), "Should return false on empty string")
	assert.False(t, IsValidFileTypeKey("    \t"), "Should return false on empty string")
	assert.False(t, IsValidFileTypeKey(".txt"), "Should return false on string without '.' ")
}

func TestPanicOnNil(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	var slice []string
	PanicOnNil(slice, "argument")
}

func TestPanicOnValue(t *testing.T) {
	defer func() {
		r := recover()
		if r != nil {
			t.Errorf("The code did panic")
		}
	}()
	PanicOnNil("string value", "argument")
}
