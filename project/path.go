package project

import "path/filepath"

var paths = map[string]*fileInfo{
	"migration": {
		path:  "db/migration",
		isDir: true,
	},
	"databaseConfig": {
		path:  "config/database.yml",
		isDir: false,
	},
}

type fileInfo struct {
	path  string
	isDir bool
	tmpl  string
}

func Path(relativePath string) string {
	return filepath.Join(Root(), relativePath)
}

func MigrationPath() string {
	return paths["migration"].path
}

func DatabaseConfigFilePath() string {
	return paths["databaseConfig"].path
}
