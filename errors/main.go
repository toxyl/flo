package errors

import (
	"io/fs"

	"github.com/toxyl/errors"
)

var (
	ErrNoEncoderImplemented     = errors.Newf("no encoder implemented")
	ErrNoDecoderImplemented     = errors.Newf("no decoder implemented")
	ErrChecksumAlgorithmInvalid = func(name string) error { return errors.Newf("%s is not a valid checksum algorithm", name) }
	ErrFailedToOpenFile         = func(file string, err error) error { return errors.Newf("failed to open file %s: %v", file, err) }
	ErrFailedToCreateDir        = func(file string, err error) error { return errors.Newf("failed to create dir %s: %v", file, err) }
	ErrFailedToCreateFile       = func(file string, err error) error { return errors.Newf("failed to create file %s: %v", file, err) }
	ErrFailedToDeleteFile       = func(file string, err error) error { return errors.Newf("failed to delete %s: %v", file, err) }
	ErrFailedToCopyFile         = func(src, dst string, err error) error {
		return errors.Newf("failed copy %s to %s: %v", src, dst, err)
	}
	ErrFailedToSetPermissions = func(file string, mode fs.FileMode, err error) error {
		return errors.Newf("failed set %s to %s: %v", file, mode.String(), err)
	}
	ErrIsNotExecutable = func(file string) error {
		return errors.Newf("%s is not an executable, use PermExec(o, g, w) first", file)
	}
	ErrIsNotDirectory = func(file string) error {
		return errors.Newf("%s is not a directory", file)
	}
	ErrMustBePointer = func(target any) error { return errors.Newf("expected *%T, but got %T", target, target) }
)
