package face

import (
	"log"
	"time"
	"golang.org/x/net/context"
	pb "github.com/nEdAy/wx_attendance_api_server/internal/face/face_recognition"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50052"
)

func GetFaceCount(prefixCosUrl string, fileName string, faceToken string) (int32, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFaceRecognitionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	r, err := c.GetFaceCount(ctx, &pb.GetFaceCountRequest{PrefixCosUrl: prefixCosUrl, FileName: fileName, FaceToken: faceToken})
	if err != nil {
		log.Printf("could not Count: %v", err)
		return -1, err
	}
	log.Printf("Count: %d", r.Count)
	return r.Count, err
}

func IsMatchFace(prefixCosUrl string, fileName string, faceToken string) (bool, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFaceRecognitionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	r, err := c.IsMatchFace(ctx, &pb.IsMatchFaceRequest{PrefixCosUrl: prefixCosUrl, FileName: fileName, FaceToken: faceToken})
	if err != nil {
		log.Printf("could not IsMatchFace: %v", err)
		return false, err
	}
	log.Printf("IsMatchFace: %t", r.IsMatchFace)
	return r.IsMatchFace, err
}
