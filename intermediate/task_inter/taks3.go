package task_inter

import (
	"fmt"
)

func sort(n int, c chan<- int) {
	mtx.Lock()
	defer func() {
		close(c)
		mtx.Unlock()
		wg.Done()
	}()
	for i := 1; i <= n; i++ {
		c <- i
	}
}

func ganjil_genap(c chan int) {
	mt.RLock()
	defer func() {
		mt.RUnlock()
		wg.Done()
	}()
	for i := range c {
		if i%2 == 0 {
			fmt.Println(i, ": Genap")
		} else {
			fmt.Println(i, ": Ganjil")
		}
	}
}

func Display_taks3(v int) {
	c := make(chan int, v)
	wg.Add(2)
	go sort(cap(c), c)
	go ganjil_genap(c)
	wg.Wait()
}
