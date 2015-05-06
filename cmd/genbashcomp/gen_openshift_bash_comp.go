package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/openshift/origin/pkg/cmd/admin"
	"github.com/openshift/origin/pkg/cmd/cli"
	"github.com/openshift/origin/pkg/cmd/openshift"
)

func OutDir(path string) (string, error) {
	outDir, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	stat, err := os.Stat(outDir)
	if err != nil {
		return "", err
	}

	if !stat.IsDir() {
		return "", fmt.Errorf("output directory %s is not a directory\n", outDir)
	}
	outDir = outDir + "/"
	return outDir, nil
}

func main() {
	// use os.Args instead of "flags" because "flags" will mess up the man pages!
	path := "rel-eng/completions/bash/"
	if len(os.Args) == 2 {
		path = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [output directory]\n", os.Args[0])
		os.Exit(1)
	}

	outDir, err := OutDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get output directory: %v\n", err)
		os.Exit(1)
	}
	outFile_openshift := outDir + "openshift"
	openshift := openshift.NewCommandOpenShift()
	openshift.GenBashCompletionFile(outFile_openshift)

	outFile_osc := outDir + "osc"
	osc := cli.NewCommandCLI("osc", "openshift cli")
	osc.GenBashCompletionFile(outFile_osc)

	outFile_osadm := outDir + "osadm"
	osadm := admin.NewCommandAdmin("osadm", "openshift admin", ioutil.Discard)
	osadm.GenBashCompletionFile(outFile_osadm)
}
