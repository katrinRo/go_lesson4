package main

import (
	"fmt"
	"unsafe"
)

type MyMarker struct{}

type TreeNode struct {
	Value    int
	Children []*TreeNode
}

func (t *TreeNode) Next() {
	for _, v := range t.Children {
		v.Next()
	}
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

type Address struct {
	Street string
	City   string
}

type User struct {
	Name    string
	Age     int
	Address Address
}

type Employee struct {
	Person
	Position string
}

type PersonWithTags struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func structsBasic() {
	var person1 Person // все поля будут иметь нулевые значения: "", "", 0
	fmt.Println(person1)
	// name := person1.FirstName

	person3 := struct {
		FirstName string
		Age       int
	}{
		FirstName: "Jane",
		Age:       25,
	}
	fmt.Println(person3)
	fmt.Println(person3.FirstName)
	fmt.Println(person3.Age)

	e := Employee{}
	e.FullName()

	person2 := Person{
		FirstName: "John",
		LastName:  "Doe", Age: 30,
	}
	fmt.Println(person2)

	// empty structs
	a := struct{}{}
	b := struct{}{}
	fmt.Printf("%p %p\n", &a, &b) // Выводит один и тот же адрес для обоих экземпляров

	set := make(map[string]struct{})

	// Добавить элементы в набор
	set["a"] = struct{}{}
	set["a"] = struct{}{}
	set["b"] = struct{}{}

	// Проверить наличие элемента
	if _, exists := set["a"]; exists {
		fmt.Println("a is in the set")
	}

	// Удалить элемент
	delete(set, "a")

}

type MyStruct struct {
	Field int
}

func (m MyStruct) GetValue() int {
	return m.Field
}

func (m MyStruct) SetValue(val int) {
	m.Field = val
}

func (m *MyStruct) SetValueByPointer(val int) {
	m.Field = val
}

func testMethod() {
	obj := MyStruct{
		Field: 1,
	}
	obj.SetValue(2)
	fmt.Println(obj.Field) // Вывод: 1

	obj.SetValueByPointer(2)
	fmt.Println(obj.Field) // Вывод: 2
}

func structsComparison() {
	type Point struct {
		X, Y int
	}

	p1 := Point{1, 2}
	p2 := Point{1, 2}
	fmt.Println(p1 == p2) // true

	type PersonNonComparable struct {
		Name    string
		Friends []string // слайсы не сравниваемы
	}
	pe1 := PersonNonComparable{"foo", []string{"bar"}}
	pe2 := PersonNonComparable{"foo", []string{"bar"}}
	// person1 == person2 будет ошибкой компиляции
	fmt.Println(pe1 == pe2) // ошибкой компиляции

}

func NewPerson() *Person {
	return &Person{
		Age: 18, // значение по умолчанию
	}
}

func structs() {

}

func structsMemoryAlignment() {
	type Example struct {
		a byte  // 1 byte
		b int64 // 8 bytes
		c byte  // 1 byte
	}
	x := &Example{}
	y := Example{}

	fmt.Println(unsafe.Sizeof(x)) // 8
	fmt.Println(unsafe.Sizeof(y)) // 0

	// (8 bytes) + a (1 byte) + c (1 byte) + padding (6 bytes) = 16 bytes

	type ExampleOptimized struct {
		b int64 // 8 bytes
		a byte  // 1 byte
		c byte  // 1 byte
	}

	x2 := &ExampleOptimized{}
	y2 := ExampleOptimized{}

	fmt.Println(unsafe.Sizeof(x2)) // 8
	fmt.Println(unsafe.Sizeof(y2)) // 0
	// (1 byte) + padding (7 bytes) + b (8 bytes) + c (1 byte) + padding (7 bytes) = 24 bytes
}

type Base struct {
	Name string
}

type Derived struct {
	Base
	Age int
}

func (b *Base) Greet() {
	fmt.Println("Привет,", b.Name)
}

func (d *Derived) Greet() {
	d.Base.Greet()
	fmt.Println("Мне", d.Age, "лет")
}

func structEmbedding() {

	d := Derived{
		Base: Base{
			Name: "John",
		},
		Age: 25,
	}
	d.Greet() // Выведет: Привет, John

	d.Base.Greet()

}
