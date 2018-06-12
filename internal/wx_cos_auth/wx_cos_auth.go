package wx_cos_auth

import (
	"log"
	"google.golang.org/grpc"
	"time"
	"golang.org/x/net/context"
)

const (
	address = "localhost:50051"
)

// AuthorizationTransport 生产鉴权签名
func AuthorizationTransport(method string, pathname string) (string, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewWXCosAuthClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	r, err := client.GetAuthData(ctx, &GetAuthDataRequest{Method: method, Pathname: pathname})
	if err != nil {
		log.Printf("could not AuthData: %v", err)
		return "", err
	}
	log.Printf("AuthData: %s", r.AuthData)
	return r.AuthData, err
}
