package templates

import (
	"embed"
	"path/filepath"
	"strings"
)

//go:embed project
var files embed.FS

func File(path string) string {
	fullPath := filepath.Join("project", path)
	content, err := files.ReadFile(fullPath)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func CreateDirs(fileCallback, dirCallback func(path string)) {
	createDirs("project", fileCallback, dirCallback)
}

func createDirs(path string, fileCallback, dirCallback func(path string)) {
	dirs, err := files.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		fullPath := filepath.Join(path, dir.Name())
		if dir.IsDir() {
			dirCallback(relDirPath(fullPath))
			createDirs(fullPath, fileCallback, dirCallback)
		} else {
			fileCallback(relFilePath(fullPath))
		}
	}
}

func relDirPath(path string) string {
	relPath, _ := filepath.Rel("project", path)
	return relPath
}

func relFilePath(path string) string {
	relPath, _ := filepath.Rel("project", path)
	return strings.TrimSuffix(relPath, ".tmpl")
}
