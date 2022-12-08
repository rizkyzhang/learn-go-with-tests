package learnos_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCh(t *testing.T) {
	t.Run("Chdir", func(t *testing.T) {
		wd, err := os.Getwd()
		lastForwardSlashID := strings.LastIndex(wd, "/")
		parentDir := wd[:lastForwardSlashID]
		assert.NoError(t, err)

		targetDir := fmt.Sprintf("%s/learn_string", parentDir)
		err = os.Chdir("../learn_string")
		assert.NoError(t, err)

		wd, err = os.Getwd()
		assert.NoError(t, err)
		assert.Equal(t, targetDir, wd)
	})
}
