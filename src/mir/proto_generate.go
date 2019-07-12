package mir

//go:generate sh -c "pb-rs ./*.proto"
//go:generate sh -c "protoc -I. --gogo_out=plugins=grpc:. ./*.proto"
