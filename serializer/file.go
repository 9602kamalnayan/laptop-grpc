package serializer

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

func WriteProtobufToJsonFile(message proto.Message, filename string) error {
	data, err := protobufToJson(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to json: %w", err)
	}
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("can not write json data to file %w", err)
	}
	return nil

}
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary : %w", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)

	if err != nil {
		return fmt.Errorf("cannot write data to binary file: %w", err)
	}
	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read from binary file: %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshall binary to proto message: %w", err)
	}
	return nil
}
