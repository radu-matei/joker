package main

import (
	"context"
	"io"
	"log"

	"github.com/radu-matei/kube-toolkit/example/scaffold/pkg/client"
	"github.com/spf13/cobra"
)

var (
	serverStreamUsage = "starts a server stream"
)

type serverStreamCmd struct {
	out    io.Writer
	client *client.Client
}

func newServerStreamCmd(out io.Writer) *cobra.Command {
	serverStreamCmd := &serverStreamCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "stream",
		Short: serverStreamUsage,
		Long:  serverStreamUsage,
		RunE: func(cmd *cobra.Command, args []string) error {
			setupConnection(remoteServerPort, localRandomPort)

			conn, err := client.GetGRPCConnection(serverHost)
			if err != nil {
				log.Fatalf("cannot create grpc connection: %v", err)
			}
			defer conn.Close()

			serverStreamCmd.client = ensureGRPCClient(serverStreamCmd.client, conn)
			return serverStreamCmd.run()
		},
	}

	return cmd
}

func (cmd *serverStreamCmd) run() error {

	err := cmd.client.ServerStream(context.Background())

	if err != nil {
		log.Fatalf("cannot initialize cloud: %v", err)
		return err
	}

	return nil
}
