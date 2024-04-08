package main

import (
	"fmt"
	"os"

	"github.com/toxyl/flo"
	"github.com/toxyl/flo/config"
	"github.com/toxyl/flo/log"
	"github.com/toxyl/flo/utils"
	"github.com/toxyl/glog"
)

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, ".")
	}
	// config.ColorMode = false
	log.SetFns(nil, nil) // errors can happen when accessing unknown files, just ignore them
	d := flo.Dir(os.Args[1])
	maxLenUID := 0
	maxLenGID := 0
	totalSize := 0
	directories := []*flo.DirObj{}
	files := []*flo.FileObj{}
	if !d.Info().Permissions.IsDir() {
		f := flo.File(d.Path())
		files = append(files, f)
		maxLenUID = len(f.Owner())
		maxLenGID = len(f.Group())
	}

	str := utils.NewString().Rune('\n').Str(glog.Underline() + d.Path() + glog.Reset()).Rune('\n').Rune('\n')

	d.EachLimit(
		func(f *flo.FileObj) {
			files = append(files, f)
			maxLenUID = glog.Max(maxLenUID, len(f.Owner()))
			maxLenGID = glog.Max(maxLenGID, len(f.Group()))
		},
		func(d *flo.DirObj) {
			directories = append(directories, d)
			maxLenUID = glog.Max(maxLenUID, len(d.Owner()))
			maxLenGID = glog.Max(maxLenGID, len(d.Group()))
		},
		0,
	)

	for _, d := range directories {
		str = str.Rune(' ').Str(d.String(maxLenUID, maxLenGID)).Rune('\n')
	}

	for _, f := range files {
		totalSize += int(f.Size())
		str = str.Rune(' ').Str(f.String(maxLenUID, maxLenGID)).Rune('\n')
	}
	str.Rune('\n').
		StrClean(!config.ColorMode,
			fmt.Sprintf("%s, %s (%s)\n",
				glog.IntAmount(len(directories), "directory", "directories"),
				glog.IntAmount(len(files), "file", "files"),
				glog.HumanReadableBytesIEC(totalSize))).
		Rune('\n').
		Print()

}
