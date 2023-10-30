package main

import (
	"fmt"
	model "lesson4/exercises/models"
)

func AddPersonMap(p []model.Person)map[string]int {
	result := make(map[string]int, 0)
	for _, el := range p{
		result[el.Name] = el.Age
	}
	return result
}

func AddProductMap(p []model.Product)map[string][]model.Product {
	result := make(map[string][]model.Product)

	for _, el := range p{
		result[el.Category] = append(result[el.Category], model.Product{Name: el.Name, Price: el.Price})
	}
	return result
}

func AddCitizen(m map[string][]model.Citizen)map[int]int{
	result := make(map[int]int, 0)
	for _, arr := range m {
		for _, item := range arr{
			_, ok := result[item.Age]
			if ok {
				result[item.Age] += 1
			} else {
				result[item.Age] = 1
			}
		}
	}

	return result
}
func AddStudent(s []model.Student) map[string]map[int]float64 {
	result := make(map[string]map[int]float64)

	for _, el := range s {
		yearMap := make(map[int]float64, 0)
		KeySpeciality, ok := result[el.Speciality]
		if ok {
			YearKey, ok := KeySpeciality[el.Enrollment]
			if ok {
				//при условии что в map уже есть такой год происходит пересчёт среднего балла
				yearMap[el.Enrollment] = (YearKey + el.AverageMark) / 2
				result[el.Speciality] = yearMap
			} else {
				for key, item := range KeySpeciality {
					yearMap[key] = item
				}
			}
			yearMap[el.Enrollment] = el.AverageMark
			result[el.Speciality] = yearMap

		} else {
			yearMap[el.Enrollment] = el.AverageMark
			result[el.Speciality] = yearMap
		}
	}

	return result
}

func main()  {
	//task1
	persons := make([]model.Person, 0)
	persons = append(persons, model.Person{Name: "Alex", Age: 25})
	persons = append(persons, model.Person{Name: "Bob", Age: 30})

	PersonMap := AddPersonMap(persons)
	fmt.Println(PersonMap)

	//task2
	products := make([]model.Product, 0)
	products = append(products, model.Product{Name: "Laptop", Category: "Electronics", Price: 1000.5})
	products = append(products, model.Product{Name: "Shirt", Category: "Clothes", Price: 25.5})
	products = append(products, model.Product{Name: "Shirt2", Category: "Clothes", Price: 125.5})

	ProductMap := AddProductMap(products)
	fmt.Println(ProductMap)

	//task3
	city := make(map[string][]model.Citizen)
	c1 := make([]model.Citizen, 0)
	c1 = append(c1, model.Citizen{Name: "Alex", Age: 25})
	c1 = append(c1, model.Citizen{Name: "Bob", Age: 25})
	city["CityA"] = c1

	c2 := make([]model.Citizen, 0)
	c2 = append(c2,model.Citizen{Name: "Charlie", Age: 30})
	city["CityB"] = c2
	CitiPersonCount := AddCitizen(city)
	fmt.Println(CitiPersonCount)

	//task4
	students := make([]model.Student, 0)
	students = append(students, model.Student{Name: "Alex", Speciality: "Math", Enrollment: 2020, AverageMark: 4.5})
	students = append(students, model.Student{Name: "Bob", Speciality: "Math", Enrollment: 2021, AverageMark: 4.0})
	students = append(students, model.Student{Name: "Charlie", Speciality: "Physics", Enrollment: 2020, AverageMark: 4.8})
	students = append(students, model.Student{Name: "Elena", Speciality: "Physics", Enrollment: 2020, AverageMark: 4.4})
	resultStudent := AddStudent(students)
	fmt.Println(resultStudent)
	return
}