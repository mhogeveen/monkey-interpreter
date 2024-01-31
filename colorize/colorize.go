package colorize

import "fmt"

type ColorCode string

const (
	RED    = "\033[0;31m"
	GREEN  = "\033[0;32m"
	YELLOW = "\033[0;33m"
	BLUE   = "\033[0;34m"
	PURPLE = "\033[0;35m"
	CYAN   = "\033[0;36m"
	GRAY   = "\033[0;37m"
	NONE   = "\033[0m"
)

var colorCodes = map[string]ColorCode{
	"red":    RED,
	"green":  GREEN,
	"yellow": YELLOW,
	"blue":   BLUE,
	"purple": PURPLE,
	"cyan":   CYAN,
	"gray":   GRAY,
	"none":   NONE,
}

func Colorize(s string, c string) string {
	return fmt.Sprintf("%v%v %v", colorCodes[c], s, NONE)
}
