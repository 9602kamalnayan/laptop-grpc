package service

import (
	"context"
	"errors"
	"laptop-grpc/pb"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	Store LaptopStore
}

// mustEmbedUnimplementedLaptopServiceServer implements pb.LaptopServiceServer
func (s *LaptopServer) mustEmbedUnimplementedLaptopServiceServer() {
	panic("unimplemented")
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{Store: store}
}

func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with this id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop id is not valid uuid: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop id: %v", err)
		}
		laptop.Id = id.String()
	}

	err := server.Store.Save(laptop)

	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot store laptop to the store %v", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}

func (server *LaptopServer) SearchLaptop(req *pb.SearchLaptopRequest, stream pb.LaptopService_SearchLaptopServer) error {
	filter := req.GetFilter()
	log.Printf("receive a search-laptop request with a filter %v", filter)
	err := server.Store.Search(
		filter,
		func(laptop *pb.Laptop) error {
			res := &pb.SearchLaptopResponse{Laptop: laptop}
			err := stream.Send(res)
			if err != nil {
				return err
			}
			log.Printf("sent laptop with id %s", laptop.GetId())
			return nil
		},
	)
	if err != nil {
		return status.Errorf(codes.Internal, "Unexpected error: %v", err)

	}
	return nil
}
