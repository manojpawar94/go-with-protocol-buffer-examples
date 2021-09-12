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

