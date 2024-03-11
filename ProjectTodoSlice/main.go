package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"math/rand"
	"os"
)

type detail struct {
	First string
	Age   int
}

type lists struct {
	Entries []map[int]detail
}

var name string
var age int

const fileName = "sanjay"

func main() {
	lists := loadListsFromFile()

	fmt.Println("Enter the Name")
	fmt.Scan(&name)

	fmt.Println("Enter the Age")
	fmt.Scan(&age)

	lists.addTodo(name, age)

	saveListsToFile(lists)
	fmt.Println("Your list-------------")
	getList(lists)

	// fmt.Println(*lists)
}

func (l *lists) addTodo(name string, age int) {
	randNum := rand.Intn(1000)
	newEntry := map[int]detail{
		randNum: {name, age},
	}
	l.Entries = append(l.Entries, newEntry)
	fmt.Println(*l)
}

func loadListsFromFile() lists {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return lists{}
	}

	var loadedLists lists
	err = json.Unmarshal(data, &loadedLists)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return lists{}
	}

	return loadedLists
}

func saveListsToFile(l lists) {
	fmt.Println(l)
	data, err := json.MarshalIndent(l, "", "  ")

	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	err = os.WriteFile(fileName, data, 0666)
	if err != nil {
		// fmt.Println("Error writing to file:", err)
	}
}
func getList(l lists) {
	for _, mv := range l.Entries {
		// fmt.Println(mv)

		for k, sv := range mv {
			fmt.Println(k, sv.First, sv.Age)

		}

	}

}
