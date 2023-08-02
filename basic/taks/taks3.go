package taks

import (
	"fmt"
)

type ValuesBidang struct {
	Panjang float64
	Lebar   float64
	Tinggi  float64
}

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

func hitLuas(panjang *float64, lebar *float64) float64 {
	return *panjang * *lebar
}
func hitKeliling(panjang *float64, lebar *float64) float64 {
	return 2 * (*panjang + *lebar)
}
func hitVolume(panjang *float64, lebar *float64, tinggi *float64) float64 {
	return *panjang * *lebar * *tinggi
}

func (vB ValuesBidang) luas() float64 {
	return hitLuas(&vB.Panjang, &vB.Lebar)
}
func (vB ValuesBidang) keliling() float64 {
	return hitKeliling(&vB.Panjang, &vB.Lebar)
}
func (vB ValuesBidang) volume() float64 {
	return hitVolume(&vB.Panjang, &vB.Lebar, &vB.Tinggi)
}

func Proseshitung(h hitung) {
	fmt.Println("Luas: ", h.luas())
	fmt.Println("Keliling: ", h.keliling())
	fmt.Println("Volume: ", h.volume())
}
