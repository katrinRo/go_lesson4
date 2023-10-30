package models

type Person struct {
	Name string
	Age  int
}

type Product struct {
	Name     string
	Category string
	Price    float64
}

type Citizen struct {
	Name string
	Age  int
}

type Student struct {
	Name        string
	Speciality  string
	Enrollment  int // Год поступления
	AverageMark float64
}
