package main

import (
	"fmt"
)

// membuat fungsi
func swap(x, y string) (string, string) {
	return y, x
}

// membuat tipe data struct
type person struct {
	name string
	age  int
}

type student struct {
	grade int
	person
}

func main() {
	fmt.Printf("Hello World\n\n")

	//test defer
	defer fmt.Println("test defer")
	fmt.Println("coba")

	var decimalNumber = 2.68876

	fmt.Printf("Decimal Number %f\n", decimalNumber)
	fmt.Printf("Decimal Number %.4f\n\n", decimalNumber)

	var exist bool = true
	fmt.Printf("exist? %t\n\n", exist)

	// coba variabel tidak const
	nama := "Haikal"
	nama = "Haikal kali"
	fmt.Println("Nama saya : " + nama + "\n")

	// const tidak dapat diubah
	const umur int = 23
	// umur = 24
	fmt.Println("umur : ", umur)

	//penerapan array loop
	// var fruits = [0]string{"apple", "grape", "banana", "melon"}
	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	// pemanggilan function swap
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// penerapan struct
	var s1 = student{}
	s1.name = "wick"
	s1.age = 22
	s1.grade = 2

	fmt.Println("name : ", s1.name)
	fmt.Println("age : ", s1.age)
	fmt.Println("coba pake person : ", s1.person.age)
	fmt.Println("grade : ", s1.grade)

	// kombinasi struct dengan slice
	var allStudents = []person{
		{name: "Wick", age: 20},
		{name: "John", age: 21},
		{name: "Euy", age: 20},
	}
	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	// map
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40
	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"]) //mei belum dideklarasikan

	var fruits = []string{"apple", "grape", "banana", "melon"}

	var afruits = fruits[0:3]
	var bfruits = fruits[1:4]
	var aafruits = afruits[1:2]
	var bafruits = bfruits[0:1]

	fmt.Println(fruits)   //{apple grape banana melon}
	fmt.Println(afruits)  //{apple grape banana}
	fmt.Println(bfruits)  //{grape banana melon}
	fmt.Println(aafruits) // {grape}
	fmt.Println(bafruits) //{grape}

	// buah grape diubah menjadi pineapple
	bafruits[0] = "pineapple"

	fmt.Println(fruits)   //{apple grape banana melon}
	fmt.Println(afruits)  //{apple grape banana}
	fmt.Println(bfruits)  //{grape banana melon}
	fmt.Println(aafruits) // {grape}
	fmt.Println(bafruits) //{grape}
}
