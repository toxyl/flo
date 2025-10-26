//go:build linux

package ownership

import (
	"fmt"
	"os"
	"os/user"
	"syscall"
)

func (fo *FileOwnership) Update() {
	if stat, err := os.Stat(fo.file); err == nil {
		if u, err := user.LookupId(fmt.Sprint(stat.Sys().(*syscall.Stat_t).Uid)); err == nil {
			fo.user = u.Username
			fo.uid = u.Uid
		}
		if g, err := user.LookupGroupId(fmt.Sprint(stat.Sys().(*syscall.Stat_t).Gid)); err == nil {
			fo.group = g.Name
			fo.gid = g.Gid
		}
	}
}
