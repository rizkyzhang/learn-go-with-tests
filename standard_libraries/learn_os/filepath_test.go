package learnos_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilepath(t *testing.T) {
	t.Run("Abs", func(t *testing.T) {
		absPath, err := filepath.Abs("..")
		assert.NoError(t, err)
		t.Logf("absPath: %s", absPath)
	})

	t.Run("Dir", func(t *testing.T) {
		result := filepath.Dir("test1/test2")
		assert.Equal(t, "test1", result)
	})

	t.Run("Dir with trailing slash", func(t *testing.T) {
		result := filepath.Dir("test1/test2/")
		assert.Equal(t, "test1/test2", result)
	})

	t.Run("Ext", func(t *testing.T) {
		result := filepath.Ext("test.png")
		assert.Equal(t, ".png", result)
	})

	t.Run("Split", func(t *testing.T) {
		dir, file := filepath.Split("/home/test/test.jpg")

		assert.Equal(t, "/home/test/", dir)
		assert.Equal(t, "test.jpg", file)
	})
}
