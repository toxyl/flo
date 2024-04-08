package config

import (
	"github.com/toxyl/flo/codec"
	"github.com/toxyl/glog"
)

var (
	ChecksumAlgorithm = codec.SHA256
	ColorMode         = true
)
var (
	ModeNone   = glog.WrapGray("-")
	ModeSticky = glog.WrapBrightYellow("s")
	ModeUID    = glog.WrapGreen("u")
	ModeGID    = glog.WrapOrange("g")
	ModeLink   = glog.WrapPurple("L")
	ModeNoLink = glog.WrapPurple(" ")
)
var (
	PermNone  = glog.WrapGray("-")
	PermRead  = glog.WrapDarkGreen("r")
	PermWrite = glog.WrapDarkYellow("w")
	PermExec  = glog.WrapDarkRed("x")
)
var (
	RiskNone   = glog.WrapGreen(" ")
	RiskLow    = glog.WrapYellow(" ")
	RiskMedium = glog.WrapOrange("▪")
	RiskHigh   = glog.WrapRed("▪")
)
var (
	TypeUnknown     = glog.WrapGray("?")
	TypeFile        = glog.WrapDarkGray("-")
	TypeDir         = glog.WrapLightBlue("d")
	TypeFIFO        = glog.WrapPink("p")
	TypeSocket      = glog.WrapLime("s")
	TypeBlockDevice = glog.WrapBlue("b")
	TypeCharDevice  = glog.WrapBrightYellow("c")
)
var (
	IndicatorDir        = glog.WrapLightBlue("▊")
	IndicatorCharDevice = glog.WrapBrightYellow("▍")
	IndicatorDevice     = glog.WrapBlue("▍")
	IndicatorFIFO       = glog.WrapPink("▍")
	IndicatorSocket     = glog.WrapLime("▍")
	IndicatorSticky     = glog.WrapDarkRed("*")
	IndicatorFile       = glog.WrapGray(" ")
	IndicatorLink       = glog.WrapPurple("┅⮞")
	IndicatorNoLink     = glog.WrapPurple("  ")
)
