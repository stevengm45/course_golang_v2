package main

import "fmt"

func main() {
	var pi *int
	pi = nil
	if pi == nil {
		fmt.Print("no puedo hacer nada con este apuntador ")
		fmt.Println("porque no apunta a nada!")
	}

	i := 10
	fmt.Println(i)
	p := &i
	fmt.Println(p)
	a := *p
	fmt.Println(a)
	*p = 21
	fmt.Println(*p)
	fmt.Println(i)

	i = 0
	j := 0
	p1 := &i
	p2 := &j
	if p1 == p2 {
		fmt.Println("p1 y p2 apuntan a la misma direccion de memoria")
	} else {
		fmt.Println("p1 y p2 apuntan a direcciones distintas")
	}
	if *p1 == *p2 {
		fmt.Println("p1 y p2 apuntan a valores iguales")
	} else {
		fmt.Println(("p1 y p2 apuntan a valores distintos"))
	}
}
