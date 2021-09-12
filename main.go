package main

import (
	"log"

	"github.com/manojpawar94/grpc/src/sample/samplepb"
)

func main() {
	log.Println("Declaring proto struct ")
	sm := doSimpleMessage()

	log.Println("---------------------------------")

	log.Println("Writing proto struct to file")
	if err := writeToFile("simple-message.bin", sm); err != nil {
		log.Panicln("Error occur while writing the proto file.", err)
	}

	log.Println("---------------------------------")

	log.Println("Reading proto struct from file")
	readSm := new(samplepb.SimpleMessage)
	if err := readFromFile("simple-message.bin", readSm); err != nil {
		log.Panicln("Error occur while reading the proto file.", err)
	}
	log.Printf("Read Simple Message\n %v.", readSm)
}
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
