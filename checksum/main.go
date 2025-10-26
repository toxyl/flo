package checksum

import (
	"os"

	"github.com/toxyl/flo/codec"
)

type Checksum struct {
	algo    *codec.Codec
	file    string
	val     string
	changed bool
}

func (c *Checksum) sum() string {
	s := ""
	file, _ := os.Open(c.file)
	_ = c.algo.Decode(file, &s)
	return s
}

func (c *Checksum) Get() string {
	return c.val
}

func (c *Checksum) Update() {
	valid := false
	switch c.algo {
	case codec.SHA1, codec.SHA256, codec.SHA512, codec.MD5, codec.CRC32, codec.CRC64:
		valid = true
	}
	if !valid {
		panic("you must choose a valid checksum algorithm (SHA1, SHA256, SHA512, MD5, CRC32 or CRC64)")
	}
	s := c.sum()
	c.changed = c.val != s
	c.val = s
}

func (c *Checksum) Changed() bool {
	return c.changed
}

func (c *Checksum) Matches(other *Checksum) bool {
	return c.sum() == other.sum()
}

func (c *Checksum) MatchesString(str string) bool {
	return c.sum() == str
}

func (c *Checksum) MatchesBytes(data []byte) bool {
	return c.sum() == c.algo.EncodeString(data)
}

func New(algo *codec.Codec, path string) *Checksum {
	c := &Checksum{
		algo:    algo,
		file:    path,
		val:     "",
		changed: false,
	}
	return c
}
