package main

import "fmt"

func main() {
	var password = "pass12345"
	saludo(password)
}

func saludo(password string) {
	fmt.Println("Hello, World!", password)
}
