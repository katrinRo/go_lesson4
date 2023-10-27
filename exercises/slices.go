package main

import "fmt"

func exercise1() {
	s := []int{1, 2, 3, 4, 5}
	a := s[2:4]
	a[0] = 42
	fmt.Println(s)
}

func exercise2() {
	s := make([]int, 3, 5)
	s[0] = 1
	s[1] = 2
	s[2] = 3

	modifySlice(s[:2])

	s = append(s, 4)
	fmt.Println(s)
}

func modifySlice(slice []int) {
	slice = append(slice, 100)
	slice[0] = 999
}

func exercise3() {
	s1 := []int{1, 2, 3, 4}
	s2 := modifyData(s1[:2])
	s2 = append(s2, 700)

	s1[1] = 50
	fmt.Println(s1)
	fmt.Println(s2)
}

func modifyData(data []int) []int {
	data[1] = 1000
	data = append(data, 500)
	data[2] = 600
	return data
}

func exercise4() {
	original := make([]int, 3, 4)
	original[0] = 10
	original[1] = 20
	original[2] = 30

	newSlice := doubleSlice(original[:2])

	original = append(original, 40)
	newSlice[1] = 100

	fmt.Println(original)
	fmt.Println(newSlice)
}

func doubleSlice(slice []int) []int {
	slice = append(slice, slice...)
	return slice
}

func exercise5(slice []int) []int {
	slice[0] = 200
	slice = append(slice, 1000)
	return slice
}

func exercise6() {
	s := []int{1, 2, 3, 4}
	res := exercise5(s)
	s[1] = 222
	fmt.Println(s)
	fmt.Println(res)

}

func modify(slice *[]int) {
	*slice = append(*slice, 555)
	(*slice)[1] = 999
}

// extra hard
func exercise7() {
	s := make([]int, 2, 4)
	s[0] = 1
	s[1] = 2
	modify(&s)
	s = append(s, 3)
	fmt.Println(s)

}

func addElements(slice []int, values ...int) []int {
	return append(slice, values...)
}

// extra hard
func exercise8() {
	original := []int{1, 2}
	modified := addElements(original, 3, 4, 5)
	original[0] = 100
	fmt.Println(original)
	fmt.Println(modified)

}

func main() {
	fmt.Println("Упражнение 1:")
	exercise1()

	fmt.Println("Упражнение 2:")
	exercise2()

	fmt.Println("Упражнение 3:")
	exercise3()

	fmt.Println("Упражнение 4:")
	exercise4()

	fmt.Println("Упражнение 6:")
	exercise6()

	fmt.Println("Упражнение 7:")
	exercise7()

	fmt.Println("Упражнение 8:")
	exercise8()
}
