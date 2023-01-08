package main

import "fmt"

func saludo() string{
	return "Hola como estas"
}

func suma(n int, x int) (int){
	return x + n
}

func main(){
	fmt.Println("Hola como estas")
	fmt.Println(saludo())
	fmt.Println(suma(1,2))
}
