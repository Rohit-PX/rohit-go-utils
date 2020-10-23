package main

import (
	"bytes"
	//"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	pb "github.com/rohit-go-utils/proto_address/tutorial"
	"github.com/rohit-go-utils/proto_address/utils"
)

func main() {
	//fmt.Printf("RK testing write")
	buf := new(bytes.Buffer)
	// [START populate_proto]
	p1 := pb.Person{
		Id:    8888,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	p2 := pb.Person{
		Id:    3333,
		Name:  "Hello Doe",
		Email: "hello@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	// [END populate_proto]
	utils.WritePerson(buf, &p1)
	utils.WritePerson(buf, &p2)
	fname := os.Args[1]
	book := &pb.AddressBook{People: []*pb.Person{&p1, &p2}}
	// ...

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	utils.ReadAddressbook(fname)

}
