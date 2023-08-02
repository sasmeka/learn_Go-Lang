package taks

import (
	"strconv"
	"strings"
)

func Convertion(num float64) string {
	var num_to_str string = strconv.FormatFloat(num, 'f', -1, 64) + "0" // konversi ke string dan menambahkan 0 jika angka yang dimasukkan hanya 1 dibelakang koma

	if strings.Contains(num_to_str, ".") == false { // cek apakah angka yang dimasukkan desimal atau tidak
		return "you must enter a decimal number"
	}

	var num_arr = strings.Split(num_to_str, ".")  // memisahkan angka dengan koma
	var num_arr_1 = strings.Split(num_arr[1], "") // mengambil nilai array index 1 pada num_array

	for i := len(num_arr_1) - 1; i >= 2; i-- { // proses angka pada index di atas 1 pada num_array_1 dan mengurutkannya dari belakang
		numm_before, _ := strconv.Atoi(num_arr_1[i-1])
		numm, _ := strconv.Atoi(num_arr_1[i])
		if numm > 5 {
			num_arr_1[i-1] = strconv.Itoa(numm_before + 1)
		}
	}

	// proses index 0 & 1 pada num_array
	num_arr_1_0, _ := strconv.Atoi(num_arr_1[0])
	num_arr_1_1, _ := strconv.Atoi(num_arr_1[1])
	if num_arr_1_1 > 5 {
		num_arr_1[0] = strconv.Itoa(num_arr_1_0 + 1)
		num_arr_1[1] = strconv.Itoa(0)
	} else {
		num_arr_1[1] = strconv.Itoa(0)
	}

	num_arr[1] = num_arr_1[0] + num_arr_1[1]

	result, _ := strconv.ParseFloat(strings.Join(num_arr, "."), 64)
	return strconv.FormatFloat(result, 'f', 2, 64)
}
