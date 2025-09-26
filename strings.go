package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ucok Baba", "Ucok"))
	fmt.Println(strings.Split("Ucok Baba", " "))
	fmt.Println(strings.ToLower("Ucok Baba"))
	fmt.Println(strings.ToUpper("Ucok Baba"))
	fmt.Println(strings.Trim("            Ucok Baba            ", " "))
	fmt.Println(strings.ReplaceAll("Ucok Baba Ucok Baba", "Ucok", "Ucup"))
}
