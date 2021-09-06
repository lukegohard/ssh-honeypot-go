package helpers

import "os"

// CheckErr checks errors
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

// FileExists return true if "filepath" exists, else false
func FileExists(filepath string) bool {

	stat, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !stat.IsDir()
}

func DirExists(dirpath string) bool {

	_, err := os.Stat(dirpath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
