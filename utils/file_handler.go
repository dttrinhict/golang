package utils

import "os"

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}


func CheckEntryExsisted(entries []os.DirEntry, entry os.DirEntry) bool {
	for _, v := range entries {
		if entry == v {
			return true
		}
	}
	return false
}