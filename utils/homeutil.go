package utils

import (
	"os"
	"runtime"
)

func HomeDir() string {
	if home := os.Getenv("RIVER_HOME"); home != "" {
		return home
	}
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}
	return os.Getenv("HOME")
}
