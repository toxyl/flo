package errors

import (
	"fmt"
	"io/fs"
)

var (
	ErrNoEncoderImplemented     = fmt.Errorf("no encoder implemented")
	ErrNoDecoderImplemented     = fmt.Errorf("no decoder implemented")
	ErrChecksumAlgorithmInvalid = func(name string) error { return fmt.Errorf("%s is not a valid checksum algorithm", name) }
	ErrFailedToOpenFile         = func(file string, err error) error { return fmt.Errorf("failed to open file %s: %v", file, err) }
	ErrFailedToCreateDir        = func(file string, err error) error { return fmt.Errorf("failed to create dir %s: %v", file, err) }
	ErrFailedToCreateFile       = func(file string, err error) error { return fmt.Errorf("failed to create file %s: %v", file, err) }
	ErrFailedToDeleteFile       = func(file string, err error) error { return fmt.Errorf("failed to delete %s: %v", file, err) }
	ErrFailedToCopyFile         = func(src, dst string, err error) error {
		return fmt.Errorf("failed copy %s to %s: %v", src, dst, err)
	}
	ErrFailedToSetPermissions = func(file string, mode fs.FileMode, err error) error {
		return fmt.Errorf("failed set %s to %s: %v", file, mode.String(), err)
	}
	ErrIsNotExecutable = func(file string) error {
		return fmt.Errorf("%s is not an executable, use SetExec(o, g, w) first", file)
	}
	ErrIsNotDirectory = func(file string) error {
		return fmt.Errorf("%s is not a directory", file)
	}
	ErrMustBePointer = func(target any) error { return fmt.Errorf("expected *%T, but got %T", target, target) }
)
