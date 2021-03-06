package simple_example

import (
	"log"
	"testing"

	"github.com/manojpawar94/grpc/src/simple_example/simplepb"
)

func TestWriteToFile(t *testing.T) {
	sm := NewSimpleMessage()

	log.Println("Writing proto struct to file")
	if err := writeToFile("simple-message.bin", sm); err != nil {
		log.Panicln("Error occur while writing the proto file.", err)
		t.Failed()
	}
}

func TestReadFromFile(t *testing.T) {
	readSm := new(simplepb.SimpleMessage)
	if err := readFromFile("simple-message.bin", readSm); err != nil {
		log.Panicln("Error occur while reading the proto file.", err)
		t.Failed()
	}
	log.Printf("Read Simple Message\n %v.", readSm)
}
