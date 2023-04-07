package uniqueidgen

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGeneratorStruct_GenerateId(t *testing.T) {
	generator := CreateIUniqueIdGenerator()

	id1 := generator.GenerateId()
	id2 := generator.GenerateId()

	require.NotNilf(t, id1, "Result can't be nil")
	require.NotNilf(t, id2, "Result can't be nil")
	require.NotEqual(t, id1, id2, "IDs should be different")
}
