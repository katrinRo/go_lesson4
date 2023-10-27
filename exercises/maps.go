package main

import (
	"fmt"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
	exercise7()
	exercise8()
}

func exercise1() {
	// Простое упражнение 1
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(m1["one"])

}

func exercise2() {
	// Простое упражнение 2
	m2 := map[int]bool{}
	fmt.Println(m2[2])

}

func exercise3() {
	// Средней сложности упражнение 1
	m3 := map[string]*int{}
	val := 5
	m3["five"] = &val
	fmt.Println(*m3["five"])

}

func exercise4() {
	// Средней сложности упражнение 2
	m4 := map[int]string{1: "one", 2: "two", 3: "three"}
	delete(m4, 2)
	fmt.Println(len(m4))

}

func exercise5() {
	// Высокой сложности упражнение 1
	m5 := map[string]int{"a": 1, "b": 2, "c": 3}
	val, exists := m5["d"]
	fmt.Println(val, exists)

}

func exercise6() {
	// Высокой сложности упражнение 2
	m6 := make(map[int]int)
	m6[1] = m6[1] + 1
	m6[2] = m6[2] + 2
	fmt.Println(m6)
	// Ответ
}

func exerciseHigh1() {
	// Упражнение высокой сложности 1
	m1 := map[int]func() int{
		1: func() int { return 5 },
		2: func() int { return 10 },
	}
	fmt.Println(m1[1]() + m1[2]())

}

func exerciseHigh2() {
	// Упражнение высокой сложности 2
	type person struct {
		age int
	}

	m2 := map[person]string{
		{age: 20}: "Alice",
		{age: 25}: "Bob",
	}
	p := person{age: 20}
	fmt.Println(m2[p])

}

func exerciseMax1() {
	// Упражнение максимальной сложности 1
	m3 := map[string]*int{}
	val := 5
	m3["key"] = &val
	val = 10
	fmt.Println(*m3["key"])

}

func exerciseMax2() {
	// Упражнение максимальной сложности 2
	m4 := map[int]string{1: "a", 2: "b", 3: "c"}
	for k, v := range m4 {
		if k == 2 {
			m4[4] = "d"
		}
		if k == 3 {
			delete(m4, k)
		}
		fmt.Println(k, v)
	}
	fmt.Println(len(m4))

}
