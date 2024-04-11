package dothome

import (
	"os"
	"path/filepath"
	"strings"
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

// AppConfig defines how to create application specific directories.
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

// AppleAppLayout is the layout for macOS applications as defined by the [Apple
// File System Programming Guide].
//   - ConfigDir is $HOME/Library/Preferences
//   - DataDir is $HOME/Library/Application Support
//   - CacheDir is $HOME/Library/Caches
//
// [Apple File System Programming Guide]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/FileSystemOverview/FileSystemOverview.html
func AppleLayout() (Layout, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Layout{}, err
	}

	return Layout{
		Home:      home,
		ConfigDir: filepath.Join(home, "Library", "Preferences"),
		DataDir:   filepath.Join(home, "Library", "Application Support"),
		CacheDir:  filepath.Join(home, "Library", "Caches"),
	}, nil
}

// AppleAppLayout has application directories predefined for macOS applications.
// If a [AppConfig.AppleBundleID] is present in the appConfig it is used.  If not it falls
// back to the [AppConfig.Name] field with it's original casing.
//   - ConfigDir is $HOME/Library/Preferences/<bundleID>/config
//   - DataDir is $HOME/Library/Application Support/<bundleID>/data
//   - CacheDir is $HOME/Library/Caches/<bundleID>
func AppleAppLayout(appcfg AppConfig) (Layout, error) {
	layout, err := AppleLayout()
	if err != nil {
		return Layout{}, err
	}

	bundleID := appcfg.Name
	if appcfg.AppleBundleID != "" {
		bundleID = appcfg.AppleBundleID
	}

	return Layout{
		Home:      layout.Home,
		ConfigDir: filepath.Join(layout.ConfigDir, bundleID, "config"),
		DataDir:   filepath.Join(layout.DataDir, bundleID, "data"),
		CacheDir:  filepath.Join(layout.CacheDir, bundleID),
	}, nil
}

// WindowsLayout is the layout for Windows applications defined by the [Windows
// Known Folder Locations].
//   - ConfigDir is %AppData% or %HOME%\AppData\Roaming
//   - DataDir is %AppData% or %HOME%\AppData\Roaming
//   - CacheDir is %LocalAppData% or %HOME%\AppData\Local
//
// [Windows Known Folder Locations]: https://docs.microsoft.com/en-us/windows/win32/shell/knownfolderid.
func WindowsLayout() (Layout, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Layout{}, err
	}

	return Layout{
		Home:      home,
		ConfigDir: envOrDefault("AppData", filepath.Join(home, "AppData", "Roaming")),
		DataDir:   envOrDefault("AppData", filepath.Join(home, "AppData", "Roaming")),
		CacheDir:  envOrDefault("LocalAppData", filepath.Join(home, "AppData", "Local")),
	}, nil
}

func WindowsAppLayout(appcfg AppConfig) (Layout, error) {
	layout, err := WindowsLayout()
	if err != nil {
		return Layout{}, err
	}

	return Layout{
		Home:      layout.Home,
		ConfigDir: filepath.Join(layout.ConfigDir, appcfg.OrgName, appcfg.Name),
		DataDir:   filepath.Join(layout.DataDir, appcfg.OrgName, appcfg.Name),
		CacheDir:  filepath.Join(layout.CacheDir, appcfg.OrgName, appcfg.Name),
	}, nil
}

// XDGLayout is the layout for Unix-like systems as defined by the [XDG Base
// Directory Specification].
//   - ConfigDir is $XDG_CONFIG_HOME or $HOME/.config
//   - DataDir is $XDG_DATA_HOME or $HOME/.local/share
//   - CacheDir is $XDG_CACHE_HOME or $HOME/.cache
//
// [XDG Base Directory Specification]: https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html.
func XDGLayout() (Layout, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Layout{}, err
	}

	return Layout{
		Home:      home,
		ConfigDir: envOrDefault("XDG_CONFIG_HOME", filepath.Join(home, ".config")),
		DataDir:   envOrDefault("XDG_DATA_HOME", filepath.Join(home, ".local", "share")),
		CacheDir:  envOrDefault("XDG_CACHE_HOME", filepath.Join(home, ".cache")),
	}, nil
}

// XDGAppLayout has application directories predefined for Unix-like systems.
// The application name uses the [AppConfig.Name] field but spaces are replaced
// with hyphens and is lowerecased.
//   - ConfigDir is $XDG_CONFIG_HOME/<name>
//   - DataDir is $XDG_DATA_HOME/<name>
//   - CacheDir is $XDG_CACHE_HOME/<name>
func XDGAppLayout(appcfg AppConfig) (Layout, error) {
	layout, err := XDGLayout()
	if err != nil {
		return Layout{}, err
	}

	appName := strings.ToLower(appcfg.Name)
	appName = strings.ReplaceAll(appName, " ", "-")

	return Layout{
		Home:      layout.Home,
		ConfigDir: filepath.Join(layout.ConfigDir, appName),
		DataDir:   filepath.Join(layout.DataDir, appName),
		CacheDir:  filepath.Join(layout.CacheDir, appName),
	}, nil
}
