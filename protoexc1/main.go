package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	jerry := &Person{
		Name: "jerry",
		Age:  25,
		SFlowers: &SocialFlowers{
			Twitter: 1200,
			Youtube: 3500,
		},
	}

	data, err := proto.Marshal(jerry)
	if err != nil {
		log.Fatal("marshalling err:", err)
	}

	fmt.Println(data)

	newJerry := &Person{}
	err = proto.Unmarshal(data, newJerry)
	if err != nil {
		log.Fatal("unmarshalling err:", err)
	}
	fmt.Println(newJerry.GetName())
	fmt.Println(newJerry.GetAge())
	fmt.Println(newJerry.SFlowers.GetTwitter())
	fmt.Println(newJerry.SFlowers.GetYoutube())
}
