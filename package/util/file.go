package util

import "os"

func CreateDirNotExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		e := os.MkdirAll(path, os.ModePerm)
		if e != nil {
			return
		}
	}
}
