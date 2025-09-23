package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	userv1 "github.com/DelightVLG/msc-auth/pkg/api/user/v1"
	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	port = ":8080"
)

type server struct {
	userv1.UnimplementedUserServiceServer
}

func (s *server) Create(ctx context.Context, req *userv1.CreateRequest) (*userv1.CreateResponse, error) {
	id := gofakeit.Int64()
	return &userv1.CreateResponse{
		Id: &id,
	}, nil
}

func (s *server) Get(ctx context.Context, id *userv1.GetRequest) (*userv1.GetResponse, error) {
	fmt.Println("Get", id)
	id2 := int64(1)
	name := gofakeit.Name()
	email := gofakeit.Email()
	role := userv1.Role(gofakeit.IntRange(1, 2))
	return &userv1.GetResponse{
		Id: &id2,
		Data: &userv1.UserData{
			Name:  &name,
			Email: &email,
			Role:  &role,
		},
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
	}, nil
}

func (s *server) Update(ctx context.Context, req *userv1.UpdateRequest) (*userv1.UpdateResponse, error) {
	fmt.Println("Update", req)
	return &userv1.UpdateResponse{}, nil
}

func (s *server) Delete(ctx context.Context, req *userv1.DeleteRequest) (*userv1.DeleteResponse, error) {
	fmt.Println("Delete", req)
	return &userv1.DeleteResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	userv1.RegisterUserServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
