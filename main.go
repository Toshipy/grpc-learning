package main

import (
	"fmt"
	// "io/ioutil"
	"log"
	"protobuf-lesson/pb"

	// "google.golang.org/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
)

func main() {
	employee := &pb.Employee{
		Id: 1,
		Name: "John Doe",
		Email: "john.doe@example.com",
		Occupation: pb.Occupation_DEVELOPER,
		PhoneNumbers: []string{"080-1234-5678", "090-1234-5678"},
		Projects: map[string]*pb.Company_Project{
			"project1": &pb.Company_Project{},
		},
		Profile: &pb.Employee_Text{
			Text: "My name is Toshi",
		},
		BirthDate: &pb.Date{
			Year:  1990,
			Month: 1,
			Day:   1,
		},
	}

	// binData, err := proto.Marshal(employee)
	// if err != nil {
	// 	log.Fatal("Can't serialize", err)
	// }

	// if err := ioutil.WriteFile("employee.bin", binData, 0666); err != nil {
	// 	log.Fatal("Can't write", err)
	// }

	// in, err := ioutil.ReadFile("employee.bin")
	// if err != nil {
	// 	log.Fatal("Can't read", err)
	// }

	// readEmployee := &pb.Employee{}
	// err = proto.Unmarshal(in, readEmployee)
	// if err != nil {
	// 	log.Fatalln("Can't unmarshal", err)
	// }

	// fmt.Println(readEmployee)

	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatal("Can't marshal to json", err)
	}

	// fmt.Println(out)

	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatal("Can't unmarshal from json", err)
	}

	fmt.Println(readEmployee)

}
