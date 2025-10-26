package utils

import (
	"bytes"
)

type BytesIO struct {
	buf bytes.Buffer
}

func (s *BytesIO) Write(p []byte) (n int, err error) {
	return s.buf.Write(p)
}

func (s *BytesIO) Read(p []byte) (n int, err error) {
	return s.buf.Read(p)
}

func (s *BytesIO) Bytes() []byte {
	return s.buf.Bytes()
}

func (s *BytesIO) String() string {
	return s.buf.String()
}

func NewBytesIO(str []byte) *BytesIO {
	var buf BytesIO
	_, _ = buf.Write(str)
	return &buf
}
