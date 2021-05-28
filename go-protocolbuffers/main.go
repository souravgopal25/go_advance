package main

//https://developers.google.com/protocol-buffers/docs/reference/go-generated#package
import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Hello World")

	stundetObject := &Student{
		Name:  "Sourav Sharma",
		Email: "souravgopal25@gmail.com",
		Phone: "+917979970460",
		Address: &Address{
			State:   "Jharkhand",
			City:    "Jamshedpur",
			Country: "India",
		},
	}
	data, err := proto.Marshal(stundetObject)
	if err != nil {
		fmt.Println("Marshaling Error ", err)
	}
	fmt.Println(data)
	stundetObjectDecoded := &Student{}
	err = proto.Unmarshal(data, stundetObjectDecoded)
	if err != nil {
		fmt.Println("Unmarshalling error", err)
	}
	fmt.Println(stundetObjectDecoded)

}
