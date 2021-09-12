package simple_example

import (
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func writeToJson(filename string, pb proto.Message) error {
	marshaler := &protojson.MarshalOptions{Indent: "\t"}

	out, err := marshaler.Marshal(pb)

	if err != nil {
		log.Fatalln("Error occure while marshing the protocol buffer to json string")
		return err
	}

	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("Error occur during write to json file.", err)
		return err
	}

	return nil
}

func readFromJson(filename string, pb proto.Message) error {
	in, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln("Error occur during reading from file.", err)
		return err
	}

	if err := protojson.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("Error occure while marshing the protocol buffer to json string")
		return err
	}

	return nil
}
