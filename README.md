# dothome

[![Go Reference](https://pkg.go.dev/badge/github.com/nemith/dothome.svg)](https://pkg.go.dev/github.com/nemith/dothome)

`dothome` is a Go library for dealing with configuration, cache, and data
directories in a users home directory on various platforms.

It was highly inspired by Rust's [`etcetera`](https://docs.rs/etcetera/latest/etcetera/) where you can choose to override the native specification on Apple to use an XDG layout instead.  This was to be nicer to users running CLI apps who don't want their config files burried a pile of plists in a `$HOME/Library/Preferences` and prefer something like `.config/myapp/myapp.conf`.

## Basic Usage

```go
import "github.com/nemith/dothome"

layout, err := dothome.CLIAppLayout(dothome.AppConfig{Name: "foo"})
if err != nil {
    panic(err)
}
fmt.Println(layout.ConfigDir)
// Output: /home/user/.config/foo
```
