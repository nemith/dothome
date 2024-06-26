package dothome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWindowsLayout(t *testing.T) {
	t.Run("defaultPaths", func(t *testing.T) {
		t.Setenv("USERPROFILE", "C:\\Users\\user")
		t.Setenv("AppData", "")
		t.Setenv("LocalAppData", "")

		layout, err := windowsLayout()
		assert.NoError(t, err)

		assert.Equal(t, "C:\\Users\\user", layout.Home)
		assert.Equal(t, "C:\\Users\\user\\AppData\\Roaming", layout.ConfigDir)
		assert.Equal(t, "C:\\Users\\user\\AppData\\Roaming", layout.DataDir)
		assert.Equal(t, "C:\\Users\\user\\AppData\\Local", layout.CacheDir)
	})
	t.Run("envPaths", func(t *testing.T) {
		t.Setenv("USERPROFILE", "C:\\Users\\user")
		t.Setenv("AppData", "H:\\AppData\\Roaming")
		t.Setenv("LocalAppData", "H:\\AppData\\Local")

		layout, err := windowsLayout()
		assert.NoError(t, err)

		assert.Equal(t, "C:\\Users\\user", layout.Home)
		assert.Equal(t, "H:\\AppData\\Roaming", layout.ConfigDir)
		assert.Equal(t, "H:\\AppData\\Roaming", layout.DataDir)
		assert.Equal(t, "H:\\AppData\\Local", layout.CacheDir)
	})
}

func TestWindowsAppLayout(t *testing.T) {
	t.Setenv("USERPROFILE", "C:\\Users\\user")
	t.Setenv("AppData", "")
	t.Setenv("LocalAppData", "")

	layout, err := windowsAppLayout(testAppConfig)
	assert.NoError(t, err)

	assert.Equal(t, "C:\\Users\\user", layout.Home)
	assert.Equal(t, "C:\\Users\\user\\AppData\\Roaming\\MyOrg\\My App\\config", layout.ConfigDir)
	assert.Equal(t, "C:\\Users\\user\\AppData\\Roaming\\MyOrg\\My App\\data", layout.DataDir)
	assert.Equal(t, "C:\\Users\\user\\AppData\\Local\\MyOrg\\My App\\cache", layout.CacheDir)
}
