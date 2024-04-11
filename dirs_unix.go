//go:build unix

package dothome

import (
	"os"
	"path/filepath"
	"strings"
)

func xdgLayout() (Layout, error) {
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

func xdgAppLayout(appcfg AppConfig) (Layout, error) {
	layout, err := xdgLayout()
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
