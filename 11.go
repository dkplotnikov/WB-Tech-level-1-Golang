package main

// Реализовать пересечение двух неупорядоченных множеств.

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// RandIntSet добавляет в возвращаемое множество n in [0,k) случайных чисел
func RandIntSet(k int, n int) *map[int]struct{} {
	set := make(map[int]struct{})
	for i := 0; i < n; i++ {
		set[rand.Intn(k)] = struct{}{}
	}
	return &set
}

// SetToString возвращает запись множества в строке
func SetToString(set *map[int]struct{}) string {
	var s []string
	for e := range *set {
		s = append(s, strconv.Itoa(e))
	}
	return "{" + strings.Join(s, " ") + "}"
}

// Intersect возвращает пересечение двух множеств
func Intersect(set1 *map[int]struct{}, set2 *map[int]struct{}) *map[int]struct{} {
	set := make(map[int]struct{})

	// Проверка идет по сету меньшего размера
	if len(*set1) > len(*set2) {
		set1, set2 = set2, set1
	}

	// В результирующий set добавляется элемент из set1, если он содержится в set2
	for e := range *set1 {
		if _, ok := (*set2)[e]; ok {
			set[e] = struct{}{}
		}
	}

	return &set
}

func main() {
	rand.Seed(time.Now().Unix())

	set1 := RandIntSet(10, 10)
	set2 := RandIntSet(10, 10)

	fmt.Println("set1:", SetToString(set1))
	fmt.Println("set2:", SetToString(set2))

	set3 := Intersect(set1, set2)
	fmt.Println("intersect:", SetToString(set3))
}
