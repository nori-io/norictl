// Copyright © 2018 Secure2Work info@secure2work.com
//
// This program is free software: you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation, either version 3
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/testdata"

	"github.com/secure2work/nori/core/grpc"
)

const (
	addr    = "localhost:11111"
	testDir = "../testing_dir"
)

var (
	pem     = testdata.Path("server1.pem")
	key     = testdata.Path("server1.key")
	ca      = testdata.Path("ca.pem")
	testpem = filepath.Join(testDir, "server.pem")
	testkey = filepath.Join(testDir, "server.key")
)

func runServer(assert *assert.Assertions, useCerts bool) *grpc.Server {
	assert.Nil(os.RemoveAll(testDir))

	// FIXME
	server := grpc.NewServer([]string{testDir}, addr, true)
	if useCerts {
		server.SetCertificates(pem, key)
	} else {
		server.SetCertificates(testpem, testkey)
	}
	go func() {
		assert.Nil(server.Run())
	}()
	return server
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func TestPluginGet(t *testing.T) {
	assert := assert.New(t)

	server := runServer(assert, true)
	defer server.Stop()

	viper.Set("plugin-path", "github.com/secure2work/nori/plugins/cache")
	viper.Set("grpc-address", addr)
	viper.Set("ca", ca)
	viper.Set("ServerHostOverride", "x.test.youtube.com")
	viper.Set("dependencies", false)
	getCmd.Run(new(cobra.Command), []string{})

	assert.True(fileExists(filepath.Join(testDir, "cache.so")))
}

func TestCertsUpload(t *testing.T) {
	assert := assert.New(t)

	server := runServer(assert, false)
	defer server.Stop()

	assert.False(server.GetSecure())

	viper.Set("grpc-address", addr)
	viper.Set("ca", "")
	viper.Set("ServerHostOverride", "")
	viper.Set("pem", pem)
	viper.Set("key", key)

	// get passkey faster than creating a server
	time.Sleep(time.Second)
	viper.Set("passkey", server.GetPasskey())

	uploadCertsCmd.Run(new(cobra.Command), []string{})

	// restaring server
	time.Sleep(time.Second)

	assert.True(server.GetSecure())
	assert.True(fileExists(testpem))
	assert.True(fileExists(testkey))
}
