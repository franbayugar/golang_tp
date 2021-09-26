package main

import (
	"fmt"

	"com.tudai2021/services"
)

func main() {
	var srv services.Chain
	var s string

	fmt.Println("Ingrese la cadena")
	fmt.Scan(&s)

	valid, err := services.ValidChain(s)
	if err != nil {
		ch, _ := srv.ModChain(s, valid)
		fmt.Println(ch)
	} else {
		fmt.Print("Error: ", err)
	}

}
