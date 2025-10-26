package flo

import (
	"path/filepath"
)

func (f *FileObj) resolve(name string) *FileObj   { return newFile(filepath.Join(f.path, name)) }
func (f *FileObj) resolveDir(name string) *DirObj { return newDir(filepath.Join(f.path, name)) }
func (f *FileObj) File(name string) *FileObj      { return f.resolve(name) }
func (f *FileObj) Dir(name string) *DirObj        { return f.resolveDir(name) }
func (f *FileObj) Parent() *DirObj                { return Dir(filepath.Dir(f.path)) }
