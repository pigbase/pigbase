package main

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
)

type Account struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
	Salt string `json:"salt"`
}

type Value struct {
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

type Collection struct {
	Name   string  `json:"name_of_collection"`
	Values []Value `json:"values"`
}

type Data struct {
	Accounts    []Account    `json:"accounts"`
	Collections []Collection `json:"collections"`
}

func GetDataFilePath() string {
	path := ""
	switch runtime.GOOS {
	case "windows":
		path = "C:\\Program Files\\Pigbase\\data"
	case "linux":
		path = "/opt/pigbase/data"
	}
	return path
}

func GetDataFile() string {
	value, err := os.ReadFile(GetDataFilePath())

	if err != nil {
		log.Fatalf("Error while reading data file: %s", err)
	}

	val := string(value)

	return val
}

func DataFileToObject(value string) Data {
	var data Data
	json.Unmarshal([]byte(value), &data)

	return data
}

func (d Data) SaveDataConfig() {
	buffer, err := json.Marshal(d)
	if err != nil {
		log.Fatal("Error while creating string from object.")
	}
	err = os.WriteFile(GetDataFilePath(), buffer, 0666)
	if err != nil {
		log.Fatalf("Error while writing into data file. %s", err)
	}
	log.Printf("Saved data file.")
}

// f := DataFileToObject(GetDataFile())
// f.Collections[0].Values[0].Value = "Ceska republika"
// f.Collections[0].Values[0].Type = "String"
// f.SaveDataConfig()
