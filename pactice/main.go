package main

import (
	"bufio"
	"fmt"
	"os"
)

const aConstant int = 10

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number: ")
	text, _ := reader.ReadString('\n')

	fmt.Println("You entered: ", text)
}