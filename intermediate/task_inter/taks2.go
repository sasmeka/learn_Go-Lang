package task_inter

import (
	"fmt"
	"strconv"
	"strings"
)

type DeretBilang2 struct {
	Limit int
}

func Display_taks2(v int) {
	result_fibo := make(chan string)
	var limit = DeretBilang2{v}

	// for i := 0; i < 2; i++ {
	wg.Add(2)
	go limit.Fibonacci2(result_fibo)
	go limit.GanjilGenap(result_fibo)
	// }

	wg.Wait()
}

func (dB DeretBilang2) Fibonacci2(ch chan<- string) {
	mtx.Lock()
	defer func() {
		mtx.Unlock()
		wg.Done()
	}()
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
	ch <- result
}

func (dB DeretBilang2) GanjilGenap(ch <-chan string) {
	mt.RLock()
	defer func() {
		mt.RUnlock()
		wg.Done()
	}()
	var result_arr = strings.Split(<-ch, ", ")
	for i := 0; i < len(result_arr)-1; i++ {
		var num, _ = strconv.Atoi(result_arr[i])
		if num%2 == 1 {
			fmt.Println(result_arr[i], " : Ganjil")
		} else {
			fmt.Println(result_arr[i], " : Genap")
		}
	}
}
