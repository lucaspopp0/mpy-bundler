package parser_test

import (
	_ "embed"
	"testing"

	"github.com/lucaspopp0/mpy-bundler/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed testdata/app/a/a.py
	contents_a_a string

	//go:embed testdata/app/a/b.py
	contents_a_b string

	//go:embed testdata/app/a.py
	contents_a string
)

func TestParseFile(t *testing.T) {
	t.Run("a/a.py", func(t *testing.T) {
		results, err := parser.ParseFile(contents_a_a)
		require.NoError(t, err)

		require.Contains(t, results.Declarations, "x")
		assert.Equal(t, results.Declarations["x"].Kind, parser.DeclarationVariable)

		require.Contains(t, results.Declarations, "y")
		assert.Equal(t, results.Declarations["y"].Kind, parser.DeclarationVariable)

		require.Contains(t, results.Declarations, "swap")
		assert.Equal(t, results.Declarations["swap"].Kind, parser.DeclarationFunction)
	})

	t.Run("a/b.py", func(t *testing.T) {
		results, err := parser.ParseFile(contents_a_b)
		require.NoError(t, err)

		require.Contains(t, results.Declarations, "Swapper")
		assert.Equal(t, results.Declarations["Swapper"].Kind, parser.DeclarationClass)
	})

	t.Run("a.py", func(t *testing.T) {
		results, err := parser.ParseFile(contents_a)
		require.NoError(t, err)

		require.Contains(t, results.Declarations, "swap")
		assert.Equal(t, results.Declarations["swap"].Kind, parser.DeclarationFunction)

		require.Len(t, results.Imports, 2)
		assert.Equal(t, "a", results.Imports[0].Name())
		assert.Equal(t, "b", results.Imports[1].Name())
	})
}
