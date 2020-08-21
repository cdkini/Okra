package okraerr

import (
	"fmt"
	"os"
)

// An OkraError is Okra's generalized error reporting struct. It is inclusive of syntax and runtime
// errors. Go error handling is used to determine the presence of an error and if applicable, an
// OkraError is reported to the user to provide details behind the issue in an Okra specific context.
type OkraError struct {
	line int
	col  int
	msg  string
}

func NewOkraError(line int, col int, msg string) OkraError {
	return OkraError{line, col, msg}
}

func (e OkraError) Error() string {
	return fmt.Sprintf("Error [%d:%d]: %s", e.line, e.col, e.msg)
}

// CheckErr is a wrapper around Go's traditional error checking, reporting an error code and
// an instance of OkraError if an error is found as the result of another function.
//   Args: code [int]       - The error code that best fits the situation (as described in okradocs.io/errors) // TODO: Update URL of docs
//         err  [error]     - The Go error produced by another function. Erroneous if not nil
//         oe   [OkraError] - Returned to the user with the location and circumstances around the error
//   Returns: nil
func CheckErr(err error, line int, col int, msg string) {
	if err != nil {
		ReportErr(line, col, msg)
	}
}

// ReportErr reports an error and exits the program with a corresponding error code.
//   Args: code [int]       - The error code that best fits the situation (as described in okradocs.io/errors) // TODO: Update URL of docs
//         oe   [OkraError] - Returned to the user with the location and circumstances around the error
//   Returns: nil
func ReportErr(line int, col int, msg string) {
	oe := NewOkraError(line, col, msg)
	fmt.Println(oe.Error())
	os.Exit(-1)
}
