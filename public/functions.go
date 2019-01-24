package public

import (
	"github.com/fatih/color"
	"runtime"
)

func CheckErr(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		color.Red("Error got at %s, line %d with message: %s", file, line, err.Error())
	}
}
