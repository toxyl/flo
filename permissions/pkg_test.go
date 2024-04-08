package permissions

import (
	"io/fs"
	"testing"
)

func TestPermissions_String(t *testing.T) {
	perms := []fs.FileMode{
		0000,
		0755,
		0750,
		0751,
		0752,
		0644,
		0777,
		0700,
		0400,
		0007,
	}

	t.Run("Permissions From FileModes", func(t *testing.T) {
		for _, perm := range perms {
			p := New("").Set(uint32(perm))
			t.Logf("(%6.2f%% %s) %s", p.Risk()*100, p.RiskString(), p.String())
		}
	})
}

func TestPermissions_From_File_String(t *testing.T) {
	files := []string{
		"/bin/chage",
		"/bin/chfn",
		"/bin/chsh",
		"/bin/fusermount3",
		"/bin/gpasswd",
		"/bin/mount",
		"/bin/newgrp",
		"/bin/ping",
		"/bin/pkexec",
		"/bin/plocate",
		"/bin/ssh-agent",
		"/dev/core",
		"/dev/initctl", // F
		"/dev/kvm",
		"/dev/log", // p or F
		"/dev/null",
		"/dev/sda",     // D
		"/dev/urandom", // Dc
		"/dev/zero",
		"/proc/thread-self/fd/0",
		"/proc/thread-self/fd/6",
		"/sys/class/net/eth0",
		"/sys/fs/bpf",
		"/tmp",
		"/usr/bin/crontab",
		"/usr/bin/expiry",
		"/usr/bin/locate",          // l
		"/etc/alternatives/locate", // l
		"/usr/bin/plocate",
		"/usr/bin/passwd",
		"/usr/local/bin/lha",          // l
		"/var/run/mysqld/mysqld.sock", // p
		"/var/run/user/1000/systemd/inaccessible/chr",  // c
		"/var/run/user/1000/systemd/inaccessible/dir",  // d
		"/var/run/user/1000/systemd/inaccessible/fifo", // p
		"/var/run/user/1000/systemd/inaccessible/reg",
		"/var/run/user/1000/systemd/inaccessible/sock", // S
		"/var/run/user/1000/systemd/units/invocation:dbus.service",
	}

	t.Run("Permissions From Files", func(t *testing.T) {
		for _, file := range files {
			p := New(file)
			t.Logf("%s | %s", p.String(), file)
		}
		t.Logf("\n\n")
	})
}
