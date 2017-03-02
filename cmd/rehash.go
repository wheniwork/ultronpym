package cmd

import "github.com/wheniwork/ultronpym/data"

func rehashManifest(manifestFile string) {
	manifest := data.LoadManifest(manifestFile)
	for file, sha := range manifest {
		newSha, err := data.ComputeChecksum(file)
		if err != nil {
			fatal(err)
		}
		printHash(file, newSha, newSha != sha)
		manifest[file] = newSha
	}
	manifest.SaveManifest(manifestFile)
	success("🍠  Files Hashed!")
}
