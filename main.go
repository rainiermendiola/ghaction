package main

import (
	"fmt"
	"os"
)

func main() {
	varGHSecret := os.Getenv("VAR_FROM_GH_SECRET")
	if varGHSecret == "ThisIsMySecretValue" {
		fmt.Println("The secret variable is correct!")
	} else {
		fmt.Println("The secret variable is not retrieved")
	}
}
