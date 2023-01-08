package main

import "fmt"

func saludo() string{
	return "Hola como estas"
}

func main(){
	fmt.Println("Hola como estas")
	fmt.Println(saludo())
}
