package flo

import (
	c "github.com/toxyl/flo/codec"
)

func (f *FileObj) checksum(codec *c.Codec) string {
	s := ""
	f.read(codec, &s)
	return s
}

func (f *FileObj) SHA1() string   { return f.checksum(c.SHA1) }
func (f *FileObj) SHA256() string { return f.checksum(c.SHA256) }
func (f *FileObj) SHA512() string { return f.checksum(c.SHA512) }
func (f *FileObj) MD5() string    { return f.checksum(c.MD5) }
func (f *FileObj) CRC32() string  { return f.checksum(c.CRC32) }
func (f *FileObj) CRC64() string  { return f.checksum(c.CRC64) }

func (f *FileObj) ReadSHA1(target *string) error   { return f.read(c.SHA1, target) }
func (f *FileObj) ReadSHA256(target *string) error { return f.read(c.SHA256, target) }
func (f *FileObj) ReadSHA512(target *string) error { return f.read(c.SHA512, target) }
func (f *FileObj) ReadMD5(target *string) error    { return f.read(c.MD5, target) }
func (f *FileObj) ReadCRC32(target *string) error  { return f.read(c.CRC32, target) }
func (f *FileObj) ReadCRC64(target *string) error  { return f.read(c.CRC64, target) }

func (f *FileObj) MustReadSHA1(target *string) *FileObj   { return f.mustRead(c.SHA1, target) }
func (f *FileObj) MustReadSHA256(target *string) *FileObj { return f.mustRead(c.SHA256, target) }
func (f *FileObj) MustReadSHA512(target *string) *FileObj { return f.mustRead(c.SHA512, target) }
func (f *FileObj) MustReadMD5(target *string) *FileObj    { return f.mustRead(c.MD5, target) }
func (f *FileObj) MustReadCRC32(target *string) *FileObj  { return f.mustRead(c.CRC32, target) }
func (f *FileObj) MustReadCRC64(target *string) *FileObj  { return f.mustRead(c.CRC64, target) }
