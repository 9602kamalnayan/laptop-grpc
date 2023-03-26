package main

import (
	"context"
	"flag"
	"io"
	"laptop-grpc/pb"
	"laptop-grpc/sample"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateLaptop(laptopClient pb.LaptopServiceClient) {
	laptop := sample.NewLaptop()
	laptop.Id = ""
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Fatal("laptop already exists")
		} else {
			log.Fatal("cannot create laptop")
		}
		return

	}
	log.Printf("created laptop with id: %s", res.Id)
}
func searchLaptop(laptopClient pb.LaptopServiceClient, filter *pb.Filter) {
	log.Printf("search filter: %s", filter)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	req := &pb.SearchLaptopRequest{Filter: filter}
	stream, err := laptopClient.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatal("cannot search laptop: ", err)

	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}
		laptop := res.GetLaptop()
		log.Printf("-found: %s", laptop.GetId())
		log.Printf(" + brand: %s", laptop.GetBrand())
		log.Printf(" + name: %s", laptop.GetName())
		log.Printf(" + cpu cores: %d", laptop.GetCpu().GetNumberCores())
		log.Printf(" + cpu min ghz: %v", laptop.GetCpu().GetMinGhz())
		log.Printf(" + ram: %v,%v", laptop.GetRam().GetValue(), laptop.GetRam().GetUnit())
		log.Printf(" + price:%fusd", laptop.GetPriceUsd())
	}
}
func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server:", err)
	}
	laptopClient := pb.NewLaptopServiceClient(conn)
	for i := 0; i < 10; i++ {
		CreateLaptop(laptopClient)
	}
	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	}
	searchLaptop(laptopClient, filter)

}
