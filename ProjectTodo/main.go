package main

import "fmt"

type detail struct {
	name string
	age  int
}

type todo struct {
	todo []map[int]detail
}

func main() {
	content := &todo{}

	fmt.Println("Enter the Name")


	m := map[int]detail{
		1: {"sanjay", 22},
		2: {"san", 222},
	}
	m1 := map[int]detail{
		1: {"selvin", 34},
		2: {"newww", 23},
	}

	var xm = []map[int]detail{m, m1}

	content.addTodo(xm...)
	content.deleteTodo(2, "newww", xm...)

}

func (t *todo) addTodo(mx ...map[int]detail) {
	t.todo = append(t.todo, mx...)
	fmt.Println(*t)
}
func (t *todo) deleteTodo(key int, value string, mx ...map[int]detail) {

	for i, vx := range mx {
		fmt.Println(i, vx)
		for _, v := range vx {
			if value == v.name {
				delete(vx, key)
			}
		}
	}
	fmt.Println(*t)
}
