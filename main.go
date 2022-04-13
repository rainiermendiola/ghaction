package main

import (
	"fmt"
	"os"
)

func main() {
	varGHSecret := os.Getenv("VAR_FROM_GH_SEECRETS")
	fmt.Println("VAR_FROM_GH_SEECRETS: ", varGHSecret)
}
