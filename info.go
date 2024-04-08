package flo

import (
	"io/fs"
	"os"
	"strings"
	"time"

	"github.com/toxyl/flo/checksum"
	"github.com/toxyl/flo/config"
	"github.com/toxyl/flo/ownership"
	"github.com/toxyl/flo/permissions"
	"github.com/toxyl/flo/utils"
	"github.com/toxyl/glog"
)

func (f *FileObj) updateInfo() {
	s, err := os.Stat(f.path)
	f.info.Ownership.Update()
	f.info.Name = f.Name()
	f.info.Mode = 0
	f.info.LastModified = time.Time{}
	f.info.Exists = !os.IsNotExist(err)
	f.info.Size = 0
	f.info.Checksum = checksum.New(config.ChecksumAlgorithm, f.path)
	f.info.Permissions = permissions.New(f.path)
	f.info.Path = f.Path()

	if f.info.Exists && s != nil {
		f.info.LastModified = s.ModTime()
		f.info.Mode = s.Mode()
		if f.info.Permissions.HasSize() {
			f.info.Size = s.Size()
		}
	}
}

type FileInfo struct {
	Name         string
	Exists       bool
	LastModified time.Time
	Mode         fs.FileMode
	Size         int64
	Path         string
	Checksum     *checksum.Checksum
	Permissions  *permissions.Permissions
	Ownership    *ownership.FileOwnership
}

func (f *FileInfo) NewerThan(t time.Time) bool { return f.LastModified.After(t) }
func (f *FileInfo) OlderThan(t time.Time) bool { return f.LastModified.Before(t) }

func (f *FileInfo) String(maxLenOwner, maxLenGroup int) string {
	res := utils.NewString().
		Str(f.Permissions.String()).Pad(1).
		Str(glog.PadLeft(glog.Auto(f.Ownership.User()), maxLenOwner, ' ')).Pad(1).
		Str(glog.PadRight(glog.Auto(f.Ownership.Group()), maxLenGroup, ' ')).Pad(1)

	if f.Permissions.HasSize() {
		res.Str(glog.PadLeft(glog.HumanReadableBytesIEC(f.Size), 12, ' '))
	} else {
		res.Str(strings.Repeat(" ", 11) + glog.WrapGray("-"))
	}
	res.Pad(1)

	ct := f.LastModified.Format("2006-01-02 15:04:05")
	if f.NewerThan(time.Now().Add(-24 * time.Hour)) {
		res.Str(glog.WrapYellow(ct))
	} else if f.NewerThan(time.Now().Add(-720 * time.Hour)) {
		res.Str(glog.WrapCyan(ct))
	} else if f.NewerThan(time.Now().Add(-8760 * time.Hour)) {
		res.Str(glog.WrapLightBlue(ct))
	} else {
		res.Str(glog.WrapGray(ct))
	}
	res.Pad(1)

	if f.Permissions.IsDir() {
		res.Str(config.IndicatorDir + " " + glog.Bold())
	} else if f.Permissions.IsCharDevice() {
		res.Str(config.IndicatorCharDevice + " " + glog.Italic())
	} else if f.Permissions.IsBlockDevice() {
		res.Str(config.IndicatorDevice + " " + glog.Italic())
	} else if f.Permissions.IsFIFO() {
		res.Str(config.IndicatorFIFO + " " + glog.Italic())
	} else if f.Permissions.IsSocket() {
		res.Str(config.IndicatorSocket + " " + glog.Italic())
	}

	if f.Permissions.IsSticky() {
		res.Str(config.IndicatorSticky + " ")
	}
	if f.Permissions.IsLink() {
		res.Str(glog.Italic())
	}

	if res.String() == "" {
		res.Str(config.IndicatorFile)
	}
	res.Str(f.Name).Str(glog.Reset())

	if f.Permissions.IsLink() {
		res.Pad(1).Str(config.IndicatorLink).Pad(1)
		if targetPath, err := os.Readlink(f.Path); err == nil {
			res.Str(glog.File(targetPath))
		} else {
			res.Str(glog.WrapRed("DEAD"))
		}
	}

	str := res.Pad(1).String()

	if !config.ColorMode {
		str = glog.StripANSI(str)
	}
	return str
}
