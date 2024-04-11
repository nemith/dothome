package dothome

import (
	"os"
	"path/filepath"
)

func NativeLayout() (Layout, error)                 { return windowsLayout() }
func CLILayout() (Layout, error)                    { return windowsLayout() }
func NativeAppLayout(app AppConfig) (Layout, error) { return windowsAppLayout(app) }
func CLIAppLayout(app AppConfig) (Layout, error)    { return windowsAppLayout(app) }

func windowsLayout() (Layout, error) {
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

func windowsAppLayout(appcfg AppConfig) (Layout, error) {
	layout, err := windowsLayout()
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
