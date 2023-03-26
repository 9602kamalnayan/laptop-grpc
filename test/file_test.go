package test

import (
	"laptop-grpc/pb"
	"laptop-grpc/sample"
	"laptop-grpc/serializer"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../temp/laptop.bin"
	jsonfile := "../temp/laptop.json"
	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)
	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJsonFile(laptop1, jsonfile)
	require.NoError(t, err)

}
