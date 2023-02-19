package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dedihartono801/product-svc/pkg/config"
	"github.com/dedihartono801/product-svc/pkg/db"
	"github.com/dedihartono801/product-svc/pkg/services"
	pb "github.com/dedihartono801/protobuf/product/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	opts := []grpc.ServerOption{}
	tls := true

	if tls {
		certFile := "ssl/product-svc/server.crt"
		kefFile := "ssl/product-svc/server.pem"

		creds, err := credentials.NewServerTLSFromFile(certFile, kefFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
