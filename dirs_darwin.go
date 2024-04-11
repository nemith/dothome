package dothome

import (
	"os"
	"path/filepath"
)

func NativeLayout() (Layout, error)                 { return appleLayout() }
func CLILayout() (Layout, error)                    { return xdgLayout() }
func NativeAppLayout(app AppConfig) (Layout, error) { return appleAppLayout(app) }
func CLIAppLayout(app AppConfig) (Layout, error)    { return xdgAppLayout(app) }

func appleLayout() (Layout, error) {
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

func appleAppLayout(appcfg AppConfig) (Layout, error) {
	layout, err := appleLayout()
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
