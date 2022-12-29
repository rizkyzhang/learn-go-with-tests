package learnos_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMkdir(t *testing.T) {
	t.Run("Mkdir", func(t *testing.T) {
		err := os.Mkdir("test", os.ModePerm)

		assert.NoError(t, err)
		assert.DirExists(t, "test")

		t.Cleanup(func() {
			os.Remove("test")
		})
	})

	t.Run("MkdirAll", func(t *testing.T) {
		err := os.MkdirAll("test1/test2/test3", os.ModePerm)

		assert.NoError(t, err)
		assert.DirExists(t, "test1/test2/test3")

		t.Cleanup(func() {
			os.RemoveAll("test1")
		})
	})

	t.Run("MkdirTemp default pattern", func(t *testing.T) {
		wd, err := os.Getwd()
		assert.NoError(t, err)

		dir, err := os.MkdirTemp(wd, "temp")
		assert.NoError(t, err)
		assert.DirExists(t, dir)

		file := filepath.Join(dir, "tmpfile")
		err = os.WriteFile(file, []byte("content"), 0666)
		assert.NoError(t, err)
		assert.FileExists(t, file)

		t.Cleanup(func() {
			os.RemoveAll(dir)
		})
	})

	t.Run("MkdirTemp custom pattern", func(t *testing.T) {
		wd, err := os.Getwd()
		assert.NoError(t, err)

		dir, err := os.MkdirTemp(wd, "temp-*")
		assert.NoError(t, err)
		assert.DirExists(t, dir)

		file := filepath.Join(dir, "tmpfile")
		err = os.WriteFile(file, []byte("content"), 0666)
		assert.NoError(t, err)
		assert.FileExists(t, file)

		t.Cleanup(func() {
			os.RemoveAll(dir)
		})
	})
}
