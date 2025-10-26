package flo

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/toxyl/flo/errors"
)

func (f *FileObj) prepExecCommand(args ...any) (*exec.Cmd, error) {
	if !f.IsExecutable() {
		return nil, errors.ErrIsNotExecutable(f.Path())
	}
	strs := []string{}
	for _, a := range args {
		strs = append(strs, fmt.Sprintf("%s", a))
	}
	return exec.Command(f.Path(), strs...), nil
}

func (f *FileObj) Exec(args ...any) error {
	cmd, err := f.prepExecCommand(args...)
	if err != nil {
		return err
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (f *FileObj) ExecOutput(args ...any) ([]byte, error) {
	cmd, err := f.prepExecCommand(args...)
	if err != nil {
		return nil, err
	}
	return cmd.CombinedOutput()
}

func (f *FileObj) ExecQuiet(args ...any) error {
	cmd, err := f.prepExecCommand(args...)
	if err != nil {
		return err
	}
	return cmd.Run()
}

func (f *FileObj) ExecBackground(args ...any) error {
	cmd, err := f.prepExecCommand(args...)
	if err != nil {
		return err
	}
	return cmd.Start()
}
