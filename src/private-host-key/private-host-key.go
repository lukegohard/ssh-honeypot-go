package privatehostkey

import (
	"io/ioutil"

	gossh "golang.org/x/crypto/ssh"
)

// ReadHostKeyFile read the given hostkeyfile and return a gossh.Signer which contains the key
func ReadHostKeyFile(filepath string) (gossh.Signer, error) {

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
