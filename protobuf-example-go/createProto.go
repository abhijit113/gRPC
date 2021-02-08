package main

import (
	"fmt"

	simplepb "github.com/abhijit113/gRPC/protobuf-example-go/src/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	fmt.Println("hi")
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)
	fmt.Println(sm.GetId())

	sm.Id = 65478
	fmt.Println(sm.GetId())
}
