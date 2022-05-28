package helper

import "github.com/gookit/color"

func SuccessPuts(tip, format string, v ...any) {
	color.Green.Printf("    [" + tip + "]    ")
	color.Printf(format+"\n", v...)
}

func ErrorPuts(tip, format string, b ...any) {
	color.Red.Printf("    [" + tip + "]    ")
	color.Printf(format+"\n", b...)
}
