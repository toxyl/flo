package flo

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/toxyl/flo/checksum"
	"github.com/toxyl/flo/config"
	"github.com/toxyl/flo/errors"
	"github.com/toxyl/flo/log"
	"github.com/toxyl/flo/ownership"
	"github.com/toxyl/flo/permissions"
	"github.com/toxyl/flo/utils"
)

type FileObj struct {
	path string
	info *FileInfo
}

func (f *FileObj) Name() string                          { return filepath.Base(f.path) }
func (f *FileObj) BaseDir() string                       { return filepath.Dir(f.path) }
func (f *FileObj) Path() string                          { return f.path }
func (f *FileObj) Size() int64                           { return f.info.Size }
func (f *FileObj) Owner() string                         { return f.info.Ownership.User() }
func (f *FileObj) Group() string                         { return f.info.Ownership.Group() }
func (f *FileObj) LastModified() time.Time               { return f.info.LastModified }
func (f *FileObj) FileMode() fs.FileMode                 { return f.info.Permissions.FileMode() }
func (f *FileObj) Permissions() *permissions.Permissions { return f.info.Permissions }
func (f *FileObj) Checksum() *checksum.Checksum          { f.info.Checksum.Update(); return f.info.Checksum }
func (f *FileObj) Info() *FileInfo                       { return f.info }
func (f *FileObj) Exists() bool                          { return f.info.Exists }
func (f *FileObj) Depth() int                            { return strings.Count(f.Path(), string(filepath.Separator)) }
func (f *FileObj) NewerThan(t time.Time) bool            { return f.info.NewerThan(t) }
func (f *FileObj) OlderThan(t time.Time) bool            { return f.info.OlderThan(t) }
func (f *FileObj) String(lenOwner, lenGroup int) string  { return f.info.String(lenOwner, lenGroup) }
func (f *FileObj) Mkparent(perm fs.FileMode) error       { return f.Parent().Mkdir(perm) }

// Mklink creates a symlink at the given path pointing this file's path
func (f *FileObj) Mklink(path string) error {
	fSymlink := File(path)
	if fSymlink.Exists() {
		_ = fSymlink.Remove()
	}
	return os.Symlink(f.Path(), fSymlink.Path())
}

func (f *FileObj) Open() *os.File {
	file, err := os.OpenFile(f.Path(), os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if log.Error(err, "could not open file %s", f.Path()) {
		return nil
	}
	return file
}

func (f *FileObj) Remove() error {
	defer f.updateInfo()
	err := os.RemoveAll(f.Path())
	if f.Exists() {
		return errors.ErrFailedToDeleteFile(f.Path(), err)
	}
	return nil
}

func (f *FileObj) Copy(destinationPath string) error {
	return utils.FileCopy(f.Path(), destinationPath, f.Parent().FileMode(), f.FileMode())
}

func (f *FileObj) CopyFrom(file *FileObj) error {
	defer f.updateInfo()
	return file.Copy(f.Path())
}

func newFile(path string) *FileObj {
	pabs, _ := filepath.Abs(path)
	f := &FileObj{
		path: pabs,
		info: &FileInfo{
			Name:         path,
			Mode:         0,
			LastModified: time.Time{},
			Exists:       false,
			Size:         0,
			Checksum:     checksum.New(config.ChecksumAlgorithm, path),
			Path:         path,
			Permissions:  permissions.New(path),
			Ownership:    ownership.New(path),
		},
	}
	f.updateInfo()

	return f
}

func File(path string) *FileObj { return newFile(path) }
