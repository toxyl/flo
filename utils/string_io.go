package utils

import "strings"

type StringIO struct {
	strings.Builder
}

func (sw *StringIO) Write(p []byte) (n int, err error) {
	return sw.WriteString(string(p))
}

func (sw *StringIO) Read(p []byte) (n int, err error) {
	return strings.NewReader(sw.String()).Read(p)
}

func NewStringIO(str string) *StringIO {
	var buf StringIO
	_, _ = buf.Write([]byte(str))
	return &buf
}
