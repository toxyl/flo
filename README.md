# FLO (File Layer Operations) 
... is a Go library to simplify working with files and directories. Feel free to use or fork it, if it works for you.

It is intended to be fast to use with a certain degree of flexibility. 

## Examples

### `examples/codecs/main.go`
A test application for reading and writing data using the different formats supported. Also tests checksum comparisons.

### `examples/exec/main.go`
A simple application that can execute another application.

### `examples/flat-tree/main.go`
A tool that can list all files and directories contained in a directory. Recursion depth can be limited.

### `examples/hello/main.go`
An example that shows how to create a script on the fly, make it executable and execute it.

### `examples/lha/main.go`
An example application that acts like `ls -lha` with a few extras.

### `examples/ls/main.go`
An example that acts like a call to `ls`. 

### `examples/tree/main.go`
Similar to the `flat-tree` example, but displays a tree structure instead, just like good ol' `tree`. 

## Logging
There are a few cases where the library logs, in those cases it will use the `Error` and `Panic` functions found in `log/main.go`. You can overwrite these functions to handle errors yourself, e.g. to ignore errors:
```golang
// silence all logging
log.SetFns(nil, nil) 

// no errors, custom panic handler
log.SetFns(nil, func(fmtStr string, args ...any) { panic(fmt.Sprintf("lib died: "+fmtStr, args...)) })

// no panics, custom error handler
log.SetFns(func(err error, fmtStr string, args ...any) (shouldReturn bool) { fmt.Println("Error", err.Error()) ; return err != nil }, nil)
```
