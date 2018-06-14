package face_recognition

import (
	"time"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"github.com/rs/zerolog/log"
)

const (
	address = "localhost:50052"
)

func GetFaceCount(prefixCosUrl string, fileName string, faceToken string) (int32, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msgf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewFaceRecognitionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	r, err := c.GetFaceCount(ctx, &GetFaceCountRequest{PrefixCosUrl: prefixCosUrl, FileName: fileName, FaceToken: faceToken})
	if err != nil {
		log.Error().Msgf("could not Count: %v", err)
		return -1, err
	}
	log.Info().Msgf("Count: %d", r.Count)
	return r.Count, err
}

func IsMatchFace(prefixCosUrl string, fileName string, faceToken string) (bool, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msgf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewFaceRecognitionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	r, err := c.IsMatchFace(ctx, &IsMatchFaceRequest{PrefixCosUrl: prefixCosUrl, FileName: fileName, FaceToken: faceToken})
	if err != nil {
		log.Error().Msgf("could not IsMatchFace: %v", err)
		return false, err
	}
	log.Info().Msgf("IsMatchFace: %t", r.IsMatchFace)
	return r.IsMatchFace, err
}
