package utils

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/toxyl/flo/errors"
)

func FileCopy(src, destinationPath string, dirPerm, filePerm fs.FileMode) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return errors.ErrFailedToOpenFile(src, err)
	}
	defer srcFile.Close()
	dir := filepath.Dir(destinationPath)
	if err := os.MkdirAll(dir, dirPerm); err != nil && !os.IsExist(err) {
		return errors.ErrFailedToCreateDir(dir, err)
	}

	dstFile, err := os.Create(destinationPath)
	if err != nil {
		return errors.ErrFailedToCreateFile(destinationPath, err)
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return errors.ErrFailedToCopyFile(srcFile.Name(), destinationPath, err)
	}

	if err := os.Chmod(destinationPath, filePerm); err != nil {
		return errors.ErrFailedToSetPermissions(destinationPath, filePerm, err)
	}

	return nil
}

func Mkfile(file string, dirPerm, filePerm fs.FileMode) (*os.File, error) {
	p := filepath.Dir(file)
	if err := os.MkdirAll(p, dirPerm); err != nil && !os.IsExist(err) {
		return nil, errors.ErrFailedToCreateDir(p, err)
	}

	return os.Create(file)
}

func GetFileModeL(path string) fs.FileMode {
	path = filepath.Clean(path)
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return fs.FileMode(0)
	}
	return fileInfo.Mode()
}

func GetFileMode(path string) fs.FileMode {
	path = filepath.Clean(path)
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fs.FileMode(0)
	}
	return fileInfo.Mode()
}
