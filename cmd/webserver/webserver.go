package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/goldencoderam/tr-evolvers-backend/pkg/database"
	"github.com/goldencoderam/tr-evolvers-backend/pkg/grpc/service"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "Server port")

type energyMeterServer struct {
	service.UnimplementedEnergyMeterServiceServer

	mu sync.Mutex
}

func newServer() *energyMeterServer {
	s := &energyMeterServer{}
	return s
}

func (e *energyMeterServer) CreateEnergyMeter(
	ctx context.Context,
	createEnergyMeterDto *service.CreateEnergyMeterDto,
) (*service.EnergyMeterDto, error) {
	result, error := database.CreateEnergyMeter(createEnergyMeterDto)

	return &service.EnergyMeterDto{
		Serial:  result.Serial,
		BrandId: uint32(result.EnergyMeterBrandID),
	}, error
}

func (e *energyMeterServer) CreateEnergyMeterBrand(
	ctx context.Context,
	CreateEnergyMeterBrandDto *service.CreateEnergyMeterBrandDto,
) (*service.EnergyMeterBrandDto, error) {
	result, error := database.CreateEnergyMeterBrand(CreateEnergyMeterBrandDto)

	return &service.EnergyMeterBrandDto{
		Id:   uint32(result.ID),
		Name: result.Name,
	}, error
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	service.RegisterEnergyMeterServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
