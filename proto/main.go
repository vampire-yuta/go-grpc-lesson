package main

import (
	"log"
	"proto-lesson/pb"

	"github.com/gogo/protobuf/jsonpb"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Sam",
		Email:       "test@aaa.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"000-0000-0000", "000-0000-0000"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Sam.",
		},
	}

	// binData, err := proto.Marshal(employee)
	// if err != nil {
	// 	log.Fatal("marshaling error: ", err)
	// }

	// if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
	// 	log.Fatal("file write error: ", err)
	// }

	// in, err := ioutil.ReadFile("test.bin")
	// if err != nil {
	// 	log.Fatal("file read error: ", err)
	// }

	// readEmployee := &pb.Employee{}

	// err = proto.Unmarshal(in, readEmployee)
	// if err != nil {
	// 	log.Fatal("unmarshaling error: ", err)
	// }

	// fmt.Println(readEmployee)

	m := jsonpb.Marshaler{}

	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// fmt.Println(out)

	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
}
