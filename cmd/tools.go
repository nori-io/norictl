package cmd

import (
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	packageManagerPath = "github.com/golang/dep/cmd/dep"
	packageManagerBin  = "dep"
)

type ToolChain struct {
	InstallDependencies bool

	gobin          string
	gopath         string
	packageManager string
}

func SetupToolChain() (*ToolChain, error) {
	tc := new(ToolChain)
	var err error

	tc.gobin, err = exec.LookPath("go")
	if err != nil {
		return nil, err
	}

	tc.gopath = os.Getenv("GOPATH")
	if tc.gopath == "" {
		tc.gopath = build.Default.GOPATH
	}

	tc.packageManager, err = exec.LookPath("dep")
	if err != nil {
		if err.(*exec.Error).Err == exec.ErrNotFound {
			// try install dep
			cmd := exec.Command(tc.gobin, "get", packageManagerPath)
			err = cmd.Run()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	tc.InstallDependencies = true

	return tc, nil
}

func (tc ToolChain) Do(path string) error {
	// get plugin
	err := exec.Command(tc.gobin, "get", "-d", path).Run()
	if err != nil {
		return err
	}

	// install dependencies
	if tc.InstallDependencies {
		cmd := exec.Command(packageManagerBin, "ensure")
		cmd.Dir = filepath.Join(tc.gopath, "src", path)
		err = cmd.Run()
		if err != nil {
			return err
		}
	}

	pluginName := filepath.Base(path)
	pluginName = filepath.Join("plugin", pluginName)

	// build plugin
	err = exec.Command(tc.gobin, "build", "-buildmode=plugin", "-o", pluginName).Run() // TODO remove debug info
	if err != nil {
		return err
	}

	return nil
}
