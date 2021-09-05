package helpers

import "os"

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func FileExists(filepath string) bool {

	stat, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !stat.IsDir()
}
