package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	// p1 := person{First: "Jenny"}
	// p2 := person{First: "Penny"}
	// xp := []person{p1, p2}

	// bs, err := json.Marshal(xp)

	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("PRINT JSON", string(bs))

	// xp2 := []person{}
	// err = json.Unmarshal(bs, &xp2)

	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("Go Data Structure", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":3005", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{First: "Jenny"}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data:", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {

}