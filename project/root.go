package project

import "os"

func Root() string {
	path, _ := os.Getwd()
	return path
}
