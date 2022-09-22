package main

import (
	"fmt"
	"math/rand"
)

func main() {
	pasaportes := map[string]int{}

	fmt.Println(pasaportes)

	capitales := map[string]string{
		"España":      "Madris",
		"Francia":     "Paris",
		"Holanda":     "Amsterdan",
		"Reino Unido": "Londres",
		"Japon":       "Tokio",
	}

	// Agregar datos a un mapa
	capitales["Italia"] = "Roma"

	pais := "Francia"
	capital := capitales[pais]
	fmt.Printf("La capital de %s es %s\n", pais, capital)

	pais = "Narnia"
	capital2, ok := capitales[pais]

	if ok {
		fmt.Println("La capital de", pais, "es", capital2)
	} else {
		fmt.Println("NO he encontrado capital para", pais)
	}

	// Eliminar entrada con delete
	delete(capitales, "España")

	// Recorriendo mapas con range
	fmt.Println("=========Recorriendo mapas con range=========")
	for pais, capital := range capitales {
		fmt.Printf("La capital de %s es %s\n", pais, capital)
	}

	fmt.Println("========ignorar los valores============")
	// ignorar los valores
	for pais := range capitales {
		fmt.Printf("Pais %s \n", pais)
	}

	fmt.Println("========ignorar las claves==========")
	// ignorar las claves
	for _, capital := range capitales {
		fmt.Printf("Capital %s\n", capital)
	}

	// Contando el numero de elementos
	fmt.Println("======Contando el numero de elementos======")
	numero := len(capitales)
	fmt.Println("contabilizadas", numero, "capitales")

	fmt.Println("======Conjutnos======")

	fmt.Println("Los numeros ganadores de la loteria son:")
	numeros := map[int]struct{}{}
	for len(numeros) < 6 {
		n := rand.Intn(49) + 1
		// solo se muestra el numero si no ha salido antes
		if _, ok := numeros[n]; !ok {
			numeros[n] = struct{}{}
			fmt.Println("El", n, "...")
		}
	}
	fmt.Println("Felicitacion a los premiados!")
}
