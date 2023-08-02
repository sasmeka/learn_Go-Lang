package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	// segitiga(7)
	// fmt.Println("password awal: fazztrack")
	// genPass("fazztrack", "easy")
	// genPass("fazztrack", "medium")
	// genPass("fazztrack", "strong")

	defer fmt.Println("defer 1")

	defer fmt.Println("defer 2")

	fmt.Println("halo")
	// film := []int{
	// 	1, 7, 3, 4, 8, 4, 1,
	// }
	// show_recomendation(film, 5)
}

// =================Taks 1=================
func segitiga(angka int) {
	for i := angka; i >= 1; i-- {
		if angka-i != 0 {
			for iii := 1; iii <= (angka - i); iii++ {
				fmt.Print(" ")
			}
		}
		for ii := 1; ii <= ((i * 2) - 1); ii++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

// =================Taks 2=================

var randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))

func genPass(pass string, level string) {
	pass_transform := strings.Split(pass, "")
	// var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var numbers = "0123456789"
	var schar = "~`!@#$%^&*()_-+={[}]|:;<,>.?/"

	if level == "" {
		level = "low"
	}
	if len(pass_transform) < 6 {
		fmt.Print("Password must be at least 6 characters long.")
	} else {
		lvl_arr := make([]int, 4)
		if level == "strong" {
			lvl_arr[0] = len(pass_transform) * 2 / 4
			lvl_arr[1] = len(pass_transform) * 2 / 4
			lvl_arr[2] = len(pass_transform) * 2 / 4
		} else if level == "medium" {
			lvl_arr[0] = len(pass_transform) * 0
			lvl_arr[1] = len(pass_transform) * 2 / 4
			lvl_arr[2] = len(pass_transform) * 2 / 4
		} else {
			lvl_arr[0] = len(pass_transform) * 0
			lvl_arr[1] = len(pass_transform) * 2 / 4
			lvl_arr[2] = len(pass_transform) * 0
		}
		result := []string{transformPass(pass, lvl_arr[0]), randoM(numbers, lvl_arr[1]), randoM(schar, lvl_arr[2])}
		fmt.Println(level, " : ", strings.Join(result, ""))
	}
}

func transformPass(pass string, lvl int) string {
	var lettersarr = strings.Split(pass, "")
	result := make([]int, lvl)
	for i := 0; i < lvl; i++ {
		result[i] = randomizer.Intn(len(lettersarr))
	}
	for _, ii := range result {
		lettersarr[ii] = strings.ToUpper(lettersarr[ii])
	}
	letters_result := strings.Replace(strings.Join(lettersarr, ""), " ", "-", -1)
	return letters_result
}

func randoM(data string, length int) string {
	var letters = []rune(data)
	result := make([]rune, length)
	for i := range result {
		result[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(result)
}

// =================Taks 3=================

func filter_film(film []int, total int) map[string]string {
	result := map[string]string{}
	for i, v := range film {
		for a := 0; a < len(film); a++ {
			if a != i {
				if v+film[a] == total {
					if i < a {
						result[strconv.Itoa(i)+""+strconv.Itoa(a)] = strconv.Itoa(i) + " & " + strconv.Itoa(a)
					} else {
						result[strconv.Itoa(a)+""+strconv.Itoa(i)] = strconv.Itoa(a) + " & " + strconv.Itoa(i)
					}
				}
			}
		}
	}
	return result
}

func show_recomendation(film []int, total int) {
	for _, v := range filter_film(film, total) {
		fmt.Println(v)
	}
}
