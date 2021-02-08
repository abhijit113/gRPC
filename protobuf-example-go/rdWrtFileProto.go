package main

import (
	"fmt"
	"io/ioutil"
	"log"

	simplepb "github.com/abhijit113/gRPC/protobuf-example-go/src/simple"
	"github.com/golang/protobuf/proto"
)

func main() {

	sm := doSimple()
	readWriteDemo(sm)
}

func readWriteDemo(sm proto.Message) {

	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)
}

func writeToFile(flname string, pb proto.Message) error {

	out, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(flname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")

	return nil
}

func readFromFile(flname string, pb proto.Message) error {

	in, err := ioutil.ReadFile(flname)

	if err != nil {
		log.Fatalln("something went wrong while reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)

	if err2 != nil {
		log.Fatalln("Couldn't put bytes into protocol buffer struct", err2)
		return err2
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
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
	return &sm
}
