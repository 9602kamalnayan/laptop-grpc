package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func protobufToJson(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		Indent:          " ",
		UseProtoNames:   false,
		EmitUnpopulated: true,
		UseEnumNumbers:  true,
	}
	b, err := marshaler.Marshal(message)
	return string(b), err

}
