package mir

//go:generate sh -c "pb-rs ./*.proto"
//go:generate sh -c "protoc -I. --go_out=. ./*.proto"
