package flo

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/toxyl/flo/errors"
	"github.com/toxyl/flo/log"
	"github.com/toxyl/flo/permissions"
)

type DirObj struct {
	*FileObj
	dirs  []*DirObj
	files []*FileObj
}

func (f *FileObj) Mkdir(mode fs.FileMode) error {
	p := f.path
	if !f.Permissions().IsDir() {
		p = filepath.Dir(p)
	}
	if err := os.MkdirAll(p, mode); err != nil && !os.IsExist(err) {
		return err
	}
	defer f.updateInfo()
	return nil
}

func (d *DirObj) Contents() *DirObj {
	if !d.Permissions().IsDir() {
		log.Error(errors.ErrIsNotDirectory(d.Path()), "Retrieving contents failed!")
		return d
	}
	files := []*FileObj{}
	dirs := []*DirObj{}
	contents, err := os.ReadDir(d.Path())
	if err != nil {
		return d
	}
	for _, file := range contents {
		path := filepath.Join(d.Path(), file.Name())
		p := permissions.New(path)
		if p.IsDir() {
			dirs = append(dirs, Dir(path))
			continue
		}
		files = append(files, File(path))
	}
	d.dirs = dirs
	d.files = files
	sort.Slice(d.dirs, func(i, j int) bool {
		return strings.ToLower(d.dirs[i].Name()) < strings.ToLower(d.dirs[j].Name())
	})
	sort.Slice(d.files, func(i, j int) bool {
		return strings.ToLower(d.files[i].Name()) < strings.ToLower(d.files[j].Name())
	})
	return d
}

func (d *DirObj) Dirs() []*DirObj {
	return d.dirs
}

func (d *DirObj) Files() []*FileObj {
	return d.files
}

func (d *DirObj) walk(fnFile func(f *FileObj), fnDir func(d *DirObj), depth, maxDepth int) {
	if maxDepth >= 0 && depth > maxDepth {
		return
	}

	contents := d.Contents()

	for _, dir := range contents.Dirs() {
		if fnDir != nil {
			fnDir(dir)
		}
		dir.walk(fnFile, fnDir, depth+1, maxDepth)
	}

	for _, file := range contents.Files() {
		if fnFile != nil {
			fnFile(file)
		}
	}
}

func (d *DirObj) Each(fnFile func(f *FileObj), fnDir func(d *DirObj)) {
	d.walk(fnFile, fnDir, 0, -1)
}

func (d *DirObj) EachLimit(fnFile func(f *FileObj), fnDir func(d *DirObj), maxDepth int) {
	d.walk(fnFile, fnDir, 0, maxDepth)
}

func newDir(path string) *DirObj {
	d := &DirObj{
		FileObj: newFile(path),
		dirs:    []*DirObj{},
		files:   []*FileObj{},
	}
	return d
}

func Dir(path string) *DirObj { return newDir(path) }
