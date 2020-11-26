package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type persons struct {
	Name    string
	Surname string
	Age     int
}

func main() {
	jsonToString()
	stringToJSON()
}

func jsonToString() {
	p := persons{
		Name:    "Yasin",
		Surname: "Turnaoglu",
		Age:     32,
	}

	data, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println("Json string :", string(data))
}

func stringToJSON() {
	jsonData := []byte(`{"Name":"Yasin","Surname":"Turnaoglu","Age":32}`)

	var data persons

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Name : %s, Surname : %s, Age: %d", data.Name, data.Surname, data.Age)
}
