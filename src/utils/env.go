package utils

import (
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func GetEnvFile(env string) string {
	filename := []string{"../../config/properties/app.config.json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))
	return filePath
}

func GetEnvDBFile(env string) string {
	filename := []string{"../../config/properties/db/env.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))
	return filePath
}
