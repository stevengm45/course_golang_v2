package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://instagram.com",
		"http://twitter.com",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor)
	}

	tiempoPasado := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", &tiempoPasado)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "no esta disponible =(")
	} else {
		fmt.Println(servidor, "esta funcionando normalmente :D")
	}
}
