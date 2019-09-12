package main

import (
	"encoding/json"
	"fmt"
)

// import "github.com/coreos/etcd"

type MyObject struct {
	Number int    `json:"number"`
	Word   string `json:"string"`
}

func main() {
	packt := "{packt: 15}"
	jsonPackt, ok := json.Marshal(packt)
	if ok != nil {
		panic("Could not marshal object")
	}
	fmt.Println(string(jsonPackt))

	// =========
	object := MyObject{5, "Packt"}
	oJson, _ := json.Marshal(object)
	fmt.Printf("%s\n", oJson)

	jsonBytes := []byte(`{"number":5, "string":"Packt"}`)
	var object2 MyObject
	err := json.Unmarshal(jsonBytes, &object2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number is %d, Word is %s\n", object2.Number, object2.Word)
}
