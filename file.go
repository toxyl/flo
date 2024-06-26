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
func (f *FileObj) Info() *FileInfo                       { f.updateInfo(); return f.info }
func (f *FileObj) Exists() bool                          { return f.info.Exists }
func (f *FileObj) Depth() int                            { return strings.Count(f.Path(), string(filepath.Separator)) }
func (f *FileObj) NewerThan(t time.Time) bool            { return f.info.NewerThan(t) }
func (f *FileObj) OlderThan(t time.Time) bool            { return f.info.OlderThan(t) }
func (f *FileObj) String(lenOwner, lenGroup int) string  { return f.info.String(lenOwner, lenGroup) }
func (f *FileObj) Mkparent(perm fs.FileMode) error       { return f.Parent().Mkdir(perm) }
func (f *FileObj) Create(perm fs.FileMode) error {
	file, err := os.Create(f.Path())
	if log.Error(err, "could not create file %s", f.Path()) {
		return err
	}
	file.WriteString("new file")
	file.Close()
	return f.Perm(perm)
}

// Mklink creates a symlink at the given path pointing this file's path
func (f *FileObj) Mklink(path string) error {
	fSymlink := File(path)
	if fSymlink.Exists() {
		_ = fSymlink.Remove()
	}
	return os.Symlink(f.Path(), fSymlink.Path())
}

func (f *FileObj) Open() (file *os.File, closer func()) {
	file, err := os.OpenFile(f.Path(), os.O_RDWR, 0644)
	if log.Error(err, "could not open file %s", f.Path()) {
		return nil, nil
	}
	return file, func() { file.Close() }
}

func (f *FileObj) OpenReadOnly() (file *os.File, closer func()) {
	file, err := os.OpenFile(f.Path(), os.O_RDONLY, 0644)
	if log.Error(err, "could not open file %s", f.Path()) {
		return nil, nil
	}
	return file, func() { file.Close() }
}

func (f *FileObj) OpenWriteOnly() (file *os.File, closer func()) {
	file, err := os.OpenFile(f.Path(), os.O_WRONLY, 0644)
	if log.Error(err, "could not open file %s", f.Path()) {
		return nil, nil
	}
	return file, func() { file.Close() }
}

// OpenAppend opens the file for appending, creating the file if it doesn't exist.
func (f *FileObj) OpenAppend() (file *os.File, closer func()) {
	file, err := os.OpenFile(f.Path(), os.O_CREATE|os.O_APPEND, 0644)
	if log.Error(err, "could not open file %s", f.Path()) {
		return nil, nil
	}
	return file, func() { file.Close() }
}

// OpenTruncate opens the file for writing, creating the file if it doesn't exist. The file will be truncated if it exists.
func (f *FileObj) OpenTruncate() (file *os.File, closer func()) {
	file, err := os.OpenFile(f.Path(), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if log.Error(err, "could not open file %s", f.Path()) {
		return nil, nil
	}
	return file, func() { file.Close() }
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
	fp := f.Parent()
	fp.updateInfo()
	f.updateInfo()
	return utils.FileCopy(f.Path(), destinationPath, fp.FileMode(), f.FileMode())
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
