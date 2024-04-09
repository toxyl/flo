package utils

import "strings"

type StringIO struct {
	strings.Builder
}

func (s *StringIO) Write(p []byte) (n int, err error) {
	return s.WriteString(string(p))
}

func (s *StringIO) Read(p []byte) (n int, err error) {
	return strings.NewReader(s.String()).Read(p)
}

func NewStringIO(str string) *StringIO {
	var buf StringIO
	_, _ = buf.Write([]byte(str))
	return &buf
}
