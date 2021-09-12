package simple_example

import (
	"log"
	"testing"

	"github.com/manojpawar94/grpc/src/simple_example/simplepb"
)

func TestWriteToJsonFile(t *testing.T) {
	sm := NewSimpleMessage()

	log.Println("Writing proto struct to file")
	if err := writeToJson("simple-message.json", sm); err != nil {
		log.Panicln("Error occur while writing the proto file.", err)
		t.Failed()
	}
}

func TestReadFromJsonFile(t *testing.T) {
	readSm := new(simplepb.SimpleMessage)
	if err := readFromJson("simple-message.json", readSm); err != nil {
		log.Panicln("Error occur while reading the proto file.", err)
		t.Failed()
	}
	log.Printf("Read Simple Message\n %v.", readSm)
}
