package main

import (
	"encoding/json"
	"fmt"
)

type Info struct {
	Name    string `json:"name"`
	City    string `json:"city,omitempty"`
	BigCity bool   `json:"-"`
}

func main() {
	e := `
	{
		"name":"abc",
		"city":"x",
		"BigCity":false
	}
	`
	var i Info
	//var i map[string]any
	err := json.Unmarshal([]byte(e), &i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	fmt.Println("bigcity",i.BigCity)
	// i.Name = "Raj"
	// data, err := json.Marshal(i)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = os.WriteFile("hello.json", data, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
