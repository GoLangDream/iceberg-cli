package helper

import "github.com/gookit/color"

func SuccessPuts(format string, a ...any) {
	color.Green.Printf(format+"\n", a...)
}

func ErrorPuts(format string, a ...any) {
	color.Red.Printf(format+"\n", a...)
}
