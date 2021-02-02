package erro

import (
	"errors"
	"fmt"
	"io"
	"runtime"
	"strings"
)

const delim = "\x00->"

// Wrap will wrap an error and return a new error that is annotated with the
// function/file/linenumber of where Wrap() was called
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	pc, filename, linenr, _ := runtime.Caller(1)
	strs := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	function := strs[len(strs)-1]
	return fmt.Errorf(delim+" Error in %s:%d (%s) %w", filename, linenr, function, err)
}

// Dump will dump the formatted error string (with each error in its own line)
// into w io.Writer
func Dump(w io.Writer, err error) {
	pc, filename, linenr, _ := runtime.Caller(1)
	strs := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	function := strs[len(strs)-1]
	err = fmt.Errorf("Error in %s:%d (%s) %w", filename, linenr, function, err)
	fmtedErr := strings.Replace(err.Error(), " "+delim+" ", "\n\n", -1)
	fmt.Fprintln(w, fmtedErr)
}

// Sdump will return the formatted error string (with each error in its own
// line)
func Sdump(err error) string {
	pc, filename, linenr, _ := runtime.Caller(1)
	strs := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	function := strs[len(strs)-1]
	err = fmt.Errorf("Error in %s:%d (%s) %w", filename, linenr, function, err)
	fmtedErr := strings.Replace(err.Error(), " "+delim+" ", "\n\n", -1)
	return fmtedErr
}

// S1dump will return the raw error string (the entire error stack trace in one
// line)
func S1dump(err error) string {
	pc, filename, linenr, _ := runtime.Caller(1)
	strs := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	function := strs[len(strs)-1]
	err = fmt.Errorf("Error in %s:%d (%s) %w", filename, linenr, function, err)
	return err.Error()
}

// Is reports whether any error in err's chain matches the target(s). Exactly
// the same as errors.Is, but variadic
func Is(err error, target error, targets ...error) bool {
	targets = append([]error{target}, targets...)
	for _, e := range targets {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}
