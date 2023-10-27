package main

import (
	"fmt"
	"sync"
)

// Заголовок для Go map.
// type hmap struct {
// 	// Примечание: формат hmap также закодирован в cmd/compile/internal/reflectdata/reflect.go.
// 	// Убедитесь, что это остается в синхронизации с определением компилятора.
// 	count     int // Количество активных ячеек == размеру карты. Должно быть первым (используется встроенной функцией len())
// 	flags     uint8
// 	B         uint8  // log_2 от # корзин (может содержать до loadFactor * 2^B элементов)
// 	noverflow uint16 // приблизительное количество переполненных корзин; смотрите детали в incrnoverflow
// 	hash0     uint32 // начальное значение хеша

// 	buckets    unsafe.Pointer // массив из 2^B корзин. может быть nil, если count==0.
// 	oldbuckets unsafe.Pointer // предыдущий массив корзин вдвое меньшего размера, не-nil только при росте
// 	nevacuate  uintptr        // счетчик прогресса для эвакуации (корзины, меньше чем это, были эвакуированы)

// 	extra *mapextra // дополнительные поля
// }

func maps() {
	m := map[string]int{
		"apple":  5,
		"banana": 7,
	}

	// 1. Случайный порядок итерации:
	fmt.Println("Порядок итерации 1:")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Порядок итерации 2:")
	for k, v := range m {
		fmt.Println(k, v)
	}
	// При каждом выполнении порядок может меняться.

	// 2. Конкурентный доступ:
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		m["apple"] = 10
	}()
	go func() {
		defer wg.Done()
		m["banana"] = 15
	}()
	wg.Wait()
	// Этот код может вызвать панику выполнения из-за одновременной записи в мапу.

	// 3. Проверка на существование:
	value, exists := m["apple"]
	if exists {
		fmt.Println("Apple exists with value:", value)
	} else {
		fmt.Println("Apple does not exist.")
	}

	// 4. Nil мапы:
	var nilMap map[string]int
	fmt.Println(nilMap["key"]) // Это вызовет панику выполнения.

	// 5. Удаление элементов:
	delete(m, "apple")

	// 6. Размер мапы:
	fmt.Println("Размер мапы:", len(m))

	// 7. Копирование мап:
	copyMap := m
	copyMap["banana"] = 20
	fmt.Println("Original map:", m)
	fmt.Println("Copy map:", copyMap)

	// 8. Изменяемость ключей:
	invalidMap := map[[]int]string{} // Неверный код, так как срез не может быть ключом.

	// 9. Инициализация с начальным размером:
	largeMap := make(map[int]int, 1000)

	fmt.Println("End of demonstration.")
}

func setLikeType() {
	// В Go пустые структуры (`struct{}`) часто используются в качестве значения в мапах,
	// особенно когда значение само по себе не важно, и интерес представляет только наличие или отсутствие ключа.
	// Например, для создания набора уникальных элементов (аналога `set` в некоторых других языках) можно
	// использовать мапу с ключами типа элемента и значениями типа `struct{}`.
	// Пустая структура не занимает память, что делает ее особенно привлекательной для этой цели.

	set := make(map[string]struct{})

	// Добавить элементы в набор
	set["a"] = struct{}{}
	set["b"] = struct{}{}

	// Проверить наличие элемента
	if _, exists := set["a"]; exists {
		fmt.Println("a is in the set")
	}

	// Удалить элемент
	delete(set, "a")

}
