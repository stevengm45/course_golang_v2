package main

import (
	"fmt"
)

func main() {
	// numeros := []int{1, 2, 3, 3, 4, 5, 1, 6}

	// for i := 0; i <= len(numeros); i++ {
	// 	for j := i + 1; j < len(numeros); j++ {
	// 		if numeros[i] == numeros[j] {
	// 			fmt.Println(numeros[i])
	// 		}
	// 	}
	// }

	// s := "(())))))"
	// right := 0
	// left := 0
	// for i := 0; i < len(s); i++ {
	// 	if s[i] == '(' {
	// 		right++
	// 	} else if s[i] == ')' && right > 0 {
	// 		right--
	// 	} else {
	// 		left++
	// 	}
	// }

	// fmt.Println(left + right)

	// r := 0
	// l := 0
	// i := 0
	// for i < len(s) {
	// 	if s[i] == '(' {
	// 		r++
	// 	} else {
	// 		if i+1 < len(s) && s[i+1] == ')' {
	// 			if r == 0 {
	// 				l++
	// 			} else {
	// 				r--
	// 				i++
	// 			}
	// 		} else {
	// 			l++
	// 			if l == 0 {
	// 				l++
	// 			} else {
	// 				r--
	// 			}
	// 		}
	// 	}
	// 	i++
	// 	fmt.Println(len(s) + 2*r)
	// }

	message := "abaasass"
	var palabra string
	i := 0
	fmt.Println(palabra)
	for i < len(message) {
		n := 1
		j := 1
		for i < (len(message)-1) && message[i] == message[i+1] {
			n = n + 1
			i = i + 1

		}
		if n == 1 {
			palabra = palabra + string(message[j])
			fmt.Println(palabra)
		} else {

			palabra = palabra + string(message[j]) + string(n)
			fmt.Println(palabra)
		}

		i = i + 1
	}

	fmt.Println(palabra)
}
