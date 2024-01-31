package main

import (
	"encoding/json"
	"fmt"
)

type prod struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := prod{ID: 1, Name: "Notebook", Price: 2030.32, Tags: []string{"Barato", "Promoção"}}
	p1json, _ := json.Marshal(p1)
	fmt.Println(string(p1json))

	var p2 prod
	jsonString := `{"id": 1, "name": "pen", "price": 3.90, "tags": ["Promoção", "Baratissimo"]}`
	if err := json.Unmarshal([]byte(jsonString), &p2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)
}
