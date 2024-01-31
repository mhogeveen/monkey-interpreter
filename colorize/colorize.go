package colorize

import "fmt"

type ColorCode string

const (
	RED   = "\033[0;31m"
	GREEN = "\033[0;32m"
	BLUE  = "\033[0;33m"
	NONE  = "\033[0m"
)

var colorCodes = map[string]ColorCode{
	"red":   RED,
	"green": GREEN,
	"blue":  BLUE,
	"none":  NONE,
}

func Colorize(s string, c string) string {
	return fmt.Sprintf("%v%v %v", colorCodes[c], s, NONE)
}
