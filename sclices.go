package main

import (
	"fmt"
)

func slices() {
	// Исходный массив
	arr := [5]int{1, 2, 3, 4, 5}

	// Создаем слайс на основе исходного массива
	originalSlice := arr[:]

	// Создаем новый слайс на основе оригинального слайса
	newSlice := originalSlice[1:4] // [2, 3, 4]

	fmt.Println(arr)

	// Выводим значения обоих слайсов
	fmt.Println("Original Slice:", originalSlice) // [1, 2, 3, 4, 5]
	fmt.Println("New Slice:", newSlice)           // [2, 3, 4]

	// Меняем значение в новом слайсе
	newSlice[0] = 99
	fmt.Println(arr)
	// Так как оба слайса указывают на один и тот же базовый массив,
	// изменения в одном слайсе отразятся и в другом
	fmt.Println("After modifying newSlice:")
	fmt.Println("Original Slice:", originalSlice) // [1, 99, 3, 4, 5]
	fmt.Println("New Slice:", newSlice)           // [99, 3, 4]

	var s []int // s == nil

	big := make([]int, 1000000)
	small := big[:3] // память под big не будет освобождена до тех пор, пока small в использовании

	a := []int{1, 2, 3}
	b := append(a, 4)
	a[0] = 42
	// b все еще [1, 2, 3, 4], не видит изменений a

	a := []int{1, 2, 3}
	b := []int{4, 5}
	copy(a, b)
	// a теперь [4, 5, 3]

	// Use make to create an empty slice of integers.
	sliceOne := make([]int, 0)
	// Use a slice literal to create an empty slice of integers.
	sliceTwo := []int{}
	fmt.Println(sliceOne == nil) // This will print false
	fmt.Println(len(sliceOne))   // This will print 0
	fmt.Println(cap(sliceOne))   // This will print 0
	fmt.Println(sliceTwo == nil) // This will print false
	fmt.Println(len(sliceTwo))   // This will print 0
	fmt.Println(cap(sliceTwo))   // This will print 0

	/* Create a slice of integers. Contains a
	length and capacity of 5 elements.*/
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)      // Print [10 20 30 40 50]
	fmt.Println(len(slice)) // Print  5
	fmt.Println(cap(slice)) // Print  5
	/* Create a new slice.Contains a length
	   of 2 and capacity of 4 elements.*/
	newSlice := slice[1:3]
	fmt.Println(slice)         //Print [10 20 30 40 50]
	fmt.Println(len(newSlice)) //Print 2
	fmt.Println(cap(newSlice)) //Print 4
}
