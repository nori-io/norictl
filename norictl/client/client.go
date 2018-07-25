package client

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/secure2work/nori-cli/proto"
)

func NewClient(addr string, caFile string, serverHostOverride string) (commands.CommandsClient, chan<- struct{}) {
	var opts []grpc.DialOption
	if len(caFile) != 0 && fileExists(caFile) {
		creds, err := credentials.NewClientTLSFromFile(caFile, serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	closeCh := make(chan struct{})

	go func() {
		<-closeCh
		conn.Close()
	}()

	c := commands.NewCommandsClient(conn)
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
