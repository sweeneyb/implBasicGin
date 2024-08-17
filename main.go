package main

import ( 
	"fmt"
	"os"
	"bufio"
)

func main() {
  reader := bufio.NewReader(os.Stdin)

  for {
	fmt.Print("> ")
	input, _ := reader.ReadString('\n')

	fmt.Printf("You said %v", input)
  }
}