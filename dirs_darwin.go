package dothome

import (
	"os"
	"path/filepath"
)

// NativeLayout returns the base directories for the standards on the native os.
//
// On darwin it will use the following:
//   - Home: $HOME
//   - ConfigDir: $HOME/Library/Preferences
//   - DataDir: $HOME/Library/Application Support
//   - CacheDir: $HOME/Library/Caches
func NativeLayout() (Layout, error) { return appleLayout() }

// CLILayout returns the base directories that are suitable for CLI applications.
//
// On darwin this uses the same output as unix system and is as follows:
//   - Home: $HOME
//   - ConfigDir: $XDG_CONFIG_HOME or $HOME/.config
//   - DataDir: $XDG_DATA_HOME or $HOME/.local/share
//   - CacheDir: $XDG_CACHE_HOME or $HOME/.cache
func CLILayout() (Layout, error) { return xdgLayout() }

// NativeAppLayout returns application specific directories for the native os.
//
// The resulting directories will be based on [AppConfig.AppleBundleID]. If it
// is missing than the [AppConfig.Name] will be used instead.
//
//	layout, err := dothome.NativeAppLayout(dothome.AppConfig{Name: "MyApp"})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(layout.ConfigDir)
//	// Output: /home/user/Library/Preferences/myapp/config
//	fmt.Println(layout.DataDir)
//	// Ouput: /home/user/Library/Application Support/myapp/data
//	fmt.Println(layout.CacheDir)
//	// Ouput: /home/user/Library/Caches/myapp
func NativeAppLayout(app AppConfig) (Layout, error) { return appleAppLayout(app) }

// CLIAppLayout returns application specific directories that are better suited
// for CLI applications.
//
// For darwin this uses a unix XDG layout even following the same rules as unix
// with the environment variables.
//
//	  layout, err := dothome.NativeAppLayout(dothome.AppConfig{Name: "MyApp"})
//	  if err != nil {
//	      log.Fatal(err)
//	  }
//	  fmt.Println(layout.ConfigDir)
//	  // Output: /home/user/.config/my-app
//	  fmt.Println(layout.DataDir)
//	  // Ouput: /home/user/.local/share/my-app
//	  fmt.Println(layout.CacheDir)
//	// Ouput: /home/user/.cache/my-app

func CLIAppLayout(app AppConfig) (Layout, error) { return xdgAppLayout(app) }

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
		ConfigDir: filepath.Join(layout.ConfigDir, bundleID),
		DataDir:   filepath.Join(layout.DataDir, bundleID),
		CacheDir:  filepath.Join(layout.CacheDir, bundleID),
	}, nil
}
