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

package client

import (
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"github.com/nori-io/nori-grpc/pkg/api/proto"

)

func NewClient(addr string, caFile string, serverHostOverride string) (proto.NoriClient, chan<- struct{}) {
	var opts []grpc.DialOption
	if len(caFile) != 0 && fileExists(caFile) {
		creds, err := credentials.NewClientTLSFromFile(caFile, serverHostOverride)
		if err != nil {
			fmt.Println("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
		fmt.Println("Created gRPC client with cert: %s", caFile)
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}

	closeCh := make(chan struct{})

	go func() {
		<-closeCh
		conn.Close()
	}()

	c := proto.NewNoriClient(conn)
	return c, closeCh
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
