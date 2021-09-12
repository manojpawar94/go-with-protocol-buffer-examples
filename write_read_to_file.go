package main

import (
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

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
