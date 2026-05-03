package main

import (
	"fmt"
)

func main() {
	license := true
	age := 19

	switch {
	case license && age > 15:
		fmt.Println("Puede seguir avanzando")
	case age <= 15 || !license:
		fmt.Println("No puede seguir circulando")
	}
}
