package cmd

import (
	"io/ioutil"
	"testing"

	"github.com/wheniwork/ultronpym/data"
)

func TestAddCommand(t *testing.T) {
	temp := setup()
	file := temp + "/manifest.yml"
	addCommand(file, []string{"missing-file"})

	dat, _ := ioutil.ReadFile(file)
	if string(dat) != "{}\n" {
		t.Fatal("Missing files should not be added")
	}

	testFile := temp + "/to-sign.sh"
	ioutil.WriteFile(testFile, []byte("sha"), 0644)

	addCommand(file, []string{testFile})
	manifest := data.LoadManifest(file)
	if manifest[testFile] != "d600474b1b8e50d3633c91c0cf1efc454b79c9624a43fd7de441ee71745726ab" {
		t.Fatal("File was not added properly")
	}

	cleanup(temp)
}
