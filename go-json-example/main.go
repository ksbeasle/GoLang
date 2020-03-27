package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Objects struct {
	Objects []object `json:"Objects"`
}

type object struct {
	Topic string `json:"topic"`
	Inner inner  `json:"inner"`
}

type inner struct {
	Inside1 int `json:"inside1"`
	Inside2 int `json:"inside2"`
	Inside3 int `json:"inside3"`
}

func main() {
	f, err := os.Open("test.json")
	if err != nil {
		panic(err.Error())
	}
	log.Println("JSON file successfully opened.")

	jsonToByteArr, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}
	log.Println("JSON file successfully read.")
	var objects Objects
	json.Unmarshal(jsonToByteArr, &objects)

	for i := 0; i < len(objects.Objects); i++ {
		fmt.Println(objects.Objects[i].Topic)
		fmt.Println(objects.Objects[i].Inner.Inside1)
		fmt.Println(objects.Objects[i].Inner.Inside2)
		fmt.Println(objects.Objects[i].Inner.Inside3)
	}

	//read in json using interface{}
	var aMap map[string]interface{}
	json.Unmarshal(jsonToByteArr, &aMap)
	//fmt.Println(aMap)

	for key, value := range aMap {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
		//TODO figure out how to get the inner map, another interface maybe?

	}
	defer f.Close()
}
