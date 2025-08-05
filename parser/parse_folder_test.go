package parser_test

import (
	"testing"

	"github.com/lucaspopp0/mpy-bundler/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFolder(t *testing.T) {
	result, err := parser.ParseFolder("testdata")
	require.NoError(t, err)

	require.Len(t, result.Folders, 1)
	assert.Equal(t, "app", result.Folders[0].Name)
	require.Len(t, result.Folders[0].Files, 1)
	assert.Equal(t, "a.py", result.Folders[0].Files[0].Name)
	require.Len(t, result.Folders[0].Folders, 1)
	assert.Equal(t, "a", result.Folders[0].Folders[0].Name)
	require.Len(t, result.Folders[0].Folders[0].Files, 2)
	assert.Equal(t, "a.py", result.Folders[0].Folders[0].Files[0].Name)
	assert.Equal(t, "b.py", result.Folders[0].Folders[0].Files[1].Name)
}
