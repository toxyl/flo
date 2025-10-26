package flo

import (
	"embed"
	"io/fs"
	"strings"

	"github.com/toxyl/errors"
)

// InitWithEmbeddedFS unpacks the given `embeddedFS` into this directory.
//
// It will first remove the directory if it exists and then create it
// to ensure that we start with a clean slate.
//
// If your embedded FS contains a path prefix (e.g. `//go:embed mysource/*`) and
// you don't want that replicated when extracting the FS, you can provide it
// as `pathPrefix` to strip it (e.g. `pathPrefix = "mysource"`).
//
// Be aware that embedded filesystems do not store file permissions.
// Therefore all dirs and files will be written with the provided
// permissions `dirMode` and `fileMode` respectively.
func (dir *DirObj) InitWithEmbeddedFS(embeddedFS embed.FS, pathPrefix string, dirMode, fileMode fs.FileMode) error {
	if dir.Exists() {
		// remove first, so we start with clean data
		if err := dir.Remove(); err != nil {
			return errors.Newf("removing base dir %s failed", dir.Path()).Append(err)
		}
	}
	if err := dir.Mkdir(dirMode); err != nil {
		return errors.Newf("creating base dir %s failed", dir.Path()).Append(err)
	}
	pathPrefix += "/"
	err := fs.WalkDir(embeddedFS, ".", func(path string, file fs.DirEntry, err error) error {
		if strings.HasPrefix(path, pathPrefix) {
			dst := dir.Dir(strings.TrimPrefix(path, pathPrefix))
			if file.IsDir() {
				if err := dst.Mkdir(dirMode); err != nil {
					return errors.Newf("creating dir %s failed", dst.Path()).Append(err)
				}
				return nil
			}

			data, err := embeddedFS.ReadFile(path)
			if err != nil {
				return errors.Newf("reading %s failed", path).Append(err)
			}
			if err := dst.StoreBytes(data); err != nil {
				return errors.Newf("writing %s failed", dst.Path()).Append(err)
			}
			if err := dst.Perm(fileMode); err != nil {
				return errors.Newf("setting permissions on %s failed", dst.Path()).Append(err)
			}
		}
		return nil
	})
	if err != nil {
		return errors.Newf("can't walk embedded dir").Append(err)
	}
	return nil
}

// UpdateFromEmbeddedFS works similar to InitWithEmbeddedFS but will not clear the directory before extracting the embedded FS.
//
// Existing files will only be overwritten with the version in the embedded FS if you set `overwriteExisting` to `true`.
func (dir *DirObj) UpdateFromEmbeddedFS(embeddedFS embed.FS, pathPrefix string, dirMode, fileMode fs.FileMode, overwriteExisting bool) error {
	if !dir.Exists() {
		return dir.InitWithEmbeddedFS(embeddedFS, pathPrefix, dirMode, fileMode)
	}
	pathPrefix += "/"
	err := fs.WalkDir(embeddedFS, ".", func(path string, file fs.DirEntry, err error) error {
		if strings.HasPrefix(path, pathPrefix) {
			dst := dir.Dir(strings.TrimPrefix(path, pathPrefix))
			if dst.Exists() && !overwriteExisting {
				return nil // silently ignore this file
			}
			if file.IsDir() {
				if err := dst.Mkdir(dirMode); err != nil {
					return errors.Newf("creating dir %s failed", dst.Path()).Append(err)
				}
				return nil
			}
			data, err := embeddedFS.ReadFile(path)
			if err != nil {
				return errors.Newf("reading %s failed", path).Append(err)
			}
			if err := dst.Remove(); err != nil {
				return errors.Newf("removing %s failed", dst.Path()).Append(err)
			}
			if err := dst.StoreBytes(data); err != nil {
				return errors.Newf("writing %s failed", dst.Path()).Append(err)
			}
			if err := dst.Perm(fileMode); err != nil {
				return errors.Newf("setting permissions on %s failed", dst.Path()).Append(err)
			}
		}
		return nil
	})
	if err != nil {
		return errors.Newf("can't walk embedded dir").Append(err)
	}
	return nil
}
