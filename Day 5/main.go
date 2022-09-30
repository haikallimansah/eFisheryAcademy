package main

import "fmt"

type item struct {
	string
	w, v int
}

var barang = []item{
	{"BenihLele", 50000, 1},
	{"PakanLeleCapMenara", 25000, 1},
	{"ProbiotikA", 75000, 1},
	{"ProbiotikNilaB", 10000, 1},
	{"PakanNila", 20000, 1},
	{"BenihNila", 20000, 1},
	{"Cupang", 5000, 1},
	{"BenihNila2", 30000, 1},
	{"BenihCupang", 10000, 1},
	{"ProbiotikB", 10000, 1},
}

const maxWt = 100000

func main() {
	items, w, v := hitung(len(barang)-1, maxWt)
	fmt.Println(items)
	fmt.Println("Total :", w)
	fmt.Println("Jumlah terbeli :", v)
	// min dan max
	var a = [10]int{50000, 25000, 75000, 10000, 20000, 20000, 5000, 30000, 10000, 10000}
	min, max := cariMinMax(a)
	fmt.Println("Termurah: ", min)
	fmt.Println("Termahal: ", max)
}

func hitung(i, w int) ([]string, int, int) {
	if i < 0 || w == 0 {
		return nil, 0, 0
	} else if barang[i].w > w {
		return hitung(i-1, w)
	}
	i0, w0, v0 := hitung(i-1, w)
	i1, w1, v1 := hitung(i-1, w-barang[i].w)
	v1 += barang[i].v
	if v1 > v0 {
		return append(i1, barang[i].string), w1 + barang[i].w, v1
	}
	return i0, w0, v0
}

func cariMinMax(a [10]int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
