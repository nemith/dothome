package dothome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppleLayout(t *testing.T) {
	t.Setenv("HOME", "/Users/user")
	layout, err := appleLayout()
	assert.NoError(t, err)

	assert.Equal(t, "/Users/user", layout.Home)
	assert.Equal(t, "/Users/user/Library/Preferences", layout.ConfigDir)
	assert.Equal(t, "/Users/user/Library/Application Support", layout.DataDir)
	assert.Equal(t, "/Users/user/Library/Caches", layout.CacheDir)
}

func TestAppleLayoutApp(t *testing.T) {
	t.Run("withBundleID", func(t *testing.T) {
		t.Setenv("HOME", "/Users/user")
		layout, err := appleAppLayout(testAppConfig)
		assert.NoError(t, err)

		assert.Equal(t, "/Users/user", layout.Home)
		assert.Equal(t, "/Users/user/Library/Preferences/com.myorg.MyApp/config", layout.ConfigDir)
		assert.Equal(t, "/Users/user/Library/Application Support/com.myorg.MyApp/data", layout.DataDir)
		assert.Equal(t, "/Users/user/Library/Caches/com.myorg.MyApp", layout.CacheDir)
	})

	t.Run("appNameOnly", func(t *testing.T) {
		t.Setenv("HOME", "/Users/user")
		layout, err := appleAppLayout(testSimpleAppConfig)
		assert.NoError(t, err)

		assert.Equal(t, "/Users/user", layout.Home)
		assert.Equal(t, "/Users/user/Library/Preferences/My App/config", layout.ConfigDir)
		assert.Equal(t, "/Users/user/Library/Application Support/My App/data", layout.DataDir)
		assert.Equal(t, "/Users/user/Library/Caches/My App", layout.CacheDir)
	})
}
