package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/toxyl/flo"
	"github.com/toxyl/flo/log"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Prints a list of all files and directories contained in the given directory.")
		fmt.Println("Usage: " + filepath.Base(os.Args[0]) + " [directory]")
		return
	}

	log.SetFns(nil, nil) // errors can happen when accessing unknow files, just ignore them

	flo.Dir(os.Args[1]).EachLimit(
		func(f *flo.FileObj) {
			n := f.Name()
			if strings.Contains(n, " ") {
				n = "'" + n + "'"
			}
			fmt.Printf("%s ", n)
		},
		func(d *flo.DirObj) {
			n := d.Name()
			if strings.Contains(n, " ") {
				n = "'" + n + "'"
			}
			fmt.Printf("\033[1m%s\033[0m ", n)
		},
		0,
	)
	fmt.Println()
}
