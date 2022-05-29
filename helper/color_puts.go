package helper

import (
	"github.com/gookit/color"
	"github.com/jedib0t/go-pretty/v6/text"
)

func SuccessPuts(tip, format string, v ...any) {
	color.Green.Printf(text.AlignRight.Apply(tip, 10) + "  ")
	color.Printf(format+"\n", v...)
}

func ErrorPuts(tip, format string, b ...any) {
	color.Red.Printf(text.AlignRight.Apply(tip, 10) + "  ")
	color.Printf(format+"\n", b...)
}
