//go:build unix && !darwin

package dothome

// NativeLayout returns the base directories for the standards on the native os.
//
// On unix it will use the following:
//   - Home: $HOME
//   - ConfigDir: $XDG_CONFIG_HOME or $HOME/.config
//   - DataDir: $XDG_DATA_HOME or $HOME/.local/share
//   - CacheDir: $XDG_CACHE_HOME or $HOME/.cache
func NativeLayout() (Layout, error) { return xdgLayout() }

// CLILayout returns the base directories that are suitable for CLI applications.
//
// On unix this is the same as [NativeLayout].
func CLILayout() (Layout, error) { return xdgLayout() }

// NativeAppLayout returns application specific directories for the native os.
//
// The resulting directories will be based on [AppConfig.Name] which will be
// lowercases and spaces replaced with hyphens.
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
func NativeAppLayout(app AppConfig) (Layout, error) { return xdgAppLayout(app) }

// CLIAppLayout returns application specific directories that are better suited
// for CLI applications.
//
// For unix this is the same as [NativeAppLayout].
func CLIAppLayout(app AppConfig) (Layout, error) { return xdgAppLayout(app) }
