package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Hero is a struct definition of a hero
type Hero struct {
	Name string
	Role string
}

// Heroes is a slice of Hero
type Heroes []Hero

func main() {
	var h Heroes
	data, err := ioutil.ReadFile("heroes.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &h)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", h)
}
