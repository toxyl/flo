package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/toxyl/flo"
	"github.com/toxyl/flo/log"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Prints a list of all files contained in the given directory up to the given maximum depth.")
		fmt.Println("Usage: " + filepath.Base(os.Args[0]) + " [directory] <maximum depth>")
		return
	}
	max := int64(-1)
	if len(os.Args) == 3 {
		m, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Println("Second argument must be a valid unsigned integer")
			return
		}
		max = m
	}

	log.SetFns(nil, nil) // errors can happen when accessing unknow files, just ignore them

	flo.Dir(os.Args[1]).EachLimit(
		func(f *flo.FileObj) {
			fmt.Printf("%s\n", f.Path())
		},
		nil,
		int(max),
	)
}
