package task_inter

import "fmt"

func sum(d []int, ch chan int) {
	result := 0
	for _, v := range d {
		result += v
	}
	ch <- result
}

func Exm_main() {
	a := []int{7, 10, 2, 34, 33, -12, -8, 4}
	chn := make(chan int)

	go sum(a[len(a)/2:], chn)
	go sum(a[:len(a)/2], chn)

	chn1 := <-chn
	chn2 := <-chn
	fmt.Println(chn1, " + ", chn2, " = ", chn1+chn2)
}
