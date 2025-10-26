package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/toxyl/flo"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Executes the given file with the given args.")
		fmt.Println("Usage: " + filepath.Base(os.Args[0]) + " [file] <arg1> <arg2> .. <argN>")
		return
	}

	// build args list
	args := []interface{}{}
	for _, a := range os.Args[2:] {
		args = append(args, a)
	}

	// execute file with the args
	if err := flo.File(os.Args[1]).Exec(args...); err != nil {
		fmt.Printf("Command failed:\n%s\n", err.Error())
		os.Exit(1)
	}
}
