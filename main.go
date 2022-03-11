package main

import (
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/guseggert/go-ds-grpc/proto"
	grpcdssvr "github.com/guseggert/go-ds-grpc/server"

	radosds "github.com/coryschwartz/go-ds-rados"
)

func main() {
	var listen string
	var ok bool
	if listen, ok = os.LookupEnv("RADOSDS_LISTEN"); !ok {
		listen = ":8080"
	}
	cephconfig := os.Getenv("RADOSDS_CEPH_CONFIG")
	cephpool := os.Getenv("RADOSDS_CEPH_POOL")

	rds, err := radosds.NewDatastore(cephconfig, cephpool)
	if err != nil {
		panic(err)
	}
	gds := grpcdssvr.New(rds)
	s := grpc.NewServer()
	pb.RegisterDatastoreServer(s, gds)
	l, err := net.Listen("tcp", listen)
	if err != nil {
		panic(err)
	}
	s.Serve(l)
}
