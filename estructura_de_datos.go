package main

import "fmt"

func Cero(vec [3]int, porc []int) {
	vec[0] = 0
	if len(porc) > 0 {
		porc[0] = 0
	}
}

func Suma(n ...int) int {
	total := 0
	for _, s := range n {
		total += s
	}
	return total
}

func main() {
	var arr [5]int

	fmt.Println(arr)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	arr4 := [3]int{1, 2, 3}
	suma := arr4[0]
	arr4[2] = 0

	fmt.Println(suma)
	fmt.Println(arr4)

	var tableroAjedrez [8][8]uint8
	fmt.Println(tableroAjedrez)

	p := []int{1, 2, 3}
	v1 := [3]int{1, 2, 3}
	v2 := [...]int{1, 2, 3}

	fmt.Println(p)
	fmt.Println(v1)
	fmt.Println(v2)

	var x []int
	x = append(x, 1, 2, 3)
	x = append(x, 6)
	fmt.Println(x)

	var sl []int
	fmt.Printf("longitud %v. capacidad %v\n", len(sl), cap(sl))
	sl = append(sl, 1, 2, 3, 4)
	fmt.Printf("longitud %v. capacidad %v\n", len(sl), cap(sl))
	sl = append(sl, 5)
	fmt.Printf("longitud %v. capacidad %v\n", len(sl), cap(sl))
	sl = append(sl, 6, 7, 8, 9)
	fmt.Printf("longitud %v. capacidad %v\n", len(sl), cap(sl))

	altaCapacidad := make([]float32, 0, 2048)
	fmt.Println(altaCapacidad)

	original := []int{1, 2, 3, 4, 5}
	copia := make([]int, len(original))

	n := copy(copia, original)
	fmt.Println(n, "numeros copiados:", copia)

	v := [3]int{1, 2, 3}
	p2 := []int{1, 2, 3}
	Cero(v, p2)
	fmt.Println("vector:", v)
	fmt.Println("porcion:", p2)

	ciudades := []string{"Tokyo", "Amsterdan", "Toronto", "California", "Londres"}
	for i := 0; i < len(ciudades); i++ {
		fmt.Printf("[%d] %s\n", i, ciudades[i])
	}

	for i, ciudad := range ciudades {
		fmt.Printf("[%d] %s\n", i, ciudad)
	}

	for _, ciudad := range ciudades {
		fmt.Println(ciudad)
	}

	base := []int{1, 0, 3, 4, 5}
	fmt.Println("base:", base)

	vista1 := base[1:3]
	fmt.Println("vista 1:", vista1)
	vista1[0] = 2
	fmt.Println("base:", base)

	vista2 := base[2:]
	fmt.Println("vista2:", vista2)

	vista3 := base[:3]
	fmt.Println("vista 3:", vista3)

	vista4 := base[:]
	fmt.Println("vista 4:", vista4)

	pruebasuma := Suma(1, 2, 3, 4)

	fmt.Println(pruebasuma)

	valores := []int{1, 2, 3, 4, 5, 6}
	total := Suma(valores...)
	fmt.Println(total)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s1 = append(s1, s2...)
	fmt.Println("concatenacion =", s1)
}
