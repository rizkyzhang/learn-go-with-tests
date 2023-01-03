package learnos_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		file, err := os.Create("test")
		defer file.Close()
		assert.NoError(t, err)
		assert.FileExists(t, file.Name())

		_, err = file.WriteString(file.Name())
		assert.NoError(t, err)

		t.Cleanup(func() {
			os.Remove(file.Name())
		})
	})

	t.Run("CreateTemp with os.Getwd()", func(t *testing.T) {
		wd, _ := os.Getwd()
		file, err := os.CreateTemp(wd, "")
		hasPrefixHome := strings.HasPrefix(file.Name(), "/home")
		defer file.Close()
		assert.NoError(t, err)
		assert.FileExists(t, file.Name())
		assert.Equal(t, true, hasPrefixHome)

		_, err = file.WriteString(file.Name())
		assert.NoError(t, err)

		t.Cleanup(func() {
			os.Remove(file.Name())
		})
	})

	t.Run("CreateTemp with dot shorthand", func(t *testing.T) {
		file, err := os.CreateTemp(".", "")
		hasPrefixDot := strings.HasPrefix(file.Name(), ".")
		defer file.Close()
		assert.NoError(t, err)
		assert.FileExists(t, file.Name())
		assert.Equal(t, true, hasPrefixDot)

		_, err = file.WriteString(file.Name())
		assert.NoError(t, err)

		t.Cleanup(func() {
			os.Remove(file.Name())
		})
	})

	t.Run("ReadFile", func(t *testing.T) {
		file, err := os.ReadFile("dummy")
		assert.NoError(t, err)
		assert.Equal(t, "This is a test", string(file))
	})

	t.Run("Copy file", func(t *testing.T) {
		src, err := os.Open("dummy")
		defer src.Close()
		assert.NoError(t, err)

		dst, err := os.Create("target")
		defer dst.Close()
		assert.NoError(t, err)

		_, err = io.Copy(dst, src)
		assert.NoError(t, err)

		content, err := os.ReadFile("target")
		assert.NoError(t, err)
		assert.Equal(t, "This is a test", string(content))

		t.Cleanup(func() {
			os.Remove("target")
		})
	})

	// t.Run("Rename", func(t *testing.T) {
	// 	err := os.Rename("renameme", "rename")
	// 	assert.NoError(t, err)

	// 	err = os.Rename("renameme_dir", "rename_dir")
	// 	assert.NoError(t, err)
	// })

	t.Run("WriteFile", func(t *testing.T) {
		err := os.WriteFile("writeme", []byte("Write success"), os.ModePerm)
		assert.NoError(t, err)

		content, err := os.ReadFile("writeme")
		assert.NoError(t, err)
		assert.Equal(t, "Write success", string(content))

		t.Cleanup(func() {
			os.WriteFile("writeme", []byte(""), os.ModePerm)
		})
	})

	t.Run("file.Write", func(t *testing.T) {
		file, err := os.OpenFile("writeme", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
		defer file.Close()
		assert.NoError(t, err)

		_, err = file.WriteString("wow\n")
		assert.NoError(t, err)
	})
}
