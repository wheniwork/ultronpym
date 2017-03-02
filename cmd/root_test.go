package cmd

import (
	"io/ioutil"
	"log"
	"os"
)

func setup() string {
	dir, err := ioutil.TempDir("", "pym")
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func cleanup(dir string) {
	os.RemoveAll(dir)
}
