package helper

import (
	"os/exec"
	"strings"
)

func RunCmd(name string, arg ...string) {
	SuccessPuts("run", "%s %s", name, strings.Join(arg, " "))
	cmd := exec.Command(name, arg...)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
