package dothome

import (
	"os"
	"path/filepath"
)

type Layout struct {
	Home      string
	ConfigDir string
	DataDir   string
	CacheDir  string
	// Missing:
	// XDG_STATE_DIR. No equivalent on macOS or Windows
	// XDG_RUNTIME_DIR. No equivalent on macOS or Windows
}

// AppConfig defines how to create application specific directories for
// [NativeAppLayout] or [CLIAppLayout].
type AppConfig struct {
	// Name is the name of the application. This field is mandatory.
	Name string

	// OrgName is the name of the organization that owns the application. This
	// is used for [WindowsAppLayout].
	OrgName string

	// AppleBundleID is the bundle ID of the application to explicitly use for
	// [AppleAppLayout].  If missing then the Name is used instead.
	AppleBundleID string
}

func envOrDefault(env, fallback string) string {
	path := os.Getenv(env)
	if path == "" || !filepath.IsAbs(path) {
		return fallback
	}
	return path
}
