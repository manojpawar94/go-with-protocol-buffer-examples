package main

import (
	"log"
	"testing"

	"github.com/manojpawar94/grpc/src/sample/samplepb"
)

func TestWriteToJsonFile(t *testing.T) {
	sm := doSimpleMessage()

	log.Println("Writing proto struct to file")
	if err := writeToJson("simple-message.json", sm); err != nil {
		log.Panicln("Error occur while writing the proto file.", err)
		t.Failed()
	}
}

func TestReadFromJsonFile(t *testing.T) {
	readSm := new(samplepb.SimpleMessage)
	if err := readFromJson("simple-message.json", readSm); err != nil {
		log.Panicln("Error occur while reading the proto file.", err)
		t.Failed()
	}
	log.Printf("Read Simple Message\n %v.", readSm)
}
