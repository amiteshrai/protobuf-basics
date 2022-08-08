package main

import (
	"fmt"

	pb "github.com/amiteshrai/protobuf-basics/proto"
)

func simple() *pb.Simple {
	return &pb.Simple{
		Id:       1,
		IsSimple: true,
		Name:     "Simple",
		Sample:   []int32{1, 2, 3, 4, 5, 6},
	}
}

func getEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_BLUE,
	}
}

func main() {
	data := simple()
	fmt.Println(data)

	enumData := getEnum()
	fmt.Println(enumData)
}
