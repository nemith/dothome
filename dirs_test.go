package dothome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testAppConfig = AppConfig{
	Name:          "My App",
	OrgName:       "MyOrg",
	AppleBundleID: "com.myorg.MyApp",
}

var testSimpleAppConfig = AppConfig{Name: "My App"}

func TestAppleLayout(t *testing.T) {
	t.Setenv("HOME", "/Users/user")
	layout, err := AppleLayout()
	assert.NoError(t, err)

	assert.Equal(t, "/Users/user", layout.Home)
	assert.Equal(t, "/Users/user/Library/Preferences", layout.ConfigDir)
	assert.Equal(t, "/Users/user/Library/Application Support", layout.DataDir)
	assert.Equal(t, "/Users/user/Library/Caches", layout.CacheDir)
}

func TestAppleLayoutApp(t *testing.T) {
	t.Run("withBundleID", func(t *testing.T) {
		t.Setenv("HOME", "/Users/user")
		layout, err := AppleAppLayout(testAppConfig)
		assert.NoError(t, err)

		assert.Equal(t, "/Users/user", layout.Home)
		assert.Equal(t, "/Users/user/Library/Preferences/com.myorg.MyApp/config", layout.ConfigDir)
		assert.Equal(t, "/Users/user/Library/Application Support/com.myorg.MyApp/data", layout.DataDir)
		assert.Equal(t, "/Users/user/Library/Caches/com.myorg.MyApp", layout.CacheDir)
	})

	t.Run("appNameOnly", func(t *testing.T) {
		t.Setenv("HOME", "/Users/user")
		layout, err := AppleAppLayout(testSimpleAppConfig)
		assert.NoError(t, err)

		assert.Equal(t, "/Users/user", layout.Home)
		assert.Equal(t, "/Users/user/Library/Preferences/My App/config", layout.ConfigDir)
		assert.Equal(t, "/Users/user/Library/Application Support/My App/data", layout.DataDir)
		assert.Equal(t, "/Users/user/Library/Caches/My App", layout.CacheDir)
	})
}

func TestXDGLayout(t *testing.T) {
	t.Run("defaultPaths", func(t *testing.T) {
		t.Setenv("HOME", "/home/user")
		t.Setenv("XDG_CONFIG_HOME", "")
		t.Setenv("XDG_DATA_HOME", "")
		t.Setenv("XDG_CACHE_HOME", "")

		layout, err := XDGLayout()
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

		layout, err := XDGLayout()
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

	layout, err := XDGAppLayout(testAppConfig)
	assert.NoError(t, err)

	assert.Equal(t, "/home/user", layout.Home)
	assert.Equal(t, "/home/user/.config/my-app", layout.ConfigDir)
	assert.Equal(t, "/home/user/.local/share/my-app", layout.DataDir)
	assert.Equal(t, "/home/user/.cache/my-app", layout.CacheDir)
}
