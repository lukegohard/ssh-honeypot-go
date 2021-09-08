package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

// Config contains the json "struct" of config.json file
type Config struct {
	Dev  string `json:"creator"`
	Auth Auth   `json:"auth"`
}

// Auth contains authentication json struct data of config.json file
type Auth struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// ParseConfigFile returns a Config struct, which contains config data
func ParseConfigFile() Config {

	var config Config

	configBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/config.json", GetRootPath()))
	CheckErr(err)
	err = json.Unmarshal(configBytes, &config)
	CheckErr(err)

	return config
}

// GetRootPath returns the string of the root directory of the project
func GetRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../../")
}

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

// DirExists return true if "dirpath" exists, else false
func DirExists(dirpath string) bool {

	_, err := os.Stat(dirpath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
