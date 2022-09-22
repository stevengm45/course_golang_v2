package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	c := "Hola asiatico"
	fmt.Printf("%q tiene %v runas\n", c, utf8.RuneCountInString(c))

	cad := "una cadena"
	// convirtiendo de string a porcion
	bytes := []byte(cad)
	runas := []rune(cad)

	// modificando el contenido de la sporciones
	bytes[0] = 'U'
	runas = append(runas, '!')

	// convirtiendo de nuevo de porcion a cadena
	str1 := string(bytes)
	str2 := string(runas)

	fmt.Println(str1)
	fmt.Println(str2)

	var sb strings.Builder
	sb.WriteString("strings.Builder\n")
	l := sb.Len()
	for i := 0; i < l; i++ {
		sb.WriteRune('=')
	}
	sb.WriteString("\nEs la forma mas eficiente de construir ")
	sb.WriteString("cadenas.\n")
	sb.WriteString("Sin embargo, solo se pueden agregar ")
	sb.Write([]byte("cadenas o porciones de bytes o runas. "))
	sb.WriteString("Para añadir otros tipos, primero debes  ")
	sb.WriteString("transformarlos. Por ej, con strconv.Itoa(")
	sb.WriteString(strconv.Itoa(1234))
	sb.WriteByte(')')

	fmt.Println(sb.String())

	

}
