package log

import "fmt"

type ErrorFunc func(err error, fmtStr string, args ...any) (shouldReturn bool)
type PanicFunc func(fmtStr string, args ...any)

var (
	Error = func(err error, fmtStr string, args ...any) (shouldReturn bool) {
		if err == nil {
			return false
		}
		fmt.Printf("ERROR: "+fmtStr+"\n       %s\n", append(args, err)...)
		return true
	}
	Panic = func(fmtStr string, args ...any) {
		panic(fmt.Sprintf(fmtStr+"\n", args...))
	}
)

func SetFns(fnError ErrorFunc, fnPanic PanicFunc) {
	if fnError == nil {
		Error = func(err error, fmtStr string, args ...any) (shouldReturn bool) { return err != nil }
	} else {
		Error = fnError
	}
	if fnPanic == nil {
		Panic = func(fmtStr string, args ...any) {}
	} else {
		Panic = fnPanic
	}
}
