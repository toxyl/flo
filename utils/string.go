package utils

import (
	"fmt"
	"strings"

	"github.com/toxyl/glog"
)

type String struct {
	builder strings.Builder
}

func (s *String) Rune(r rune) *String {
	s.builder.WriteRune(r)
	return s
}

func (s *String) CR() *String {
	return s.Rune('\r')
}

func (s *String) LF() *String {
	return s.Rune('\n')
}

func (s *String) CRLF() *String {
	return s.CR().LF()
}

func (s *String) Str(str string) *String {
	s.builder.WriteString(str)
	return s
}

func (s *String) StrUnderline(str string) *String {
	return s.Str(glog.Underline() + str + glog.Reset())
}

func (s *String) StrBold(str string) *String {
	return s.Str(glog.Bold() + str + glog.Reset())
}

func (s *String) StrItalic(str string) *String {
	return s.Str(glog.Italic() + str + glog.Reset())
}

func (s *String) StrClean(clean bool, str string) *String {
	if clean {
		return s.Str(glog.StripANSI(str))
	}
	return s.Str(str)
}

func (s *String) RuneAlt(cond bool, yes, no rune) *String {
	if cond {
		return s.Rune(yes)
	}
	return s.Rune(no)
}

func (s *String) StrAlt(cond bool, yes, no string) *String {
	if cond {
		return s.Str(yes)
	}
	return s.Str(no)
}

func (s *String) Pad(len int) *String {
	if len > 0 {
		return s.Str(strings.Repeat(" ", len))
	}
	return s
}

func (s *String) Rst() {
	s.builder.Reset()
}

func (s *String) String() string {
	return s.builder.String()
}

func (s *String) Print() {
	fmt.Print(s.String())
}

func NewString() *String {
	s := &String{
		builder: strings.Builder{},
	}
	return s
}
