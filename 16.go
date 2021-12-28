package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

type Sortable interface {
	Quicksort()
	IsSorted() bool
}

type mySlice struct {
	arr    []int
	nElems int
}

var _ Sortable = (*mySlice)(nil)

func NewRandIntSlice(k int, n int) Sortable {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = (rand.Intn(2)*2 - 1) * rand.Intn(k)
	}
	return &mySlice{s, n}
}

func (s *mySlice) Quicksort() {
	s.recQuicksort(0, s.nElems-1)
}

func (s *mySlice) recQuicksort(left, right int) {
	if right-left+1 <= 3 {
		s.manualSort(left, right)
		return
	}
	median := s.medianOf3(left, right)
	p := s.partition(left, right-1, median)
	s.recQuicksort(left, p-1)
	s.recQuicksort(p+1, right)
}

func (s *mySlice) partition(left, right, pivot int) int {
	l := left
	r := right - 1
	for l < r {
		if s.arr[l] < pivot {
			l++
			continue
		}
		if s.arr[r] > pivot {
			r--
			continue
		}
		s.swap(l, r)
		l++
		r--
	}
	if s.arr[l] < pivot {
		l++
	}
	s.swap(l, right)
	return l
}

func (s *mySlice) medianOf3(left, right int) int {
	center := (left + right) / 2
	if s.arr[left] > s.arr[center] {
		s.swap(left, center)
	}
	if s.arr[left] > s.arr[right] {
		s.swap(left, right)
	}
	if s.arr[center] > s.arr[right] {
		s.swap(center, right)
	}
	s.swap(center, right-1)
	return s.arr[right-1]
}

func (s *mySlice) manualSort(left, right int) {
	switch right - left + 1 {
	case 2:
		if s.arr[left] > s.arr[right] {
			s.swap(left, right)
		}
	case 3:
		if s.arr[left] > s.arr[right-1] {
			s.swap(left, right-1)
		}
		if s.arr[left] > s.arr[right] {
			s.swap(left, right)
		}
		if s.arr[right-1] > s.arr[right] {
			s.swap(right-1, right)
		}
	}
}

func (s *mySlice) swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}

func (s *mySlice) IsSorted() bool {
	for i := 0; i < s.nElems-1; i++ {
		if s.arr[i] > s.arr[i+1] {
			return false
		}
	}
	return true
}

func (s *mySlice) String() string {
	str := make([]string, s.nElems)
	for i, x := range s.arr {
		str[i] = strconv.Itoa(x)
	}
	return fmt.Sprint("[" + strings.Join(str, " ") + "]")
}

func main() {
	rand.Seed(time.Now().Unix())

	s := NewRandIntSlice(100, 20)
	fmt.Println(s)

	start := time.Now()
	s.Quicksort()
	total := time.Since(start)
	fmt.Println(s)

	if s.IsSorted() {
		fmt.Printf("sorted in %fs\n", total.Seconds())
	} else {
		fmt.Println("not sorted")
	}
}
