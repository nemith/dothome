//go:build unix && !darwin

package dothome

// NativeLayout is the strategy for the native platform follwing defined practices.
// For Unix (non-Darwin) it uses the XDGLayout.
// For Darwin it uses the AppleLayout.
// For Windows it uses the WindowsLayout.
func NativeLayout() (Layout, error) { return XDGLayout() }

// CLILayout is an opinionated layout for CLI applications.  It follows the
// For Unix and Apple it uses the XDGLayout.  Apple won't use the AppleLayout.
// Fow Windows it uses continues to use the WindowsLayout.
func CLILayout() (Layout, error) { return XDGLayout() }
