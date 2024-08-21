package main

import (
	"fmt"
)

var films = [4]string{"Shawshank Redemption", "Parasite", "Back To the Future", "Eternal Sunshine of the Spotless Mind"}

func movies() {
	for i := 0; i < 4; i++ {
		fmt.Printf("%v: %v \n", i+1, films[i])
	}
}

func main() {
	movies()
}
