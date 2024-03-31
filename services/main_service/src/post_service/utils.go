package post_service

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Client PostServiceClient

func Connect(conn_string string) error {
	conn, err := grpc.Dial(conn_string, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	Client = NewPostServiceClient(conn)
	return nil
}
