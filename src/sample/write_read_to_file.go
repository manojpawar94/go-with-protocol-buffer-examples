package sample

import (
	"io/ioutil"
	"log"

	"github.com/manojpawar94/grpc/src/sample/samplepb"
	"google.golang.org/protobuf/proto"
)

func doSimpleMessage() *samplepb.SimpleMessage {
	sm := &samplepb.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "This is simple message",
		SampleList: []int32{1, 4, 7},
	}

	log.Printf("Simple message is %v\n", sm)

	// renaming the Name property value of struct
	sm.Name = "This is simple renamed messaged"

	log.Printf("Message Id: %d\n", sm.GetId())

	return sm
}

func writeToFile(filename string, pb proto.Message) error {
	out, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Error occur during marshing the proto message.", err)
		return err
	}

	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("Error occur during write to file.", err)
		return err
	}

	return nil
}

func readFromFile(filename string, pb proto.Message) error {
	in, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln("Error occur during reading from file.", err)
		return err
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Error occur during unmarshing the proto message from file.", err)
		return err
	}

	return nil
}
