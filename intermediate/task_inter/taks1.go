package task_inter

import (
	"fmt"
	"strconv"
	"sync"
)

var wg = sync.WaitGroup{}
var mt = sync.RWMutex{}
var mtx = sync.Mutex{}

type DeretBilang struct {
	Limit int
}

func Display_taks1(v int) {
	wg.Add(4)
	var limit = DeretBilang{v}
	go limit.Prima()
	go limit.Ganjil()
	go limit.Genap()
	go limit.Fibonacci()
	wg.Wait()
}

// ====================Prima====================
// bilangan yang hanya dapat dibagi dengan 1 dan bilangan itu sendiri
func (dB DeretBilang) Prima() {
	result := ""
	for i := 1; i <= dB.Limit; i++ {
		var check int = 0
		for ii := 1; ii <= i; ii++ {
			if i%ii == 0 {
				check = check + 1
			}
		}
		if check == 2 {
			result = result + strconv.Itoa(i) + ", "
		}
	}
	fmt.Println("Prima: ", result)
	wg.Done()
}

// ====================Ganjil====================
func (dB DeretBilang) Ganjil() {
	result := ""
	for i := 1; i <= (dB.Limit); i++ {
		if i%2 == 1 {
			result = result + strconv.Itoa(i) + ", "
		}
	}
	fmt.Println("Ganjil: ", result)
	wg.Done()
}

// ====================Genap====================

func (dB DeretBilang) Genap() {
	result := ""
	for i := 1; i <= (dB.Limit); i++ {
		if i%2 == 0 {
			result = result + strconv.Itoa(i) + ", "
		}
	}
	fmt.Println("Genap: ", result)
	wg.Done()
}

// ====================Genap====================
func (dB DeretBilang) Fibonacci() {
	var arr_fibo = make([]int, dB.Limit)
	arr_fibo[0] = 0
	arr_fibo[1] = 1
	for i := 2; i < (dB.Limit); i++ {
		arr_fibo[i] = arr_fibo[i-1] + arr_fibo[i-2]
	}

	result := ""
	for _, v := range arr_fibo {
		result = result + strconv.Itoa(v) + ", "
	}
	fmt.Println("Fibonacci: ", result)
	wg.Done()
}
