package main

import (
	"fmt"
	"log"

	simplepb "github.com/abhijit113/gRPC/protobuf-example-go/src/simple"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {

	sm := doSimple()

	jsonDemo(sm)

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

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct:", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)

	if err != nil {
		log.Fatalln("can't convert to JSON", err)
		return ""
	}

	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}
}
