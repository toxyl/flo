package errors

import (
	"fmt"
	"io/fs"

	"github.com/toxyl/errors"
)

var (
	ErrNoEncoderImplemented     = errors.Newf("no encoder implemented")
	ErrNoDecoderImplemented     = errors.Newf("no decoder implemented")
	ErrChecksumAlgorithmInvalid = func(name string) error { return errors.Newf("%s is not a valid checksum algorithm", name) }
	ErrFile                     = func(operation, file string, err error) error {
		if err != nil {
			return errors.Newf("failed to %s %s", operation, file).Append(err)
		}
		return nil
	}
	ErrFailedToOpenFile       = func(file string, err error) error { return ErrFile("open file", file, err) }
	ErrFailedToCreateDir      = func(file string, err error) error { return ErrFile("create dir", file, err) }
	ErrFailedToCreateFile     = func(file string, err error) error { return ErrFile("create file", file, err) }
	ErrFailedToDeleteFile     = func(file string, err error) error { return ErrFile("delete", file, err) }
	ErrFailedToCopyFile       = func(src, dst string, err error) error { return ErrFile(fmt.Sprintf("copy %s to", src), dst, err) }
	ErrFailedToSetPermissions = func(file string, mode fs.FileMode, err error) error {
		return ErrFile(fmt.Sprintf("set %s permissions on", mode.String()), file, err)
	}
	ErrIsNotExecutable = func(file string) error {
		return errors.Newf("%s is not an executable, use PermExec(o, g, w) first", file)
	}
	ErrIsNotDirectory = func(file string) error { return errors.Newf("%s is not a directory", file) }
	ErrMustBePointer  = func(target any) error { return errors.Newf("expected *%T, but got %T", target, target) }
)
