package main

import "fmt"

func Hola(nombre, apellido string) {
	fmt.Printf("!Hola, %s %s!\n", nombre, apellido)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxMin(a, b int) (max int, min int) {
	if a > b {
		max = a
		min = b
	} else {
		min = a
		max = b
	}
	return
}

func Incrementa(a *int) {
	*a++
}

func Suma(a, b int) int {
	return a + b
}

func Multiplica(a, b int) int {
	return a * b
}

func EsPar(v int64) bool {
	return v%2 == 0
}

// func() {
// 	fmt.Println("esto se invocara")
// 	fmt.Println("Directamente, pero en otro contexto")
// }()

func Int() int
func Intn(n int) int {
	return n
}
func Int31n(n int32) int32
func Int63n(n int64) int64

func main() {
	Hola("Steven", "Gonzalez")
	Hola("Alvaro", "Cruz")
	Max(4, 10)

	m := Max(3, 5)
	a := 10 + Max(m+2, 6)
	fmt.Println(a)

	max, min := MaxMin(34, 15)
	fmt.Println(max)
	fmt.Println(min)

	ma, _ := MaxMin(10, 29)
	fmt.Println(ma)

	s := 3
	fmt.Println("s = ", s)
	Incrementa(&s)
	fmt.Println("s = ", s)

	var generador func() int
	contador := 0
	generador = func() int {
		contador++
		return contador
	}
	fmt.Println("generador contador:", generador(), generador(), generador())

	var operador func(int, int) int
	operador = Suma
	fmt.Println("suma=", operador(3, 4))
	operador = Multiplica
	fmt.Println("multiplica =", operador(3, 4))

	var val int32 = 123
	fmt.Println("Par?", EsPar(int64(val)))

	fmt.Println(Int())
	fmt.Println(Intn(10))
	fmt.Println(Int31n(20))
	fmt.Println(Int63n(11))
}
