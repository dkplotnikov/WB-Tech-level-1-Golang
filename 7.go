package main

// Реализовать конкурентную запись данных в map.

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Структура с mutex
type Counter struct {
	mu        sync.Mutex
	numsCount map[int]int
}

// Count заполняет map данными о количестве вхождений эдемента в слайс
// На время работы с map блокируется Counter
func (cnt *Counter) Count(nums []int) {
	for _, num := range nums {
		cnt.mu.Lock()
		cnt.numsCount[num]++
		cnt.mu.Unlock()
	}
}

// syncCount заполняет sync.Map
func syncCount(nums []int, sm *sync.Map) {
	for _, num := range nums {
		n, loaded := sm.LoadOrStore(num, 1)
		if loaded {
			sm.Store(num, n.(int)+1)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixMilli())

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(10)
	}
	fmt.Println(nums)

	cnt := Counter{numsCount: make(map[int]int)}

	// Две рутины вызывают Count от своей половины слайса
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		cnt.Count(nums[:len(nums)/2+1])
		wg.Done()
	}()
	go func() {
		cnt.Count(nums[len(nums)/2+1:])
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("mutex: ", cnt.numsCount)

	// То же для sync.Map
	sm := &sync.Map{}
	wg.Add(2)
	go func() {
		syncCount(nums[:len(nums)/2+1], sm)
		wg.Done()
	}()
	go func() {
		syncCount(nums[len(nums)/2+1:], sm)
		wg.Done()
	}()
	wg.Wait()

	fmt.Print("sync.Map: ")
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("%d:%d ", key, value)
		return true
	})
	fmt.Println()

}
