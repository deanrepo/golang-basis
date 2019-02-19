package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	anonymoutField()
}

func anonymoutField() {
	type Point struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	type Circle struct {
		Point  `json:"point"`
		Radius int `json:"radius"`
	}

	type Wheel struct {
		Circle `json:"circle"`
		Spokes int `json:"spokes"`
	}

	var wheel Wheel
	wheel.X = 2
	wheel.Y = 3
	wheel.Radius = 5
	wheel.Spokes = 25
	fmt.Printf("%+v\n", wheel)

	fmt.Println()

	data, err := json.Marshal(wheel)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("json string: %v\n", string(data))

}
