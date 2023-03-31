// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

/*
Package bindata provides a tool for embedding binary data into a Go executable.

Usage:

	import "github.com/jteeuwen/go-bindata"

	var asset bindata.Asset
	// Use asset.Name and asset.Func to access the asset data.

*/
package bindata

// Asset holds information about a single asset to be processed.
type Asset struct {
	FilePath string // Full file path of the asset.
	Name     string // Key used in the table of contents to reference the asset.
	Func     string // Function name that returns the asset contents.
}
