// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package bindata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

func writeRestore(w io.Writer) error {
	_, _ = fmt.Fprintf(w, `
// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(path.Join(dir, path.Dir(name)), os.ModePerm); err != nil {
		return err
	}
	if err := ioutil.WriteFile(path.Join(dir, name), data, os.ModePerm); err != nil {
		return err
	}
	if err := os.Chtimes(path.Join(dir, name), time.Now(), time.Now()); err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil {
		return RestoreAsset(dir, name)
	}
	for _, child := range children {
		if err := RestoreAssets(dir, path.Join(name, child)); err != nil {
			return err
		}
	}
	return nil
}

func canonicalPath(name string) string {
	return strings.ReplaceAll(name, "\\", "/")
}
`)
	return nil
}

