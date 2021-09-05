package privatehostkey

import (
	"fmt"
	"io/ioutil"

	"github.com/Ex0dIa-dev/ssh-honeypot-go/helpers"
	gossh "golang.org/x/crypto/ssh"
)

//read the given hostkeyfile and return a gossh.Signer which contains the key
func ReadHostKeyFile(filepath string) (gossh.Signer, error) {

	if !helpers.FileExists(filepath) {
		return nil, fmt.Errorf("filepath: %v not exists", filepath)
	}

	keyBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	key, err := gossh.ParsePrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}
