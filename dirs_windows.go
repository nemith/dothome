package dothome

import (
	"os"
	"path/filepath"
)

// NativeLayout returns the base directories for the standards on the native os.
//
// On windows it will use the following:
//   - Home: %USERPROFILE%
//   - ConfigDir: %AppData% (or %USERPROFILE%\AppData\Roaming)
//   - DataDir: %AppData% (or %USERPROFILE%\AppData\Roaming)
//   - CacheDir: %LocalAppData% (or %USERPROFILE%\AppData\Local)
func NativeLayout() (Layout, error) { return windowsLayout() }

// CLILayout returns the base directories that are suitable for CLI applications.
//
// On Windows this is the same as [NativeLayout].
func CLILayout() (Layout, error) { return windowsLayout() }

// NativeAppLayout returns application specific directories for the native os.
//
//   layout, err := dothome.NativeAppLayout(dothome.AppConfig{Name: "MyApp", OrgName: "MyOrg"})
//   if err != nil {
//       log.Fatal(err)
//   }
//   fmt.Println(layout.ConfigDir)
//   // Output: C:\Users\user\AppData\Roaming\MyOrg\MyApp
//   fmt.Println(layout.DataDir)
//   // Ouput: C:\Users\user\AppData\Roaming\MyOrg\MyApp
//   fmt.Println(layout.CacheDir)
//   // Ouput: C:\Users\user\AppData\Local\MyOrg\MyApp

func NativeAppLayout(app AppConfig) (Layout, error) { return windowsAppLayout(app) }

// CLIAppLayout returns application specific directories that are better suited
// for CLI applications.
//
// For windows this is the same as [NativeAppLayout].
func CLIAppLayout(app AppConfig) (Layout, error) { return windowsAppLayout(app) }

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
		ConfigDir: filepath.Join(layout.ConfigDir, appcfg.OrgName, appcfg.Name, "config"),
		DataDir:   filepath.Join(layout.DataDir, appcfg.OrgName, appcfg.Name, "data"),
		CacheDir:  filepath.Join(layout.CacheDir, appcfg.OrgName, appcfg.Name, "cache"),
	}, nil
}
