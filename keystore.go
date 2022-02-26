package gokystr

import (
	"crypto/rsa"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/abecodes/goutls/rsakys"
)

const (
	keyExt = ".pem"
)

// Keystore holds all private keys
type Keystore struct {
	mu          sync.RWMutex
	base        string
	keyExt      string
	acceptEmpty bool
	write       bool
	keys        map[string]*rsa.PrivateKey
}

// New returns a new keystore instance
func New(dir string, options ...option) (*Keystore, error) {
	if !validatePath(dir) {
		return nil, ErrWrongDirFormat
	}
	k := &Keystore{
		base:        dir,
		write:       true,
		acceptEmpty: false,
	}

	// apply options
	for _, opt := range options {
		err := opt(k)
		if err != nil {
			return nil, err
		}
	}

	if k.keyExt == "" {
		k.keyExt = keyExt
	}

	// collect keys, if available
	s := map[string]*rsa.PrivateKey{}
	err := readFiles(dir, k.keyExt, func(p string) {
		key, err := rsakys.ReadPrivate(p)
		if err != nil {
			// not able to read the key for some reason
			return
		}

		kid := parseKID(p, k.keyExt)

		// since there can not exist two files with the same name, no checking here
		s[kid] = key
	})
	if err != nil {
		return nil, err
	}

	// no valid keys where found
	if len(s) == 0 && !k.acceptEmpty {
		return nil, ErrNoKeyFound
	}

	k.keys = s

	return k, nil
}

// PrivateKey returns the private key from the keystore
func (k *Keystore) PrivateKey(kid string) (*rsa.PrivateKey, error) {
	k.mu.Lock()
	defer k.mu.Unlock()

	key, ok := k.keys[kid]
	if !ok {
		return nil, ErrNoSuchKey
	}
	return key, nil
}

// PublicKey returns the public key from the keystore
func (k *Keystore) PublicKey(kid string) (*rsa.PublicKey, error) {
	k.mu.Lock()
	defer k.mu.Unlock()

	key, ok := k.keys[kid]
	if !ok {
		return nil, ErrNoSuchKey
	}
	return &key.PublicKey, nil
}

// Add appends a private key to the keystore
func (k *Keystore) Add(kid string, key *rsa.PrivateKey) error {
	k.mu.Lock()
	defer k.mu.Unlock()

	if _, ok := k.keys[kid]; ok {
		return ErrKeyExistsAlready
	}

	if k.write {
		err := rsakys.WritePKCS1PrivateKey(key, k.base+"/"+kid+k.keyExt)
		if err != nil {
			return err
		}
	}

	k.keys[kid] = key
	return nil
}

// Remove deletes a private key from the keystore
func (k *Keystore) Remove(kid string) error {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.write {
		err := os.Remove(k.base + "/" + kid + k.keyExt)
		if err != nil {
			return err
		}
	}

	delete(k.keys, kid)

	return nil
}

func parseKID(p string, ext string) string {
	return strings.TrimSuffix(path.Base(p), ext)
}
