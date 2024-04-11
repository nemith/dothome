// Package path implements basic utilities for determining special applciation
// paths in a users home directory.
//
// This is essentiall an extension of [os.UserHomeDir], [os.UserConfigDir], and
// [os.UserCacheDir] functions.
//
// The main entry point to this package is the [NativeLayout] or [CLILayout]
// functions.
//
// [NativeLayout] returns a [Layout] object that uses the native OS paths for
// directory paths while [CLILayout] follows the same except on Apple where is
// chooses the [xdgLayout].  The, opinionated, idea is that CLI applications
// have files that are meant to be dealt with directly by users and discovery
// and editing of these files should be just like any other unix system.  You
// may not like this and that is ok, just use [NativeLayout].
//
//	layout, err := dothome.CLILayout()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	cfg, err := os.ReadFile(filepath.Join(layout.ConfigDir, "myapp", "config.json"))
package dothome
