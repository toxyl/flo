//go:build windows

package ownership

import (
	"os/user"
)

func (fo *FileOwnership) Update() {
	// windows doesn't have the same user concept as linux,
	// let's be lazy and just assume the current user to be the owner
	u, _ := user.Current()
	fo.user = u.Name
	fo.uid = u.Uid
	fo.group = fo.uid
	fo.gid = u.Gid
}
