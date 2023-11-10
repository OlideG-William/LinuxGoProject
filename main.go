package main

import (
	"fmt"
	"time"
)

func fif() {

}

func main() {
	statrtime := time.Now()

	fmt.Println("Hello world!")

	fmt.Println("Compile time program: ", time.Since(statrtime))
}
