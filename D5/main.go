package main

import (
	"fmt"
	"time"
)

type Vertex struct {
	X int
	Y int
	Z string
}

func main() {
	// fmt.Println("Main function")
	// calculation.Do()

	//demoMaps()
	//demo()
	//demo3()
	demoConcurrency()

}

func demoMaps() {
	var countryCapitalMap map[string]string
	/* create a map*/
	countryCapitalMap = make(map[string]string)

	hm := map[string]string{ // assignment 3
		"name": "Sumit",
		"age":  "14",
	}

	fmt.Println(hm)

	s := "sumit"
	for i, c := range s {
		fmt.Println(i, c)
	}

	/* insert key-value pairs in the map*/
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* test if entry is present in the map or not*/
	capital, ok := countryCapitalMap["United States"]

	/* if ok is true, entry is present otherwise entry is absent*/
	if ok {
		fmt.Println("Capital of United States is", capital)
	} else {
		fmt.Println("Capital of United States is not present")
	}

}

func demo() {

	hm := map[string]int{
		"Goose":    13,
		"Maverick": 16,
		"Rooster":  15,
	}
	hm["Fiona"] = 3

	for key := range hm {
		fmt.Printf("Student name %v and age %d\n", key, hm[key])
	}

	value, ok := hm["Fiona"]

	if ok && value > 15 {
		fmt.Println(value, ok)
	} else {
		fmt.Println("Value less than 15")
	}

}

type Student struct {
	name string
	age  int
}

func demo3() {
	/* create a map*/
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}

	fmt.Println("Original map")

	/* print map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* delete an entry */
	delete(countryCapitalMap, "France")

	fmt.Println("Entry for France is deleted")

	fmt.Println("Updated map")

	/* print map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

}

func demoConcurrency() {
	go say("world")
	say("hello")

}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
