//go:build unix && !darwin

package dothome

func NativeLayout() (Layout, error)                 { return XDGLayout() }
func CLILayout() (Layout, error)                    { return XDGLayout() }
func NativeAppLayout(app AppConfig) (Layout, error) { return XDGAppLayout(app) }
func CLIAppLayout(app AppConfig) (Layout, error)    { return XDGAppLayout(app) }
