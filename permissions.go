package flo

import (
	"io/fs"
	"os"
	"os/user"
	"strconv"

	"github.com/toxyl/flo/log"
)

func (f *FileObj) Own(username string) error {
	u, err := user.Lookup(username)
	if err != nil {
		return err
	}
	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return err
	}
	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		return err
	}
	return os.Chown(f.path, uid, gid)
}

func (f *FileObj) Perm(mode fs.FileMode) error {
	f.updateInfo()
	defer f.updateInfo()
	return os.Chmod(f.path, mode)
}

func (f *FileObj) PermOwner(r, w, x bool) *FileObj {
	f.info.Permissions.Owner().Set(r, w, x)
	log.Error(f.Perm(f.info.Permissions.FileMode()), "Setting owner permissions on %s failed!", f.Path()) // aka +x
	return f
}

func (f *FileObj) PermGroup(r, w, x bool) *FileObj {
	f.info.Permissions.Group().Set(r, w, x)
	log.Error(f.Perm(f.info.Permissions.FileMode()), "Setting group permissions on %s failed!", f.Path()) // aka +x
	return f
}

func (f *FileObj) PermWorld(r, w, x bool) *FileObj {
	f.info.Permissions.World().Set(r, w, x)
	log.Error(f.Perm(f.info.Permissions.FileMode()), "Setting world permissions on %s failed!", f.Path()) // aka +x
	return f
}

func (f *FileObj) SetExec(owner, group, world bool) *FileObj {
	if owner {
		f.info.Permissions.Owner().SetExec()
	}
	if group {
		f.info.Permissions.Group().SetExec()
	}
	if world {
		f.info.Permissions.World().SetExec()
	}
	log.Error(f.Perm(f.info.Permissions.FileMode()), "Making %s executable failed!", f.Path()) // aka +x
	return f
}

func (f *FileObj) IsExecutable() bool      { return f.info.Permissions.World().HasExec() }
func (f *FileObj) IsExecutableGroup() bool { return f.info.Permissions.Group().HasExec() }
func (f *FileObj) IsExecutableOwner() bool { return f.info.Permissions.Owner().HasExec() }

func (f *FileObj) SetRead(owner, group, world bool) *FileObj {
	if owner {
		f.info.Permissions.Owner().SetRead()
	}
	if group {
		f.info.Permissions.Group().SetRead()
	}
	if world {
		f.info.Permissions.World().SetRead()
	}
	log.Error(f.Perm(f.info.Permissions.FileMode()), "Making %s readable failed!", f.Path()) // aka +x
	return f
}

func (f *FileObj) IsReadable() bool      { return f.info.Permissions.World().HasRead() }
func (f *FileObj) IsReadableGroup() bool { return f.info.Permissions.Group().HasRead() }
func (f *FileObj) IsReadableOwner() bool { return f.info.Permissions.Owner().HasRead() }

func (f *FileObj) SetWrite(owner, group, world bool) *FileObj {
	if owner {
		f.info.Permissions.Owner().SetWrite()
	}
	if group {
		f.info.Permissions.Group().SetWrite()
	}
	if world {
		f.info.Permissions.World().SetWrite()
	}
	log.Error(f.Perm(f.info.Permissions.FileMode()), "Making %s writable failed!", f.Path()) // aka +x
	return f
}

func (f *FileObj) IsWritable() bool      { return f.info.Permissions.World().HasWrite() }
func (f *FileObj) IsWritableGroup() bool { return f.info.Permissions.Group().HasWrite() }
func (f *FileObj) IsWritableOwner() bool { return f.info.Permissions.Owner().HasWrite() }
