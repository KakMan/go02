package main

import (
	"fmt"

	"gitlab.com/jirapol/hello/name"
)

type human struct {
	name    string
	surname string
	age     int
	isAdult bool
}

func (h human) printInfo() {
	fmt.Println("Name : ", h.name, " ", h.surname, "  Age:", h.age, " Adult:", h.isAdult)
}

func setAdult(h *human) {
	h.isAdult = h.age > 18
}

func getHuman(iname string, isurname string, iage int) human {

	return human{name: iname, surname: isurname, age: iage}

}

func main2() {
	var even_msg string
	var odd_msg string

	even_msg = "Even"
	odd_msg = "odd"

	fmt.Println(name.Name)
	/*
		even_msg := "Even"
		odd_msg := "Odd"
	*/

	/*var names [2]string

	names[0] = "name1"
	names[1] = "name2"
	*/
	// หรือ ประกาศแบบ นี้ก็ได้
	// names := [3]string{"name1", "name2", "name3"}
	// names := []string{"name1","name2","name3"}

	//
	var names []string
	names = append(names, "name1")
	names = append(names, "name2")
	names = append(names, "name3")
	names = append(names, "name4")

	for index, name := range names {
		fmt.Println(index, name)
	}

	fmt.Println("Hello world.")

	for i := 0; i <= 5; i++ {
		if i > 0 {
			fmt.Print(",")
		}
		if i%2 == 0 {
			fmt.Print(i, "-", even_msg)
		} else {
			fmt.Print(i, "-", odd_msg)
		}
	}
	fmt.Println(" ")
	var age int
	var h human
	age = 38
	h = getHuman("Jirapol", "Munthawonwong", age)
	setAdult(&h)

	h.printInfo()
}
