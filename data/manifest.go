package data

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/crypto/openpgp"

	yaml "gopkg.in/yaml.v2"
)

// Manifest is the structure of the manifest
type Manifest map[string]string

func LoadManifest(filename string) Manifest {
	manifest := Manifest{}
	if _, err := os.Stat(filename); err == nil {
		data, err := ioutil.ReadFile(filename)
		if err == nil {
			yaml.Unmarshal(data, &manifest)
		}
	}
	return manifest
}

func VerifySignature(manifest string, signature string, key string) (bool, error) {
	ringReader := strings.NewReader(key)
	signatureReader, err := os.Open(signature)
	if err != nil {
		return false, fmt.Errorf("Signature access failed: '%v'", err)
	}
	manifestReader, err := os.Open(manifest)
	if err != nil {
		return false, fmt.Errorf("Manifest access failed: '%v'", err)
	}

	keyring, err := openpgp.ReadArmoredKeyRing(ringReader)
	if err != nil {
		return false, fmt.Errorf("Keyring access failed: '%v'", err)
	}
	_, err = openpgp.CheckDetachedSignature(keyring, manifestReader, signatureReader)
	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("Check Detached Signature: " + err.Error())
	}

	return true, nil
}

func VerifySignatureGPG(manifest string, signature string) (bool, error) {
	out, err := exec.Command("gpg2", "--verify", signature, manifest).CombinedOutput()
	if err != nil {
		return false, errors.New(string(out))
	}
	return true, nil
}

func ComputeChecksum(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (m *Manifest) SaveManifest(filename string) error {
	data, err := yaml.Marshal(&m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}
