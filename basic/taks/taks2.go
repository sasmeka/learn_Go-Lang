package taks

import (
	"fmt"
	"strconv"
)

type DeretBilang struct {
	Limit int
}

// ====================Prima====================
// bilangan yang hanya dapat dibagi dengan 1 dan bilangan itu sendiri
func hitPrima(limit *int) string {
	result := ""
	for i := 1; i <= *limit; i++ {
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
	return result
}

func (dB DeretBilang) Prima() {
	fmt.Print("Prima: ")
	fmt.Print(hitPrima(&dB.Limit))
	fmt.Print("\n")
}

// ====================Ganjil====================
func hitGanjil(limit *int) string {
	result := ""
	for i := 1; i <= (*limit); i++ {
		if i%2 == 1 {
			result = result + strconv.Itoa(i) + ", "
		}
	}
	return result
}

func (dB DeretBilang) Ganjil() {
	fmt.Print("Ganjil: ")
	fmt.Print(hitGanjil(&dB.Limit))
	fmt.Print("\n")
}

// ====================Genap====================
func hitGenap(limit *int) string {
	result := ""
	for i := 1; i <= (*limit); i++ {
		if i%2 == 0 {
			result = result + strconv.Itoa(i) + ", "
		}
	}
	return result
}

func (dB DeretBilang) Genap() {
	fmt.Print("Genap: ")
	fmt.Print(hitGenap(&dB.Limit))
	fmt.Print("\n")
}

// ====================Genap====================
func hitFibonacci(limit *int) string {
	var arr_fibo = make([]int, *limit)
	arr_fibo[0] = 0
	arr_fibo[1] = 1
	for i := 2; i < (*limit); i++ {
		arr_fibo[i] = arr_fibo[i-1] + arr_fibo[i-2]
	}

	result := ""
	for _, v := range arr_fibo {
		result = result + strconv.Itoa(v) + ", "
	}
	return result
}

func (dB DeretBilang) Fibonacci() {
	fmt.Print("Fibonacci: ")
	fmt.Print(hitFibonacci(&dB.Limit))
	fmt.Print("\n")
}
