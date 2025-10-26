package flo

import (
	c "github.com/toxyl/flo/codec"
	"github.com/toxyl/flo/config"
)

func (f *FileObj) compare(codec *c.Codec, file *FileObj) bool {
	return f.checksum(codec) == file.checksum(codec)
}

func (f *FileObj) SameAs(file *FileObj) bool       { return f.compare(config.ChecksumAlgorithm, file) }
func (f *FileObj) SameSHA1As(file *FileObj) bool   { return f.compare(c.SHA1, file) }
func (f *FileObj) SameSHA256As(file *FileObj) bool { return f.compare(c.SHA256, file) }
func (f *FileObj) SameSHA512As(file *FileObj) bool { return f.compare(c.SHA512, file) }
func (f *FileObj) SameMD5As(file *FileObj) bool    { return f.compare(c.MD5, file) }
func (f *FileObj) SameCRC32As(file *FileObj) bool  { return f.compare(c.CRC32, file) }
func (f *FileObj) SameCRC64As(file *FileObj) bool  { return f.compare(c.CRC64, file) }
