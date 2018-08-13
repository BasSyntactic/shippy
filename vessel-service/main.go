package main

import (
	pb "github.com/shippy/vessel-service/proto/vessel"
)

func main() {
	
}

type Repository interface {
	FindAvailable(*pb.Specification)
}
