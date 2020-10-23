package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/rohit-go-utils/proto_address/tutorial"
)

func WritePerson(w io.Writer, p *pb.Person) {
	fmt.Fprintln(w, "Person ID:", p.Id)
	fmt.Fprintln(w, "  Name:", p.Name)
	if p.Email != "" {
		fmt.Fprintln(w, "  E-mail address:", p.Email)
	}

	for _, pn := range p.Phones {
		switch pn.Type {
		case pb.Person_MOBILE:
			fmt.Fprint(w, "  Mobile phone #: ")
		case pb.Person_HOME:
			fmt.Fprint(w, "  Home phone #: ")
		case pb.Person_WORK:
			fmt.Fprint(w, "  Work phone #: ")
		}
		fmt.Fprintln(w, pn.Number)
	}
}

func listPeople(w io.Writer, book *pb.AddressBook) {
	for _, p := range book.People {
		WritePerson(w, p)
	}
}

func ReadAddressbook(fileName string) {
	// Read the existing address book.
	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found.  Creating new file.\n", fileName)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}
	newBook := string(in)
	log.Printf("CURRENT BOOK CONTENTS\n: %s", newBook)
}
