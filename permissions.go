package flo

import (
	"io/fs"
	"os"
	"os/user"
	"strconv"

	"github.com/toxyl/flo/log"
)

func (f *FileObj) Own(username string) error {
	defer f.updateInfo()
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
	defer f.updateInfo()
	f.info.Permissions.Owner().Set(r, w, x)
	log.Error(f.Perm(f.info.Permissions.FileMode()), "Setting owner permissions on %s failed!", f.Path()) // aka +x
	return f
}

func (f *FileObj) PermGroup(r, w, x bool) *FileObj {
	defer f.updateInfo()
	f.info.Permissions.Group().Set(r, w, x)
	log.Error(f.Perm(f.info.Permissions.FileMode()), "Setting group permissions on %s failed!", f.Path()) // aka +x
	return f
}

func (f *FileObj) PermWorld(r, w, x bool) *FileObj {
	defer f.updateInfo()
	f.info.Permissions.World().Set(r, w, x)
	log.Error(f.Perm(f.info.Permissions.FileMode()), "Setting world permissions on %s failed!", f.Path()) // aka +x
	return f
}

func (f *FileObj) PermExec(owner, group, world bool) *FileObj {
	defer f.updateInfo()
	// implicitly we will also set +r permissions as they are required for +x to work
	if owner {
		f.info.Permissions.Owner().SetExec()
		f.info.Permissions.Owner().SetRead()
	} else {
		f.info.Permissions.Owner().ClearExec()
	}
	if group {
		f.info.Permissions.Group().SetExec()
		f.info.Permissions.Group().SetRead()
	} else {
		f.info.Permissions.Group().ClearExec()
	}
	if world {
		f.info.Permissions.World().SetExec()
		f.info.Permissions.World().SetRead()
	} else {
		f.info.Permissions.World().ClearExec()
	}
	log.Error(f.Perm(f.info.Permissions.FileMode()), "Making %s executable failed!", f.Path()) // aka +x
	return f
}

func (f *FileObj) PermExecAll() *FileObj {
	return f.PermExec(true, true, true)
}

func (f *FileObj) PermRead(owner, group, world bool) *FileObj {
	defer f.updateInfo()
	// implicitly we will also clear +x permissions as they don't work with +r
	if owner {
		f.info.Permissions.Owner().SetRead()
	} else {
		f.info.Permissions.Owner().ClearRead()
		f.info.Permissions.Owner().ClearExec()
	}
	if group {
		f.info.Permissions.Group().SetRead()
	} else {
		f.info.Permissions.Group().ClearRead()
		f.info.Permissions.Group().ClearExec()
	}
	if world {
		f.info.Permissions.World().SetRead()
	} else {
		f.info.Permissions.World().ClearRead()
		f.info.Permissions.World().ClearExec()
	}
	log.Error(f.Perm(f.info.Permissions.FileMode()), "Making %s readable failed!", f.Path())
	return f
}

func (f *FileObj) PermReadAll() *FileObj {
	return f.PermRead(true, true, true)
}

func (f *FileObj) PermWrite(owner, group, world bool) *FileObj {
	defer f.updateInfo()
	if owner {
		f.info.Permissions.Owner().SetWrite()
	} else {
		f.info.Permissions.Owner().ClearWrite()
	}
	if group {
		f.info.Permissions.Group().SetWrite()
	} else {
		f.info.Permissions.Group().ClearWrite()
	}
	if world {
		f.info.Permissions.World().SetWrite()
	} else {
		f.info.Permissions.World().ClearWrite()
	}
	log.Error(f.Perm(f.info.Permissions.FileMode()), "Making %s writable failed!", f.Path()) // aka +x
	return f
}

func (f *FileObj) PermWriteAll() *FileObj {
	return f.PermWrite(true, true, true)
}

func (f *FileObj) IsExecutable() bool      { return f.info.Permissions.World().HasExec() }
func (f *FileObj) IsExecutableGroup() bool { return f.info.Permissions.Group().HasExec() }
func (f *FileObj) IsExecutableOwner() bool { return f.info.Permissions.Owner().HasExec() }

func (f *FileObj) IsReadable() bool      { return f.info.Permissions.World().HasRead() }
func (f *FileObj) IsReadableGroup() bool { return f.info.Permissions.Group().HasRead() }
func (f *FileObj) IsReadableOwner() bool { return f.info.Permissions.Owner().HasRead() }

func (f *FileObj) IsWritable() bool      { return f.info.Permissions.World().HasWrite() }
func (f *FileObj) IsWritableGroup() bool { return f.info.Permissions.Group().HasWrite() }
func (f *FileObj) IsWritableOwner() bool { return f.info.Permissions.Owner().HasWrite() }
