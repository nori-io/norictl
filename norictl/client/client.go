package client

import (
	"log"

	"google.golang.org/grpc"

	"github.com/secure2work/nori-cli/proto"
)

func NewClient(addr string) (commands.CommandsClient, chan<- struct{}) {
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
