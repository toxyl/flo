package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/toxyl/flo"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Says hi using a bash script that is created on the fly, then made executable and run with the given args.")
		fmt.Println("Usage: " + filepath.Base(os.Args[0]) + " [your name]")
		return
	}

	script := `#!/bin/bash
echo "Hello $1"
`
	tmpFile := flo.File("/tmp/hello.sh")
	defer func() { _ = tmpFile.Remove() }()

	if err := tmpFile.WriteBytes([]byte(script)).PermExec(true, true, true).Exec(os.Args[1]); err != nil {
		fmt.Printf("Command failed:\n%s\n", err.Error())
		os.Exit(1)
	}
}
