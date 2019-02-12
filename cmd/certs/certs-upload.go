// Copyright © 2018 Nori info@nori.io
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

package certs_cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/nori-io/nori/core/grpc"
	"github.com/nori-io/nori/proto"
	"github.com/nori-io/norictl/client"
)

var uploadCertsCmd = &cobra.Command{
	Use:   "upload",
	Short: "uploading certs",
	Run: func(cmd *cobra.Command, args []string) {
		pem := viper.GetString("pem")
		key := viper.GetString("key")
		passkey := viper.GetString("passkey")

		if len(pem) == 0 || len(key) == 0 {
			logrus.Fatal("Required pem and key files")
		}

		if len(passkey) == 0 {
			logrus.Fatal("Required passkey")
		}

		pemFile, err := os.Open(pem)
		if err != nil {
			logrus.Fatal(err)
		}
		defer pemFile.Close()

		keyFile, err := os.Open(key)
		if err != nil {
			logrus.Fatal(err)
		}
		defer keyFile.Close()

		pemBytes, err := ioutil.ReadAll(pemFile)
		if err != nil {
			logrus.Fatal(err)
		}

		keyBytes, err := ioutil.ReadAll(keyFile)
		if err != nil {
			logrus.Fatal(err)
		}

		client, closeCh := client.NewClient(
			viper.GetString("grpc-address"),
			viper.GetString("ca"),
			viper.GetString("ServerHostOverride"),
		)

		// encrypting...
		pk, err := grpc.PasskeyFromString(passkey)
		if err != nil {
			logrus.Fatal(err)
		}

		var hmac []byte

		pemBytes, hmac, err = pk.Encrypt(pemBytes)
		if err != nil {
			logrus.Fatal(err)
		}

		hmacLen := len(hmac)

		pemBs := append([]byte{byte(hmacLen)}, hmac...)
		pemBs = append(pemBs, pemBytes...)

		keyBytes, hmac, err = pk.Encrypt(keyBytes)
		if err != nil {
			logrus.Fatal(err)
		}

		keyBs := append([]byte{byte(hmacLen)}, hmac...)
		keyBs = append(keyBs, keyBytes...)

		reply, err := client.CertsUploadCommand(context.Background(), &commands.CertsUploadRequest{
			Pem: pemBs,
			Key: keyBs,
		})

		close(closeCh)

		if err != nil {
			logrus.Fatal(err)
			if reply != nil {
				logrus.Fatal(reply.Error)
			}
		} else {
			fmt.Printf("Certificates %q and %q successfully uploaded\n", pem, key)
		}
	},
}

func init() {
	CertsCmd.AddCommand(uploadCertsCmd)
}
