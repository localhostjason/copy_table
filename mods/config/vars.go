package config

import (
	"os"
	"path/filepath"
)

const (
	TimeFormat        = "2006-01-02 15:04:05"
	TimeFormatNoBlank = "20060102_150405"
	TimeFormatDay     = "20060102"
)

var RootDir string

func init() {
	RootDir = getExePath()
}

func getExePath() string {
	ex, _ := os.Executable()
	exeDir, _ := filepath.Abs(filepath.Dir(ex))
	return exeDir
}
