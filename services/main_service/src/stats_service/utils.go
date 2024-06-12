package stats_service

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Client StatsServiceClient

func Connect(conn_string string) error {
	fmt.Println(conn_string)
	conn, err := grpc.Dial(conn_string, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	Client = NewStatsServiceClient(conn)
	return nil
}
