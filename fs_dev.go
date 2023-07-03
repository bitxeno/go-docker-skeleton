//go:build dev

package main

import (
	"io/fs"
	"os"
)

func ViewAssets() fs.FS {
	return os.DirFS("view/dist")
}
