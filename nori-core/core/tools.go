// Copyright © 2018 Secure2Work info@secure2work.com
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package core

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

	// FIXME add filepath.SplitList()
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
