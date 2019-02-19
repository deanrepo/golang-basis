package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func main() {

	test1()

	// p1 := &person{
	// 	FirstName: "Eli",
	// 	LastName:  "Zhang",
	// 	Age:       32,
	// }

	// p2 := &person{
	// 	FirstName: "Dean",
	// 	LastName:  "Zhang",
	// 	Age:       30,
	// }

	// var xp = []*person{p1, p2}

	// data, err := json.Marshal(xp)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("json:", string(data))

	// j := `[{"firstName":"Eli","lastName":"Zhang","age":32},{"firstName":"Dean","lastName":"Zhang","age":30}]`

	// var xp = []*person{}
	// err := json.Unmarshal([]byte(j), &xp)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v, %+v\n", *xp[0], *xp[1])
}

func test1() {
	p1 := &person{
		FirstName: "Eli",
		LastName:  "Zhang",
		Age:       32,
	}

	p2 := &person{}

	data, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("p1 json: %s\n", string(data))

	err = json.Unmarshal(data, p2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("p2 : %+v\n", *p2)
}
