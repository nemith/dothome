//go:build unix

package dothome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXDGLayout(t *testing.T) {
	t.Run("defaultPaths", func(t *testing.T) {
		t.Setenv("HOME", "/home/user")
		t.Setenv("XDG_CONFIG_HOME", "")
		t.Setenv("XDG_DATA_HOME", "")
		t.Setenv("XDG_CACHE_HOME", "")

		layout, err := xdgLayout()
		assert.NoError(t, err)

		assert.Equal(t, "/home/user", layout.Home)
		assert.Equal(t, "/home/user/.config", layout.ConfigDir)
		assert.Equal(t, "/home/user/.local/share", layout.DataDir)
		assert.Equal(t, "/home/user/.cache", layout.CacheDir)
	})

	t.Run("envPaths", func(t *testing.T) {
		t.Setenv("HOME", "/home/user")
		t.Setenv("XDG_CONFIG_HOME", "/home/user/Configs")
		t.Setenv("XDG_DATA_HOME", "/home/user/Data")
		t.Setenv("XDG_CACHE_HOME", "/home/user/Cache")

		layout, err := xdgLayout()
		assert.NoError(t, err)

		assert.Equal(t, "/home/user", layout.Home)
		assert.Equal(t, "/home/user/Configs", layout.ConfigDir)
		assert.Equal(t, "/home/user/Data", layout.DataDir)
		assert.Equal(t, "/home/user/Cache", layout.CacheDir)
	})
}

func TestXDGAppLayout(t *testing.T) {
	t.Setenv("HOME", "/home/user")
	t.Setenv("XDG_CONFIG_HOME", "")
	t.Setenv("XDG_DATA_HOME", "")
	t.Setenv("XDG_CACHE_HOME", "")

	layout, err := xdgAppLayout(testAppConfig)
	assert.NoError(t, err)

	assert.Equal(t, "/home/user", layout.Home)
	assert.Equal(t, "/home/user/.config/my-app", layout.ConfigDir)
	assert.Equal(t, "/home/user/.local/share/my-app", layout.DataDir)
	assert.Equal(t, "/home/user/.cache/my-app", layout.CacheDir)
}
