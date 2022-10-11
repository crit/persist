package errors

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	File    string `json:"file"`
	Line    int    `json:"line"`
}

func (err Error) Error() string {
	return fmt.Sprintf("%s:%d %s", err.File, err.Line, err.Message)
}

// New returns a formatted error.
func New(code int, format string, args ...any) error {
	_, file, line, ok := runtime.Caller(1)

	if !ok {
		file = "???"
		line = 0
	} else {
		file = srcFileParse(file)
	}

	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
		File:    file,
		Line:    line,
	}
}

// Code returns the code from an error.
func Code(err error) int {
	e, ok := err.(Error)

	if !ok {
		return http.StatusInternalServerError
	}

	return e.Code
}

// Message returns the message from an error.
func Message(err error) string {
	e, ok := err.(Error)

	if !ok {
		return err.Error()
	}

	return e.Message
}

// srcFileParse returns either the filename and extension, or the last directory
// (which is also usually the package name in Go) with the filename and extension.
// "project/src/model/user.go" => "model/user.go"
// "main.go" => "main.go
func srcFileParse(filename string) string {
	// "project/src/model/user.go" => "project/src/model", "user.go"
	dir, file := filepath.Split(filename)

	// "project/src/model" => ["project", "src", "model"]
	parts := strings.FieldsFunc(dir, func(c rune) bool {
		return c == filepath.Separator
	})

	if len(parts) > 0 {
		// => "model/user.go"
		return filepath.Join(parts[len(parts)-1], file)
	}

	// "user.go"
	return file
}
